package service

import (
	"fmt"

	"github.com/alberto077/pokedexcli/internal/config"
	"github.com/alberto077/pokedexcli/internal/pokecache"
)

type Inspect struct{}

func (i *Inspect) Execute(appConfig *config.AppConfig, cache *pokecache.Cache) error {
	if len(appConfig.CliInput) < 2 {
		fmt.Println("Please pass in the pokemon that you want to inspect")
		return fmt.Errorf("Not enough arguments passed in")
	}

	pokemon := appConfig.CliInput[1]

	pokemonInfo, ok := appConfig.Pokedex[pokemon]

	if !ok {
		fmt.Printf("You have not caught %s no info avalible", pokemon)
		return nil
	}

	var statsCategories []string = []string{"hp", "attack", "defense", "special-attack", "special-defense", "speed"}

	fmt.Printf("Name: %s\n", pokemonInfo.Name)
	fmt.Printf("Height: %d\n", pokemonInfo.Height)
	fmt.Printf("Weight: %d\n", pokemonInfo.Weight)
	fmt.Println("Stats:")
	for i, stat := range pokemonInfo.Stats {
		fmt.Printf(" -%s: %d\n", statsCategories[i], stat.BaseStat)
	}
	fmt.Println("Types")
	for _, types := range pokemonInfo.Types {
		fmt.Printf(" -%s\n", types.Type.Name)
	}
	return nil
}
