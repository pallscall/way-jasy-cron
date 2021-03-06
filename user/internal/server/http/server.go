package http

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/log"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"github.com/go-kratos/kratos/pkg/net/http/blademaster/binding"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	"net/http"
	"sync"
	utilerr "way-jasy-cron/common/util/err"
	"way-jasy-cron/user/internal/server/grpc"
	"way-jasy-cron/user/internal/service"
	pb "way-jasy-cron/user/userapi"
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
	g := e.Group("/account",openVerify)
	{
		g.POST("/login", verifyLogin)
		g.POST("/register", register)

		g.GET("/pubKey", getPublicKey)
		g.POST("/crypt", generateRSA)

		g.GET("/info", getUserInfo)
	}
	t := e.Group("/test")
	{
		t.POST("", test)
	}
}

func ping(ctx *bm.Context) {
	if err := svc.Ping(ctx); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

func MustStart(s *service.Service) {
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
		svc = s
		pb.RegisterTokenVerifyBMServer(e, &grpc.Server{Svc: s})
	}()
	wg.Wait()
}

func test(ctx *bm.Context) {
	var req struct{
		Rec  string `json:"rec" form:"rec"`
	}
	if err := ctx.BindWith(&req, binding.JSON); err != nil {
		return
	}
	ctx.JSON(nil, svc.TestShow(ctx, req.Rec))
}