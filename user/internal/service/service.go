package service

import (
	"context"
	"sync"
	"way-jasy-cron/user/internal/dao/ent"
	"way-jasy-cron/user/internal/dao/mail"
	"way-jasy-cron/user/internal/dao/mysql"
	"way-jasy-cron/user/internal/dao/redis"
	//"github.com/google/wire"
)

//var Provider = wire.NewSet(New, wire.Bind(new(pb.DemoServer), new(*Service)))

// Service service.
type Service struct {
	ent         *ent.Manager
	mysql		*mysql.Manager
	redis       *redis.Manager
	mail        *mail.Manager
	stop        chan bool
	wg          sync.WaitGroup
}

// New new a service and return.
func New() *Service{
	return &Service{
		ent:     ent.New(),
		mysql:   mysql.New(),
		redis:   redis.New(),
		mail:    mail.New(),
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

