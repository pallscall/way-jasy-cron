package redis

import (
	"github.com/go-kratos/kratos/pkg/cache/redis"
	utilpaladin "way-jasy-cron/common/util/paladin"
)

type Manager struct {
	Redis *redis.Pool
}

type Config struct {
	Redis *redis.Config
}

func (c *Config) Filename() string {
	return "redis.toml"
}

func New() *Manager{
	c := &Config{}
	utilpaladin.MustUnmarshalTOML(c)
	return &Manager{
		Redis: redis.NewPool(c.Redis),
	}
}