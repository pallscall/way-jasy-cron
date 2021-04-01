// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package dao

import (
	"github.com/google/wire"
	"way-jasy-cron/cron-logger/internal/dao/mysql"
)

//go:generate kratos tool wire
func newTestDao() (*dao, func(), error) {
	panic(wire.Build(newDao, mysql.NewDB, NewRedis, NewMC))
}
