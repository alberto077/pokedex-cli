package main

import (
	"fmt"
	"github.com/alberto077/pokedexcli/configuration"
)

func commandHelp(sharedState *configuration.Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _ , caller := range configuration.CommandCallers {

		callerInfo, ok := sharedState.AllCliCommands[caller]

		if !ok {
			return fmt.Errorf("Could not find command %s in command list", caller)
		}
		
		fmt.Printf("%s: %s\n", callerInfo.Name, callerInfo.Description)
	}

	return nil

}
