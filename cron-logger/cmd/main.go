package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"
	"way-jasy-cron/cron-logger/internal/server/grpc"
	"way-jasy-cron/cron-logger/internal/server/http"
	"way-jasy-cron/cron-logger/internal/service"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/log"
)

func main() {
	flag.Parse()
	log.Init(nil) // debug flag: log.dir={path}
	defer log.Close()
	log.Info("cron-logger start")
	paladin.Init()

	svc := service.New()
	ws, err := grpc.New(svc)
	http.New(svc)
	if err != nil {
		panic(err)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ws.Shutdown(context.TODO())
			log.Info("cron-logger exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
