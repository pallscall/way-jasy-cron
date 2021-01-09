package service

import (
	"context"
	"way-jasy-cron/user/internal/dao/ent"
	"way-jasy-cron/user/internal/dao/mysql"
	"sync"
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

func (svc *Service) InitJob() {
	//svc.ent.
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
