package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	"sync"
	pb2 "way-jasy-cron/cron/api"
	pb "way-jasy-cron/user/userapi"
	"way-jasy-cron/user/internal/dao/ent"
	"way-jasy-cron/user/internal/dao/mysql"
	//"github.com/google/wire"
)

//var Provider = wire.NewSet(New, wire.Bind(new(pb.DemoServer), new(*Service)))

// Service service.
type Service struct {
	ent         *ent.Manager
	mysql		*mysql.Manager
	stop        chan bool
	wg          sync.WaitGroup
}

// New new a service and return.
func New() *Service{
	return &Service{
		ent:     ent.New(),
		mysql:   mysql.New(),
		stop:    make(chan bool, 1),
	}

}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) error {
	return s.mysql.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close(ctx context.Context) {
	s.stop <- true
	if s.mysql != nil {
		s.mysql.Close(ctx)
	}
	if s.ent != nil {
		s.ent.Client.Close()
	}
	s.wg.Wait()
}

func (s *Service) LoginUrl(ctx context.Context, req *pb.LoginReq) (reply *pb.LoginResp, err error) {
	fmt.Printf("server1 login username: %s, passwd: %s", req.Username, req.Passwd)

	cfg := &warden.ClientConfig{}
	paladin.Get("grpc.toml").UnmarshalTOML(cfg)

	var demoClient pb2.JobClient
	if demoClient,err = pb2.NewClient(cfg); err != nil {
		panic(err)
	}

	reply2, err := demoClient.Login(ctx, (*pb2.LoginReq)(req))
	reply = (*pb.LoginResp)(reply2)

	return
}