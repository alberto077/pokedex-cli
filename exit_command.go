package main

import (
	"fmt"

	"github.com/alberto077/pokedexcli/internal/pokecache"
)

func commandExit(a *appConfig, cache *pokecache.Cache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	close(a.ExitSignal)
	return nil
}
