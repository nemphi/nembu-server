package database

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/nemphi/nembu-server/config"
	"github.com/nemphi/nembu-server/models"
)

type Connection struct {
	pgdb *pg.DB
}

func New(cfg *config.Config) (*Connection, error) {
	conn := &Connection{}
	conn.pgdb = pg.Connect(&pg.Options{
		Addr:     cfg.DB.Host + ":" + cfg.DB.Port,
		User:     cfg.DB.Username,
		Password: cfg.DB.Password,
		Database: cfg.DB.Database,
	})

	ctx := context.Background()
	if err := conn.pgdb.Ping(ctx); err != nil {
		return nil, err
	}

	return conn, nil
}

func (conn *Connection) CreateTables() error {
	return conn.createTables(models.AllModels())
}

func (conn *Connection) createTables(models []interface{}) error {
	for _, model := range models {
		err := conn.pgdb.Model(model).CreateTable(&orm.CreateTableOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}
