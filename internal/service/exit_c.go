package service

import (
	"fmt"

	"github.com/alberto077/pokedexcli/internal/config"
	"github.com/alberto077/pokedexcli/internal/pokecache"
)

type Exit struct{}

func (e *Exit) Execute(AppConfig *config.AppConfig, cache *pokecache.Cache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	close(AppConfig.ExitSignal)
	return nil
}
