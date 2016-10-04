package event

type Handler struct {
	EventType string
	f         func(e *Event) string
}

type Event interface{}
