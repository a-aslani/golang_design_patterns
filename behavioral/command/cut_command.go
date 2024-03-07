package command

type CutCommand struct {
	Command
}

func (c CutCommand) Execute() error {
	c.saveBackup()
	c.app.clipboard = c.editor.getSelection()
	c.editor.deleteSelection()
	return nil
}
