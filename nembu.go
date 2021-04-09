package nembu

import (
	"sync"

	"github.com/nemphi/nembu-server/config"
	"github.com/nemphi/nembu-server/database"
)

type ServerStatus byte

const (
	StatusStopped ServerStatus = iota
	StatusStarting
	StatusStarted
	StatusStopping
)

type Server struct {
	cfg         *config.Config
	db          *database.Connection
	events      chan Event
	eventMapper *sync.Map
	status      ServerStatus
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
	sv.status = StatusStarting
	sv.listenToEvents()
	sv.status = StatusStarted
}

func (sv *Server) Stop() {
	sv.status = StatusStopping
	close(sv.events)
	sv.status = StatusStopped
}
