package command

import "fmt"

type CommandBus struct {
	handlers map[string]*CommandHandler
	commands map[string]*Command
}

func (cb *CommandBus) RegisterHandler(commandName string, handler *CommandHandler) {
	cb.handlers[commandName] = handler
}

func (cb *CommandBus) Execute(command *Command) error {
	handler := cb.handlers[command.CommandName]
	if handler == nil {
		return fmt.Errorf("handler for command %s not found", command.CommandName)
	}
	handler.Execute(command.Args)
	return nil
}
