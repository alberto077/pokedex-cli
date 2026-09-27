package main

import (
	"github.com/alberto077/pokedexcli/configuration"
)


func init() {

	configuration.ExitSignal = make(chan struct{})

	configuration.CommandCallers = []string{"exit", "help"}

	configuration.SharedState = configuration.Config{
		AllCliCommands: map[string]configuration.CliCommand{
			"exit": {
				Name:        "exit",
				Description: "Exit the Pokedex",
				Command:     commandExit,
			},
			"help": {
				Name:        "help",
				Description: "Displays a help message",
				Command:     commandHelp,
				},
			"map": {
				Name: "map",
				Description: "Displays a list of locations and moves foward on every request",
				Command: commandMap,

			},
			},
			
		LocationRequestState: configuration.LocationData{
			Offset: 0,
			Page_limit: 20, // default state needs to be changable for tests
		},

		
	}

}


func main() {

	startRepl(&configuration.SharedState)
}
