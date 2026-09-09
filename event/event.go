package event

type Event struct {
	topic    string
	Callback func(data interface{})
}

func NewEvent(topic string, callback func(data interface{})) *Event {
	return &Event{
		topic:    topic,
		Callback: callback,
	}
}
