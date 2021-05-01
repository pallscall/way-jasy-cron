package http

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	"net/http"
	"sync"
	"way-jasy-cron/common/middleware"
	utilerr "way-jasy-cron/common/util/err"
	"way-jasy-cron/cron/internal/service"

	"github.com/go-kratos/kratos/pkg/log"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
)

var (
	svc              *service.Service
	e                *bm.Engine
	onlyOnceShutdown = make(chan struct{})
)

//var svc pb.DemoServer
//
//// New new a bm server.
//func New(s pb.DemoServer) (engine *bm.Engine, err error) {
//	var (
//		cfg bm.ServerConfig
//		ct paladin.TOML
//	)
//	if err = paladin.Get("http.toml").Unmarshal(&ct); err != nil {
//		return
//	}
//	if err = ct.Get("Server").UnmarshalTOML(&cfg); err != nil {
//		return
//	}
//	svc = s
//	engine = bm.DefaultServer(&cfg)
//	pb.RegisterDemoBMServer(engine, s)
//	initRouter(engine)
//	err = engine.Start()
//	return
//}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	g := e.Group("/cron",middleware.Verify, Verify)
	{
		h := g.Group("/http")
		{
			h.GET("", listJob)
			h.GET("/query", queryJob)
			h.POST("/create", createJob)
			h.POST("/update", updateJob)
			h.POST("/switch", switchJobStatus)
		}

		s := g.Group("/shell")
		{
			s.GET("", listShellJob)
			s.GET("/query", queryShellJob)
			s.POST("/create", createShellJob)
			s.POST("/update", updateShellJob)
			s.POST("/switch", switchShellJobStatus)
			s.POST("/ping", connHost)
		}

		d := g.Group("dashboard")
		{
			d.GET("/todo",getNextCronJobList)
			d.GET("/count",getJobCount)
		}
	}
}

func ping(ctx *bm.Context) {
	if err := svc.Ping(ctx); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

func MustStart(s *service.Service) {
	svc = s
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
		initRouter(e)
		utilerr.Check(e.Start())
		utilerr.Check(svc.InitJob())
	}()
	wg.Wait()
}

