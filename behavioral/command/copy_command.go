package command

type CopyCommand struct {
	Command
}

func (c CopyCommand) Execute() error {
	c.app.clipboard = c.editor.getSelection()
	return nil
}
