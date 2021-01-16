package grpc

import (
	"context"
	utilerr "way-jasy-cron/common/util/err"
	pb "way-jasy-cron/user/userapi"

	"way-jasy-cron/user/internal/service"

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
	pb.RegisterUserServer(ws.Server(), &server{svc: svc})
	ws, _ = ws.Start()
	return
}

type server struct {
	svc *service.Service
}

//Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
//SayHello(context.Context, *HelloReq) (*emptypb.Empty, error)
//SayHelloURL(context.Context, *HelloReq) (*HelloResp, error)
func (s *server) LoginUrl(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	return s.svc.LoginUrl(ctx, req)
}

