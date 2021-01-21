package service

import (
	"context"
	"net/http"
	"sync"
	"way-jasy-cron/common/ecron"
	"way-jasy-cron/cron/internal/dao/ent"
	"way-jasy-cron/cron/internal/dao/mail"
	"way-jasy-cron/cron/internal/dao/mysql"
	"way-jasy-cron/cron/internal/model/ent_ex"

	//"github.com/google/wire"
)

//var Provider = wire.NewSet(New, wire.Bind(new(pb.DemoServer), new(*Service)))

// Service service.
type Service struct {
	ent         *ent.Manager
	mysql		*mysql.Manager
	mail        *mail.Manager
	cron        *ecron.Cron
	http        *http.Client
	jobex       *ent_ex.Manager
	stop        chan bool
	wg          sync.WaitGroup
}

// New new a service and return.
func New() *Service{
	return &Service{
		ent:     ent.New(),
		cron:    ecron.New(),
		mysql:   mysql.New(),
		jobex:   ent_ex.New(),
		http:    http.DefaultClient,
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
