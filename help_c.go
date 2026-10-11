package main

import (
	"fmt"

	"github.com/alberto077/pokedexcli/internal/pokecache"
)

func commandHelp(appConfig *appConfig, cache *pokecache.Cache) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, caller := range appConfig.CommandCallers {

		callerInfo, ok := appConfig.AllCliCommands[caller]

		if !ok {
			return fmt.Errorf("Could not find command %s in command list", caller)
		}

		fmt.Printf("%s: %s\n", callerInfo.Name, callerInfo.Description)
	}

	return nil

}
