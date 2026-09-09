package event

type EventHandler struct {
	channel chan Event
	event   Event
}

func NewEventHandler(event Event, eventBus *EventBus) *EventHandler {
	eh := &EventHandler{
		channel: make(chan Event),
		event:   event,
	}

	eventBus.Subscribe(event, eh)

	go eh.listen()

	return eh
}

func (eh *EventHandler) listen() {
	for e := range eh.channel {
		eh.event.Callback(e)
	}
}

func (eh *EventHandler) Channel() chan Event {
	return eh.channel
}

func (eh *EventHandler) Topic() string {
	return eh.event.topic
}
