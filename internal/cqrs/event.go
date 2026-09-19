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
}

func Subscribe[E any](bus *EventBus, handler func(e E)) {
}

func Publish(event interface{}) {}
