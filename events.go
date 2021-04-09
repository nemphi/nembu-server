package nembu

import "time"

type EventType string

const (
	EventUserLogin           EventType = "user_login"
	EventUserLogout          EventType = "user_logout"
	EventResourceViewed      EventType = "resource_viewed"
	EventResourceCreated     EventType = "resource_created"
	EventResourceUpdated     EventType = "resource_updated"
	EventResourceDeleted     EventType = "resource_deleted"
	EventResourceSoftDeleted EventType = "resource_soft_deleted"
)

type EventListener func(Event, *Server)

type Event struct {
	Type      EventType
	Meta      map[string]interface{}
	CreatedAt time.Time
}

func (sv *Server) RegisterEventListener(etype EventType, f EventListener) {
	Ilisteners, ok := sv.eventMapper.Load(etype)
	listeners := Ilisteners.([]EventListener)
	if !ok {
		listeners = []EventListener{}
	}
	listeners = append(listeners, f)
	sv.eventMapper.Store(etype, listeners)
}

func (sv *Server) EmitEvent(event Event) {
	if sv.status != StatusStarted {
		return
	}
	sv.events <- event
}

func (sv *Server) listenToEvents() {
	for event := range sv.events {
		Ilisteners, ok := sv.eventMapper.Load(event.Type)
		if !ok {
			continue
		}
		listeners := Ilisteners.([]EventListener)
		go func(listeners []EventListener, sv *Server, event Event) {
			for _, listener := range listeners {
				go listener(event, sv)
			}
		}(listeners, sv, event)
	}
}
