package main

import (
	"fmt"

	"github.com/alberto077/pokedexcli/internal/config"
	"github.com/alberto077/pokedexcli/internal/pokecache"
	"github.com/alberto077/pokedexcli/internal/service"
)

func startRepl(appConfig *config.AppConfig, cache *pokecache.Cache) {

	fmt.Println("Welcome to the pokedex CLI")

	for {
		select {
		case <-appConfig.ExitSignal:
			return
		default:

			fmt.Print("Pokedex > ")
			appConfig.Scanner.Scan()
			if err := appConfig.Scanner.Err(); err != nil {
				fmt.Printf("Scanner eskror: %v\n", err)
				return
			}
			// could possible make a service functions that handles cli input and cleans it to seperate service functions completly from the repl loop
			appConfig.CliInput = service.CleanInput(appConfig.Scanner.Text())

			command, ok := appConfig.CommandRegistry.AllCliCommands[appConfig.CliInput[0]]

			if !ok {
				fmt.Println("Invalid command, use the help command to discover what you can do")
			} else {
				command.Execute(appConfig, cache)
			}
		}

	}

}
