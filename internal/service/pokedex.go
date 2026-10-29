package service

import (
	"fmt"
	"github.com/alberto077/pokedexcli/internal/config"
	"github.com/alberto077/pokedexcli/internal/pokecache"
)

type OpenPokedex struct{}

func (pdx *OpenPokedex) Execute(appConfig *config.AppConfig, cache *pokecache.Cache) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range appConfig.Pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
