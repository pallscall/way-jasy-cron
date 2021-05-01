package http

import (
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
)

func show(ctx *bm.Context){
	ctx.JSON(nil, svc.Show(ctx))
}
