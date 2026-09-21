package cqrs

import (
	"reflect"
	"sync"
)

type EventBus struct {
	mu       sync.RWMutex
	handlers map[reflect.Type][]func(any)
}

func NewEventBus() *EventBus {
	return &EventBus{handlers: make(map[reflect.Type][]func(any))}
}

func Subscribe[E any](bus *EventBus, handler func(e E)) {
	// add bus lock + unlock
}

func Publish[E any](bus *EventBus, e E) {
	t := reflect.TypeOf(e)

	bus.mu.RLock()
	hanlders := bus.handlers[t]
	bus.mu.RUnlock()

	for _, h := range hanlders {
		// go routine
		go h(e)
	}

}
