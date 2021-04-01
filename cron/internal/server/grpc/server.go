package grpc

import (
	"context"
	utilerr "way-jasy-cron/common/util/err"
	pb "way-jasy-cron/cron/api"
	"way-jasy-cron/cron/internal/service"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
)

var (
	ws *warden.Server
)
func MustStart(svc *service.Service) {
	ws = New(svc)
}
// New new a grpc server.
func New(svc *service.Service) (ws *warden.Server) {
	var (
		cfg warden.ServerConfig
		ct paladin.TOML
	)
	utilerr.Check(paladin.Get("grpc.toml").Unmarshal(&ct))
	utilerr.Check(ct.Get("Server").UnmarshalTOML(&cfg))
	ws = warden.NewServer(&cfg)
	pb.RegisterJobServer(ws.Server(), &server{svc: svc})
	ws, _ = ws.Start()
	return
}

type server struct {
	svc *service.Service
}

//Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
//SayHello(context.Context, *HelloReq) (*emptypb.Empty, error)
//SayHelloURL(context.Context, *HelloReq) (*HelloResp, error)
func (s *server) WriteLog(ctx context.Context, req *pb.WriteLogReq) (*pb.NoReply, error) {
	return s.svc.WriteLog(ctx, req)
}

func (s *server) Verify(ctx context.Context, req *pb.VerifyReq) (*pb.VerifyReply, error) {
	return s.svc.VerifyToken(ctx, req)
}

