package service

import (
	"fmt"

	"github.com/alberto077/pokedexcli/internal/api"
	"github.com/alberto077/pokedexcli/internal/config"
	"github.com/alberto077/pokedexcli/internal/pokecache"
)

type ExploreLocation struct{}

func (el *ExploreLocation) Execute(appConfig *config.AppConfig, cache *pokecache.Cache) error {
	fmt.Println("Please Enter the location that you want to explore")

	appConfig.Scanner.Scan()
	var locationName string = appConfig.Scanner.Text()

	res, err := api.GetAreaEncounterResponse(cache, locationName)
	if err != nil {
		return fmt.Errorf("Could not get area information, %w", err)
	}

	if len(res) < 1 {
		fmt.Printf("No Pokemon found in %s\n", locationName)
		return nil

	}

	fmt.Printf("Pokemon in %s:\n", locationName)

	for _, pokemonEncounter := range res {
		fmt.Printf("%s\n", pokemonEncounter.Name)
	}

	return nil
}
