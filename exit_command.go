package main

import (
	"fmt"
	"github.com/alberto077/pokedexcli/configuration"
)

func commandExit(sharedState *configuration.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	close(configuration.ExitSignal)
	return nil
}
