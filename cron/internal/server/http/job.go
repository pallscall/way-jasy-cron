package http

import (
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"github.com/go-kratos/kratos/pkg/net/http/blademaster/binding"
	"graduate/cron/internal/model/ent"
	"graduate/cron/internal/model/ent_ex"
)

func queryJob(ctx *bm.Context) {
	var req struct{
		ID int `form:"id"`
	}
	if err := ctx.BindWith(&req, binding.Query); err != nil {
		return
	}
	ctx.JSON(svc.QueryJob(ctx, req.ID))
}

func listJob(ctx *bm.Context) {
	req := &ent_ex.ListJobReq{}
	if err := ctx.Bind(req); err != nil {
		return
	}
	ctx.JSON(svc.ListJob(ctx, req))
}

func createJob(ctx *bm.Context) {
	req := &ent.Job{}
	if err := ctx.BindWith(req, binding.JSON); err != nil {
		return
	}
	ctx.JSON(nil, svc.CreateJob(ctx, req))
}

func updateJob(ctx *bm.Context) {
	req := &ent.Job{}
	if err := ctx.BindWith(req, binding.JSON); err != nil {
		return
	}
	ctx.JSON(nil, svc.UpdateJob(ctx, req))
}

func switchJobStatus(ctx *bm.Context) {
	var req struct{
		ID int `form:"id"`
		Opt int `form:"opt"`
	}
	if err := ctx.BindWith(&req, binding.Query); err != nil {
		return
	}
	ctx.JSON(nil, svc.SwitchJobStatus(ctx, req.ID, req.Opt))
}

