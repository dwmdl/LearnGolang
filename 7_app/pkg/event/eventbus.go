package event

const (
	LinkVisited = "link.visited"
)

type Event struct {
	Type string
	Data any
}

type Bus struct {
	main chan Event
}

func NewEventBus() *Bus {
	return &Bus{
		main: make(chan Event),
	}
}

func (b *Bus) Publish(event Event) {
	b.main <- event
}

func (b *Bus) Subscribe() <-chan Event {
	return b.main
}
