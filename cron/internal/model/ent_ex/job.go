package ent_ex

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"time"

	"net/http"
	"strings"
	"way-jasy-cron/common/ecron"
	"way-jasy-cron/common/email"
	"way-jasy-cron/cron/ecode"
	"way-jasy-cron/cron/internal/dao/mail"
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
type logger func(ctx context.Context, msg, operator string) error

type JobExecutor struct {
	job *ent.Job
	httpClient *http.Client
	closer func(ctx context.Context, id, opt int) error
	email          *mail.Manager
	logger func(ctx context.Context, msg, operator string) error
}

type Manager struct {
}

type Config struct {
}

func (c *Config) Filename() string {
	return "job.toml"
}

func New() *Manager {
	//config := &Config{}
	//utilpaladin.MustUnmarshalTOML(config)
	return &Manager{
	}
}

func (m *Manager) Create(
	job *ent.Job,
	closer closer,
	client *http.Client,
	email *mail.Manager,
	logger logger,
) *JobExecutor {
	return &JobExecutor{
		job: job,
		closer: closer,
		httpClient: client,
		email: email,
		logger: logger,
	}
}

func (j  *JobExecutor) Validate() error {
	spec := j.job.Spec
	if _, err := ecron.ParseStandard(spec); err != nil {
		return errors.Wrap(ecode.InvalidSpec, "job spec is incorrect！")
	}
	return nil
}

const errMsg = `定时任务出错通知
	您创建的定时任务 (id: %d, name: %s) 执行出错!!!
	您的任务已被禁用！请联系相关人员进行处理`

func (j *JobExecutor) Run() {
	j.logger(context.TODO(), fmt.Sprintf("任务id(%d)开始执行", j.job.ID), j.job.Creator)
	var (
		err error
		request *http.Request
	)
	defer func() {
		if err != nil {
			j.job.Retry--
			if j.job.Retry == 0 {
				log.Info("close the job:", j.job)
				if err = j.closer(context.TODO(), j.job.ID, int(JobStopping)); err != nil {
					log.Error("method: Run#ent_ex/job err:",err)
				}
			}
			m := email.NewEmail(j.email.Host, j.email.Username, j.email.Password, j.email.Port)
			m.WithInfo("定时任务失败报警",fmt.Sprintf(errMsg, j.job.ID, j.job.Name), []string{"306698601@qq.com"})
			if err := m.Send(); err != nil {
				log.Error("send mail notice err:", err)
			}
			return
		}
		j.job.Retry = j.job.RetryTemp
	}()
	request, err = http.NewRequest(j.job.Method, j.job.URL, strings.NewReader(j.job.Body))
	if err != nil {
		log.Error("NewRequest err:", err)
		return
	}
	_, err = j.httpClient.Do(request)
	if err != nil {
		log.Error("do request err:", err)
		return
	}
	j.job.Count--
	if j.job.Count == 0 {
		if err = j.closer(context.TODO(), j.job.ID, int(JobStopping)); err != nil {
			log.Error("method: Run#ent_ex/job err:",err)
		}
	}
	return
}

func (j *JobExecutor) EntryID() ecron.EntryID{
	return ecron.EntryID(j.job.ID)
}

func (j *JobExecutor) GetSpec() string{
	return j.job.Spec
}

type ToDoList struct {
	Name   string     `json:"name"`
	DoTime time.Time  `json:"do_time"`
}
