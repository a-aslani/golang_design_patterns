package command

type PasteCommand struct {
	Command
}

func (p PasteCommand) Execute() error {
	p.saveBackup()
	p.editor.replaceSelection(p.app.clipboard)
	return nil
}
