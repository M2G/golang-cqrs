package command

import "uuid"

type Command struct {
	CommandId   string
	CommandName string
	Args        []interface{}
}

func NewCommand(commandName string, args ...interface{}) *Command {
	return &Command{
		CommandId:   uuid.New().String(),
		CommandName: commandName,
		Args:        args,
	}
}
