package mail

import (
	utilpaladin "way-jasy-cron/common/util/paladin"
)

type Manager struct {
	Host 		string
	Port 		int
	Username 	string
	Password 	string
}

type Config struct {
	Host 		string
	Port 		int
	Username 	string
	Password 	string
}

func (c *Config) Filename() string {
	return "mail.toml"
}

func New() *Manager{
	c := &Config{}
	utilpaladin.MustUnmarshalTOML(c)
	return &Manager{
		Host: c.Host,
		Port: c.Port,
		Username: c.Username,
		Password: c.Password,
	}
}