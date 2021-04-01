package ent_ex

import "way-jasy-cron/cron-logger/internal/model/ent"

type ListLoggerReq struct {
	ListBaseReq
	Operator    string `form:"operator"`
}

type ListLoggerOptions struct {
	ListBaseOptions
	Operator    string `form:"operator"`
}

type ListLoggerResp struct {
	ListBaseResp
	Logs []*ent.Logger `json:"logs"`
}

func (r ListLoggerReq) ToListLoggerOptions() (o *ListLoggerOptions) {
	o = &ListLoggerOptions{
		ListBaseOptions: ListBaseOptions{
			PN: r.PN,
			PS: r.PS,
		},
		Operator: r.Operator,
	}
	o.Completed()
	return
}