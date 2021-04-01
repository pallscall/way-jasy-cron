package service

import (
	"context"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	pb "way-jasy-cron/cron/api"
	pbu "way-jasy-cron/user/userapi"
)

func (s *Service) VerifyToken(ctx context.Context, req *pb.VerifyReq) (reply *pb.VerifyReply, err error) {
	cfg := &warden.ClientConfig{}
	paladin.Get("grpc.toml").UnmarshalTOML(cfg)

	var userClient pbu.TokenVerifyClient
	if userClient, err = pbu.NewClient(cfg); err != nil {
		panic(err)
	}

	v, err := userClient.Verify(ctx, (*pbu.VerifyReq)(req))
	if err != nil {
		return &pb.VerifyReply{Code: -1}, err
	}
	reply = (*pb.VerifyReply)(v)
	return
}
