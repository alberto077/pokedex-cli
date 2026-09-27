package main


type cliCommand struct {
	name        string
	description string
	command     func(*config) error
}

type locationData struct{
	offset uint16
	page_limit uint16

}

type config struct {
	allCliCommands map[string]cliCommand
	locationRequestState locationData
}

var (
	commandCallers []string // ordered callers
	allCliCommands map[string]cliCommand
	exitSignal chan struct{}
	sharedState config
)




func init() {

	exitSignal = make(chan struct{})

	commandCallers = []string{"exit", "help"}

	sharedState = config{
		allCliCommands: map[string]cliCommand{
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
			},
		locationRequestState: locationData{
			offset: 0,
			page_limit: 20, // default state needs to be changable for tests
		},

		
	}

}


func main() {

	startRepl(&sharedState)
}
