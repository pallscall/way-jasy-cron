package ent_ex

import (
	"github.com/pkg/errors"
	"graduate/common/ecron"
	"graduate/cron/ecode"
	"graduate/cron/internal/model/ent"
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

type Job ent.Job
func (j Job) Validate() error {
	spec := j.Spec
	if _, err := ecron.ParseStandard(spec); err != nil {
		return errors.Wrap(ecode.InvalidSpec, "job spec is incorrect！")
	}
	return nil
}

func (j *Job) Run() {

}

func (j *Job) EntryID() ecron.EntryID{
	return ecron.EntryID(j.ID)
}

func (j *Job) GetSpec() string{
	return j.Spec
}