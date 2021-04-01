package http

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	"sync"
	"way-jasy-cron/common/middleware"
	utilerr "way-jasy-cron/common/util/err"
	pb "way-jasy-cron/cron-logger/api"
	"way-jasy-cron/cron-logger/internal/server/grpc"
	"way-jasy-cron/cron-logger/internal/service"
)

var (
	svc              *service.Service
	e                *bm.Engine
	onlyOnceShutdown = make(chan struct{})
)

// New new a bm server.
func New(s *service.Service) {
	var c struct {
		Server *bm.ServerConfig
		Auth   *warden.ClientConfig
	}
	if err := paladin.Get("http.toml").UnmarshalTOML(&c); err != nil {
		return
	}
	e = bm.DefaultServer(c.Server)
	var wg sync.WaitGroup
	// must use goroutine here, because initRouter recovered the panic error
	// it don't make sense here.
	go func() {
		wg.Add(1)
		defer wg.Done()
		svc = s
		initRouter(e)
		utilerr.Check(e.Start())
		pb.RegisterCronLoggerBMServer(e, &grpc.Service{Svc: s})
	}()
	wg.Wait()
	return
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	g := e.Group("/logger",middleware.Verify)
	{
		g.GET("", listLog)
	}
}

func ping(ctx *bm.Context) {
	//if _, err := svc.Ping(ctx, nil); err != nil {
	//	log.Error("ping error(%v)", err)
	//	ctx.AbortWithStatus(http.StatusServiceUnavailable)
	//}
}

