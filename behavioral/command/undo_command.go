package command

type UndoCommand struct {
	Command
}

func (u UndoCommand) Execute() error {
	u.app.undo()

	return nil
}
