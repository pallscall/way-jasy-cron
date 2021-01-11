package middleware

import (
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"way-jasy-cron/common/ecode"
)

const _session = "access_key"
func Verify(ctx *bm.Context){
	s := ctx.Request.Header.Get(_session)
	if s == "" || s == "null"{
		ctx.JSON(nil, ecode.Unauthorized)
		ctx.Abort()
	}
	return
}
