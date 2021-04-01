package http

import (
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"way-jasy-cron/cron-logger/internal/model/ent_ex"
)

func listLog(ctx *bm.Context) {
	req := &ent_ex.ListLoggerReq{}
	if err := ctx.Bind(req); err != nil {
		return
	}
	ctx.JSON(svc.ListLog(ctx, req))
}