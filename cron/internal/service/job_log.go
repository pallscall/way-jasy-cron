package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	pbl "way-jasy-cron/cron-logger/api"
	pb "way-jasy-cron/cron/api"
)

func (s *Service) WriteLog(ctx context.Context, req *pb.WriteLogReq) (reply *pb.NoReply, err error) {
	message := req.Ctime.String() + req.Opt
	fmt.Printf("%s\n", message)

	cfg := &warden.ClientConfig{}
	paladin.Get("grpc.toml").UnmarshalTOML(cfg)

	var cronLoggerClient pbl.CronLoggerClient
	if cronLoggerClient,err = pbl.NewClient(cfg); err != nil {
		panic(err)
	}

	r, err := cronLoggerClient.WriteLog(ctx, (*pbl.WriteLogReq)(req))
	reply = (*pb.NoReply)(r)

	return
}

