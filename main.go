package main


type cliCommand struct {
	name        string
	description string
	command     func() error
}

var (
	commandCallers []string // ordered callers
	allCliCommands map[string]cliCommand
	exitSignal chan struct{}
)




func init() {

	exitSignal = make(chan struct{})

	commandCallers = []string{"exit", "help"}

	allCliCommands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			command:     commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			command:     commandHelp,
		},
	}

}

func main() {
	startRepl()
}
