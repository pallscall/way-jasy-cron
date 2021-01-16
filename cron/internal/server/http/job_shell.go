package http

import (
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"github.com/go-kratos/kratos/pkg/net/http/blademaster/binding"
	"way-jasy-cron/cron/internal/model/ent"
	"way-jasy-cron/cron/internal/model/ent_ex"
)

func queryShellJob(ctx *bm.Context) {
	var req struct{
		ID int `form:"id"`
	}
	if err := ctx.BindWith(&req, binding.Query); err != nil {
		return
	}
	ctx.JSON(svc.QueryJob(ctx, req.ID))
}

func listShellJob(ctx *bm.Context) {
	req := &ent_ex.ListShellJobReq{}
	if err := ctx.Bind(req); err != nil {
		return
	}
	ctx.JSON(svc.ListShellJob(ctx, req))
}

func createShellJob(ctx *bm.Context) {
	req := &ent.Machine{}
	if err := ctx.BindWith(req, binding.JSON); err != nil {
		return
	}
	ctx.JSON(nil, svc.CreateShellJob(ctx, req))
}

func updateShellJob(ctx *bm.Context) {
	req := &ent.Machine{}
	if err := ctx.BindWith(req, binding.JSON); err != nil {
		return
	}
	ctx.JSON(nil, svc.UpdateShellJob(ctx, req))
}

func switchShellJobStatus(ctx *bm.Context) {
	var req struct{
		ID int `form:"id"`
		Opt int `form:"opt"`
	}
	if err := ctx.BindWith(&req, binding.Query); err != nil {
		return
	}
	ctx.JSON(nil, svc.SwitchShellJobStatus(ctx, req.ID, req.Opt))
}

