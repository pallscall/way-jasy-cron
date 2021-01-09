package http

import (
	"github.com/go-kratos/kratos/pkg/ecode"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
)

const _session = "access_key"

func verify(ctx *bm.Context){
	s := ctx.Request.Header.Get(_session)
	if s == "" {
		ctx.JSON(nil, ecode.Unauthorized)
		ctx.Abort()
		return
	}
	ctx.JSON(nil, nil)
}

func verifyLogin(ctx *bm.Context) {
	s := ctx.Request.Header.Get(_session)
	if err := svc.Verify(ctx, s); err != nil {
		ctx.JSON(nil, err)
		ctx.Abort()
	}
	return
}