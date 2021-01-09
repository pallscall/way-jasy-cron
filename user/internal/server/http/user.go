package http

import (
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"graduate/user/internal/model/ent"
)

func register(ctx *bm.Context) {
	req := &ent.User{}
	if err := ctx.Bind(req); err != nil {
		return
	}
	ctx.JSON(nil, svc.Register(ctx, req))
}
