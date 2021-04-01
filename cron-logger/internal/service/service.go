package service

import (
	"context"
	"sync"
	pb "way-jasy-cron/cron-logger/api"
	"way-jasy-cron/cron-logger/internal/dao/ent"

	"github.com/google/wire"
)

var Provider = wire.NewSet(New, wire.Bind(new(pb.CronLoggerServer), new(*Service)))

// Service service.
type Service struct {
	//mysql *mysql.Manager
	ent         *ent.Manager
	stop        chan bool
	wg          sync.WaitGroup
}

// New new a service and return.
func New() *Service {
	return &Service{
		//mysql:  mysql.New(),
		ent:  ent.New(),
		stop: make(chan bool, 1),
	}
}


//// Ping ping the resource.
//func (s *Service) Ping(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
//	return &empty.Empty{}, s.dao.Ping(ctx)
//}

// Close close the resource.
func (s *Service) Close(ctx context.Context) {
	s.stop <- true
	//if s.mysql != nil {
	//	s.mysql.Close(ctx)
	//}
	if s.ent != nil {
		s.ent.Client.Close()
	}
	s.wg.Wait()
}
