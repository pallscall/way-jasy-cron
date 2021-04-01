package http

import (
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"way-jasy-cron/common/ecode"
	pb "way-jasy-cron/cron/api"
)

const(
	token = "token"
	session = "access_key"
)

func Verify(ctx *bm.Context) {
	s := ctx.Request.Header.Get(session)
	t := ctx.Request.Header.Get(token)
	if t == "" || t == "null" {
		ctx.JSON(nil, ecode.Unauthorized)
		ctx.Abort()
		return
	}
	v, err := svc.VerifyToken(ctx, &pb.VerifyReq{AccessKey: s, Token: t})
	if err != nil {
		ctx.JSON(nil, err)
		ctx.Abort()
		return
	}
	if v.Code != 0 {
		ctx.JSON(nil, ecode.Unauthorized)
		ctx.Abort()
	}
	return
}
