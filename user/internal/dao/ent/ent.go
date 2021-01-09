package ent

import (
	utilerr "graduate/common/util/err"
	utilpaladin "graduate/common/util/paladin"
	"graduate/user/internal/model/ent"
)

type Manager struct {
	Client *ent.Client
}

type Config struct {
	Driver  string
	Source  string
}

func (c *Config) Filename() string {
	return "ent.toml"
}

func New() *Manager {
	c := &Config{}
	utilpaladin.MustUnmarshalTOML(c)
	client, err := ent.Open(c.Driver, c.Source)
	utilerr.Check(err)
	return &Manager{
		Client:  client,
	}
}
