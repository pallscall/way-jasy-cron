package mysql

import (
	"context"
	"github.com/go-kratos/kratos/pkg/database/sql"
	"github.com/go-kratos/kratos/pkg/log"
	utilpaladin "graduate/common/util/paladin"
)

const _wildcards = "%"

type Manager struct {
	db         *sql.DB
	maxRoutine int
}

type Config struct {
	Mysql      *sql.Config
	MaxRoutine int
}

func (c *Config) Filename() string {
	return "db.toml"
}

func New() *Manager {
	c := &Config{}
	utilpaladin.MustUnmarshalTOML(c)
	return &Manager{
		db:         sql.NewMySQL(c.Mysql),
		maxRoutine: c.MaxRoutine,
	}
}

func (m *Manager) Close(ctx context.Context) {
	done := make(chan bool, 1)
	go func() {
		m.db.Close()
		done <- true
	}()
	select {
	case <-ctx.Done():
		log.Error("close db timeout error: %v", ctx.Err())
	case <-done:
		log.Info("mysql client closed")
	}
}

func (m *Manager) Ping(ctx context.Context) error {
	return m.db.Ping(ctx)
}

func fuzzyColumn(k string) string {
	return _wildcards + k + _wildcards
}

func IsErrNoRows(err error) bool {
	return err == sql.ErrNoRows
}
