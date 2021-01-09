package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"graduate/common/ecron"
	"graduate/cron/internal/dao/ent"
	"graduate/cron/internal/dao/mysql"
	"sync"
	//"github.com/google/wire"
)

//var Provider = wire.NewSet(New, wire.Bind(new(pb.DemoServer), new(*Service)))

// Service service.
type Service struct {
	ent         *ent.Manager
	mysql		*mysql.Manager
	cron        *ecron.Cron
	stop        chan bool
	wg          sync.WaitGroup
}

// New new a service and return.
func New() *Service{
	return &Service{
		ent:     ent.New(),
		cron:    ecron.New(),
		mysql:   mysql.New(),
		stop:    make(chan bool, 1),
	}
}

func (svc *Service) InitJob() error{
	jobs, err := svc.ent.ListAllRunningJobs(context.TODO())
	if err != nil {
		log.Error("InitJob err:", err)
		return err
	}
	for _, job := range jobs {
		if _, err := svc.cron.AddJob(job); err != nil {
			log.Error("AddJob err:", err)
			return err
		}
	}
	return nil
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
