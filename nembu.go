package nembu

import (
	"sync"

	"github.com/nemphi/nembu-server/config"
	"github.com/nemphi/nembu-server/database"
)

type Server struct {
	cfg         *config.Config
	db          *database.Connection
	events      chan Event
	eventMapper *sync.Map
	started     bool
}

func New(options ...Option) (*Server, error) {
	sv := &Server{
		events:      make(chan Event),
		eventMapper: &sync.Map{},
	}
	for _, option := range options {
		err := option(sv)
		if err != nil {
			return nil, err
		}
	}
	return sv, nil
}

func (sv *Server) Start() {
	sv.listenToEvents()
	sv.started = true
}

func (sv *Server) Stop() {
	close(sv.events)
	sv.started = false
}
