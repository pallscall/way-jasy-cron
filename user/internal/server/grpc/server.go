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

// New new a grpc server.
func New(svc *service.Service) (ws *warden.Server, err error) {
	var (
		cfg warden.ServerConfig
		ct paladin.TOML
	)
	utilerr.Check(paladin.Get("grpc.toml").Unmarshal(&ct))
	utilerr.Check(ct.Get("Server").UnmarshalTOML(&cfg))
	ws = warden.NewServer(&cfg)
	pb.RegisterTokenVerifyServer(ws.Server(), &Server{Svc: svc})
	ws, err = ws.Start()
	return
}

type Server struct {
	Svc *service.Service
}

func (s *Server)Verify(ctx context.Context, req *pb.VerifyReq)(*pb.VerifyReply, error)  {
	code, err := s.Svc.VerifyToken(ctx, req.AccessKey, req.Token)
	var reply *pb.VerifyReply
	reply = &pb.VerifyReply{
		Code: code,
	}
	//reply := &pb.VerifyReply{Code: code}
	return reply, err
}

