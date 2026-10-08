package main

import "github.com/alberto077/pokedexcli/internal/pokecache"

type CliCommand struct {
	Name        string
	Description string
	Command     func(*appConfig, *pokecache.Cache) error
}

type LocationData struct {
	Offset     int16
	Page_limit int16
}

type appConfig struct {
	ExitSignal           chan struct{}
	CommandCallers       []string
	AllCliCommands       map[string]CliCommand
	LocationRequestState LocationData
}

func GetCommands() *appConfig {

	cfg := appConfig{

		ExitSignal: make(chan struct{}),

		CommandCallers: []string{"exit", "help"},

		AllCliCommands: map[string]CliCommand{
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
				Name:        "map",
				Description: "Displays a list of locations and moves foward on every request",
				Command:     commandMap,
			},
			"mapb": {
				Name:        "mapb",
				Description: "Displays a list of locations and moves backwards on every request",
				Command:     commandMapB,
			},
		},

		LocationRequestState: LocationData{
			Offset:     0,
			Page_limit: 20, // default state needs to be changable for tests
		},
	}

	return &cfg

}
