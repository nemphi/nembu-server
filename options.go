package nembu

import (
	"github.com/nemphi/nembu-server/config"
	"github.com/nemphi/nembu-server/database"
)

type Option func(*Server) error

func WithConfig(cfg *config.Config) Option {
	return func(sv *Server) error {
		sv.cfg = cfg
		return nil
	}
}

func WithDBConnection(db *database.Connection) Option {
	return func(sv *Server) error {
		sv.db = db
		return nil
	}
}
