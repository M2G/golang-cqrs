package facade

import (
	"golang-cqrs/command"
	"golang-cqrs/event"
	"golang-cqrs/query"
)

type Facade struct {
	CommandBus *command.CommandBus
	QueryBus   *query.QueryBus
	EventBus   *event.EventBus
}

func NewFacade(commandHandlers map[string]*command.CommandHandler, queryHandlers map[string]query.Handler, eventBus *event.EventBus) *Facade {
	facade := &Facade{
		CommandBus: command.NewCommandBus(),
		QueryBus:   query.NewQueryBus(),
		EventBus:   eventBus,
	}

	for name, handler := range commandHandlers {
		facade.CommandBus.RegisterHandler(name, handler)
	}
	for name, handler := range queryHandlers {
		facade.QueryBus.AddHandler(name, handler)
	}

	return facade
}

func (f *Facade) Dispatch(cmd *command.Command) error {
	return f.CommandBus.Execute(cmd)
}

func (f *Facade) Ask(q *query.Query) (interface{}, error) {
	return f.QueryBus.Execute(q)
}
