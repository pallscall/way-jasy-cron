// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"way-jasy-cron/cron-logger/internal/dao"
	"way-jasy-cron/cron-logger/internal/service"
	"way-jasy-cron/cron-logger/internal/server/grpc"
	"way-jasy-cron/cron-logger/internal/server/http"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
