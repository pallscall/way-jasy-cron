package http

import (
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"github.com/go-kratos/kratos/pkg/net/http/blademaster/binding"
)

func getNextCronJobList(ctx *bm.Context) {
	ctx.JSON(svc.GetNextCronJobList(ctx))
}

func getJobCount(ctx *bm.Context) {
	var req struct{
		Creator string `json:"creator" form:"creator"`
	}
	if err := ctx.BindWith(&req, binding.Query); err != nil {
		return
	}
	ctx.JSON(svc.GetJobCount(ctx, req.Creator))
}
