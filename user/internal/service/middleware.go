package service

import (
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	log "github.com/sirupsen/logrus"
	"graduate/user/ecode"
	"graduate/user/internal/model/ent"
)

func (svc *Service) Verify(ctx *bm.Context, accessKey string) error{
	user, err := svc.ent.QueryUser(ctx, accessKey)
	if err != nil {
		log.Error("QueryUser err:(%v)", err)
		return err
	}
	if user == nil {
		return ecode.InvalidUser
	}
	req := &ent.User{}
	if err = ctx.Bind(req); err != nil {
		return err
	}
	if req.Password != user.Password {
		return ecode.InvalidPassword
	}
	return nil
}
