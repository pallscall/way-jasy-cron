package ent_ex

import (
	"context"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"way-jasy-cron/common/ecron"
	utilpaladin "way-jasy-cron/common/util/paladin"
	"way-jasy-cron/cron/ecode"
	"way-jasy-cron/cron/internal/model/ent"
)

type JobStatus int

const (
	JobDelete   JobStatus = -1
	JobRunning  JobStatus = 1
	JobStopping JobStatus = 0
)

type ListJobReq struct {
	ListBaseReq
	Name    string `form:"name"`
	Creator string `form:"creator"`
	Comment string `form:"comment"`
}

type ListJobOptions struct {
	ListBaseOptions
	Name    string
	Creator string
	Comment string
}

type ListJobResp struct {
	ListBaseResp
	Jobs []*ent.Job `json:"jobs"`
}

func (r ListJobReq) ToListJobOptions() (o *ListJobOptions) {
	o = &ListJobOptions{
		ListBaseOptions: ListBaseOptions{
			PN: r.PN,
			PS: r.PS,
		},
		Name:    r.Name,
		Creator: r.Creator,
		Comment: r.Comment,
	}
	o.Completed()
	return
}

type CronResp struct {
	IDs		[]int 	`json:"ids"`
	Total	int		`json:"total"`
}

type closer func(ctx context.Context, id, opt int) error

type JobExecutor struct {
	job *ent.Job
	httpClient *http.Client
	closer func(ctx context.Context, id, opt int) error
	RetryCount     int
	RetryTempCount int
}

type Manager struct {
	RetryCount int
}

type Config struct {
	RetryCount int
}

func (c *Config) Filename() string {
	return "job.toml"
}

func New() *Manager {
	config := &Config{}
	utilpaladin.MustUnmarshalTOML(config)
	return &Manager{
		RetryCount: config.RetryCount,
	}
}

func (m *Manager) Create(
	job *ent.Job,
	closer closer,
	client *http.Client,
) *JobExecutor {
	return &JobExecutor{
		job: job,
		closer: closer,
		httpClient: client,
		RetryCount: m.RetryCount,
		RetryTempCount: m.RetryCount,
	}
}

func (j  *JobExecutor) Validate() error {
	spec := j.job.Spec
	if _, err := ecron.ParseStandard(spec); err != nil {
		return errors.Wrap(ecode.InvalidSpec, "job spec is incorrect！")
	}
	return nil
}

func (j *JobExecutor) Run() {
	var (
		err error
		request *http.Request
	)
	defer func() {
		if err != nil {
			j.RetryCount--
			if j.RetryCount == 0 {
				log.Info("close the job:", j.job)
				if err = j.closer(context.TODO(), j.job.ID, int(JobStopping)); err != nil {
					log.Error("method: Run#ent_ex/job err:",err)
				}
			}
			return
		}
		j.RetryCount = j.RetryTempCount
	}()
	request, err = http.NewRequest(j.job.Method, j.job.URL, strings.NewReader(j.job.Body))
	if err != nil {
		log.Error("NewRequest err:", err)
		return
	}
	_, err = j.httpClient.Do(request)
	if err != nil {
		log.Error("do request err:", err)
	}
	return
}

func (j *JobExecutor) EntryID() ecron.EntryID{
	return ecron.EntryID(j.job.ID)
}

func (j *JobExecutor) GetSpec() string{
	return j.job.Spec
}