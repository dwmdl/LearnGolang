package stat

import (
	"api/pkg/event"
	"log"
)

type ServiceDeps struct {
	EventBus   *event.Bus
	Repository *Repository
}

type Service struct {
	EventBus   *event.Bus
	Repository *Repository
}

func NewService(deps *ServiceDeps) *Service {
	return &Service{
		EventBus:   deps.EventBus,
		Repository: deps.Repository,
	}
}

func (s *Service) AddDirection() {
	for msg := range s.EventBus.Subscribe() {
		if msg.Type == event.LinkVisited {
			linkId, ok := msg.Data.(uint)
			if !ok {
				log.Println("Bad EventLinkVisited data: ", msg.Data)
				continue
			}
			s.Repository.AddDirection(linkId)
		}
	}
}
