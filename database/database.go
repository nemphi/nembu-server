package database

import (
	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"github.com/nemphi/nembu-server/config"
	"google.golang.org/grpc"
)

type Connection struct {
	*dgo.Dgraph
}

func New(cfg *config.Config) (*Connection, error) {
	conn := &Connection{}
	options := []grpc.DialOption{}
	if cfg.DB.Insecure {
		options = append(options, grpc.WithInsecure())
	}
	gconn, err := grpc.Dial(cfg.DB.Host+":"+cfg.DB.Port, options...)
	if err != nil {
		return nil, err
	}

	conn.Dgraph = dgo.NewDgraphClient(api.NewDgraphClient(gconn))

	return conn, nil
}
