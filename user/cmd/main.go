package main

import (
	"flag"
	"way-jasy-cron/user/internal/server/grpc"
	"way-jasy-cron/user/internal/server/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"way-jasy-cron/user/internal/service"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/log"
	//"github.com/go-kratos/kratos/tool/kratos/cron/internal/di"
)

func main() {
	flag.Parse()
	log.Init(nil) // debug flag: log.dir={path}
	defer log.Close()
	log.Info("account system start")
	paladin.Init()
	//_, closeFunc, err := di.InitApp()
	//if err != nil {
	//	panic(err)
	//}
	svc := service.New()
	http.MustStart()
	grpc.MustStart(svc)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Info("account system exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
