package query

import "fmt"

type QueryBus struct {
	handlers map[string]Handler
}

func NewQueryBus() *QueryBus {
	return &QueryBus{handlers: make(map[string]Handler)}
}

func (qb *QueryBus) GetHandler(queryName string) Handler {
	return qb.handlers[queryName]
}

func (qb *QueryBus) AddHandler(queryName string, handler Handler) {
	if _, exists := qb.handlers[queryName]; exists {
		panic(fmt.Errorf("cannot override handler for query: %s", queryName))
	}
	qb.handlers[queryName] = handler
}

func (qb *QueryBus) Execute(q *Query) (interface{}, error) {
	handler := qb.GetHandler(q.QueryName)
	if handler == nil {
		return nil, fmt.Errorf("no handler found for query: %s", q.QueryName)
	}
	return handler.Execute(q.Args)
}
