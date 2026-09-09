package command

type CommandHandler struct {
	Command func(args []interface{})
}

func NewHandler(command func(args []interface{})) *CommandHandler {
	return &CommandHandler{
		Command: command,
	}
}

func (h *CommandHandler) Execute(args []interface{}) {
	h.Command(args)
}
