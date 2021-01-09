package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"way-jasy-cron/user/ecode"
	"way-jasy-cron/user/internal/model/ent"
)

func (svc *Service) Register(ctx context.Context, user *ent.User) error{
	u, err := svc.ent.QueryUser(ctx, user.Username)
	if err != nil {
		log.Error("method: Register#service. QueryUser err:", err)
		return err
	}
	if u != nil {
		return ecode.UserExist
	}
	if err := svc.ent.Register(ctx, user); err != nil {
		log.Error("method: Register#service. Register err:", err)
		return err
	}
	log.Info("Register success!")
	return nil
}
