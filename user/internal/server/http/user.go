package http

import (
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"way-jasy-cron/user/internal/model/ent"
)

func register(ctx *bm.Context) {
	req := &ent.User{}
	if err := ctx.Bind(req); err != nil {
		return
	}
	ctx.JSON(nil, svc.Register(ctx, req))
}

func generateRSA(ctx *bm.Context) {
	s := ctx.Request.Header.Get(_openSession)
	ctx.JSON(svc.GenerateRSA(ctx, s))
}

func getPublicKey(ctx *bm.Context) {
	s := ctx.Request.Header.Get(_openSession)
	ctx.JSON(svc.GetPublicKey(ctx, s))
}
