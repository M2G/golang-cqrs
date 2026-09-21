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
	var zero E
	t := reflect.TypeOf(zero)

	bus.mu.Lock()
	defer bus.mu.Unlock()
	bus.handlers[t] = append(bus.handlers[t], func(e any) { handler(e.(E)) })
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
