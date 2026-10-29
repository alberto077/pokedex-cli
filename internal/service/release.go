package service

import (
	"fmt"

	"github.com/alberto077/pokedexcli/internal/config"
	"github.com/alberto077/pokedexcli/internal/pokecache"
)

type Release struct{}

func (r *Release) Execute(appConfig *config.AppConfig, cache *pokecache.Cache) error {
	if len(appConfig.CliInput) < 2 {
		fmt.Println("Specify the pokemon that you want to release")
		return fmt.Errorf("Pokemon for removal not selected")
	}

	pokemonForRemoval := appConfig.CliInput[1]
	fmt.Printf("Removing %s...\n", pokemonForRemoval)
	delete(appConfig.Pokedex, pokemonForRemoval)
	fmt.Printf("%s was removed\n", pokemonForRemoval)
	return nil

}
