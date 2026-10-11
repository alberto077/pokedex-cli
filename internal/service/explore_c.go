package service

import (
	"fmt"

	"github.com/alberto077/pokedexcli/internal/api"
	"github.com/alberto077/pokedexcli/internal/config"
	"github.com/alberto077/pokedexcli/internal/pokecache"
)

type ExploreLocation struct{}

func (el *ExploreLocation) Execute(appConfig *config.AppConfig, cache *pokecache.Cache) error {
	var locationName string = appConfig.CliInput[1]

	res, err := api.GetAreaEncounterResponse(cache, locationName)
	if err != nil {
		return fmt.Errorf("Could not get area information, %w", err)
	}

	fmt.Println("Pokemon in area:")

	for _, pokemonEncounter := range res {
		fmt.Printf("%s\n", pokemonEncounter.Name)
	}

	return nil
}
