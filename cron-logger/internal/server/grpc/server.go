package grpc

import (
	"context"
	pb "way-jasy-cron/cron-logger/api"
	"way-jasy-cron/cron-logger/internal/service"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
)

// New new a grpc server.
func New(svc *service.Service) (ws *warden.Server, err error) {
	var (
		cfg warden.ServerConfig
		ct paladin.TOML
	)
	if err = paladin.Get("grpc.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Server").UnmarshalTOML(&cfg); err != nil {
		return
	}
	ws = warden.NewServer(&cfg)
	pb.RegisterCronLoggerServer(ws.Server(), &Service{svc})
	ws, err = ws.Start()
	return
}

type Service struct {
	Svc *service.Service
}

func (s *Service) WriteLog(ctx context.Context, req *pb.WriteLogReq) (*pb.NoReply, error) {
	return s.Svc.WriteLog(ctx, req)
}




