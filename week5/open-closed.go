type Command interface {
	Execute() ([]byte, error) ValidateInput() bool
}
type CommandExecutor struct{}

func (c CommandExecutor) Execute(command Command) { 
	if command.ValidateInput() {
		command.Execute()
	}
type FooCommand struct {
args []string // needed args
}

func (c FooCommand) ValidateInput() bool { 
	// validate args
	return len(args) >= 1 && len(args[0]) > 0
}
func (c FooCommand) Execute() ([]byte, error) { 
	// logic for FooCommand execution
	 return nil, nil
}

type BarCommand struct {
}
func (c BarCommand) ValidateInput() {
	 // does nothing
     return
}
func (c BarCommand) Execute() ([]byte, error) { // logic for BarCommand execution 
	return nil, nil
}

// Let’s take the example of the CommandExecutor, which is responsible 
// for executing Commands. The Execute() and ValidateInput() methods
//  need to handle each command separately. So every time a new command is added Execute()
//  implementation needs to change.

