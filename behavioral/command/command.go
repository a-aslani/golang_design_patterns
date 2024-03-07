package command

type Commander interface {
	Execute() error
}

type Command struct {
	app    Application
	editor Editor
	backup string
}

func NewCommand(app Application, editor Editor) *Command {
	return &Command{
		app:    app,
		editor: editor,
	}
}

func (c *Command) saveBackup() {
	c.backup = c.editor.text
}

func (c *Command) undo() {
	c.editor.text = c.backup
}

type Application struct {
	clipboard    string
	editors      []Editor
	activeEditor Editor
	history      CmdHistory
}

func (a *Application) createUI() {

}

func (a *Application) executeCommand(cmd Commander) {
	if err := cmd.Execute(); err != nil {
		a.history.push(cmd)
	}
}

func (a *Application) undo() {
	cmd := a.history.pop()
	if cmd != nil {
	}
}

type CmdHistory struct {
	history []Commander
}

func NewCmdHistory() *CmdHistory {
	return &CmdHistory{
		history: make([]Commander, 0),
	}
}

func (cmd *CmdHistory) push(c Commander) {
	cmd.history = append(cmd.history, c)
}

func (cmd *CmdHistory) pop() Commander {
	return cmd.history[len(cmd.history)-1]
}

type Editor struct {
	text string
}

func NewEditor(text string) Editor {
	return Editor{text: text}
}

func (e Editor) getSelection() string {
	// Return selected text.
	return ""
}

func (e Editor) deleteSelection() {
	// Delete selected text.
}

func (e Editor) replaceSelection(s string) {
	// Insert the clipboard's contents at the current
	// position.
}
