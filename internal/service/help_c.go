package service

import (
	"fmt"

	"github.com/alberto077/pokedexcli/internal/config"
	"github.com/alberto077/pokedexcli/internal/pokecache"
)

type Help struct{}

func (h *Help) Execute(appConfig *config.AppConfig, cache *pokecache.Cache) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, caller := range appConfig.CommandCallers {

		_, ok := appConfig.CommandRegistry.AllCliCommands[caller.Name]

		if !ok {
			return fmt.Errorf("Could not find command %s in command list", caller)
		}

		fmt.Printf("%s: %s\n", caller.Name, caller.Description)
	}

	return nil

}
