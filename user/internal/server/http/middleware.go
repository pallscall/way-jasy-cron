package http

import (
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	xecode "way-jasy-cron/user/ecode"
)

const (
	_session = "access_key"
	_openSession = "wj_cron"
)

func openVerify(ctx *bm.Context) {
	s := ctx.Request.Header.Get(_openSession)
	if s == "" || s == "null"{
		ctx.JSON(nil, xecode.InLegalOpt)
		ctx.Abort()
	}
	return
}

func verifyLogin(ctx *bm.Context) {
	s := ctx.Request.Header.Get(_openSession)
	var (
		token string
		err error
	)
	if token, err = svc.Verify(ctx, s); err != nil {
		ctx.JSON(nil, err)
		ctx.Abort()
	}
	ctx.JSON(token, nil)
}