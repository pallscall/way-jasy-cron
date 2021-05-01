package ent_ex

import (
	"way-jasy-cron/cron/internal/model/ent"
)

type ShellJobStatus int

const (
	ShellJobDelete   ShellJobStatus = -1
	ShellJobRunning  ShellJobStatus = 1
	ShellJobStopping ShellJobStatus = 0
)

type ListShellJobReq struct {
	ListBaseReq
	Creator string `form:"creator"`
	Host    string `form:"host"`
	Comment string `form:"comment"`
}

type ListShellJobOptions struct {
	ListBaseOptions
	Creator string
	Host    string
	Comment string
}

type ListShellJobResp struct {
	ListBaseResp
	Jobs []*ent.Machine `json:"jobs"`
}

func (r ListShellJobReq) ToListShellJobOptions() (o *ListShellJobOptions) {
	o = &ListShellJobOptions{
		ListBaseOptions: ListBaseOptions{
			PN: r.PN,
			PS: r.PS,
		},
		Creator: r.Creator,
		Host:    r.Host,
		Comment: r.Comment,
	}
	o.Completed()
	return
}

