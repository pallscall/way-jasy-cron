package service

import (
	"context"
	"github.com/go-kratos/kratos/pkg/log"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	uuid "github.com/satori/go.uuid"
	xecode "way-jasy-cron/common/ecode"
	"way-jasy-cron/user/ecode"
	"way-jasy-cron/user/internal/model/ent"
)

func (svc *Service) Verify(ctx *bm.Context, accessKey string) (string,error){
	//if do, err := svc.redis.Redis.Do(ctx, "exists", accessKey)
	user, err := svc.ent.QueryUser(ctx, accessKey)
	if err != nil {
		log.Error("QueryUser err:(%v)", err)
		return "", err
	}
	if user == nil {
		return "", ecode.InvalidUser
	}
	req := &ent.User{}
	if err = ctx.Bind(req); err != nil {
		return "", err
	}
	if err := svc.DecodePwd(ctx, req); err != nil {
		log.Error("method: Verify#middleware decode pwd err:", err)
		return "", err
	}
	if req.Password != user.Password {
		return "", ecode.InvalidPassword
	}
	return svc.setToken(ctx, accessKey)

}

func (svc *Service) setToken(ctx context.Context, accessKey string) (string,error){
	token := GetUUID()
	r := svc.redis.Redis.Get(ctx)
	defer r.Close()
	if _, err := r.Do("set", accessKey, token); err != nil {
		log.Error("set token err(%v)", err)
		return "", err
	}
	if _, err := r.Do("expire", accessKey, 86400); err != nil {
		log.Error("expire token err(%v)", err)
		return "",err
	}
	return token, nil
}

func GetUUID() string {
	u := uuid.NewV4()
	return u.String()
}

const (
	tokenVerifyPass = 0
	tokenVerifyUnPass = -1
)
func (svc *Service) VerifyToken(ctx context.Context, accessKey, token string) (code int32, err error){
	r := svc.redis.Redis.Get(ctx)
	defer r.Close()
	var (
		val interface{}
	)
	if val, err = r.Do("get", accessKey); err != nil {
		log.Error("get accessKey err(%v)",err)
		return tokenVerifyUnPass, err
	}

	if val == nil || string(val.([]uint8)) != token {
		return tokenVerifyUnPass, xecode.Unauthorized
	}

	return tokenVerifyPass, nil
}