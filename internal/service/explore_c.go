package service

import (
	"fmt"

	"github.com/alberto077/pokedexcli/internal/api"
	"github.com/alberto077/pokedexcli/internal/config"
	"github.com/alberto077/pokedexcli/internal/pokecache"
)

type ExploreLocation struct{}

func (el *ExploreLocation) Execute(appConfig *config.AppConfig, cache *pokecache.Cache) error {
	// fmt.Println("Please Enter the location that you want to explore")

	// appConfig.Scanner.Scan()

	if len(appConfig.CliInput) < 2 {
		fmt.Println("Please pass in the area name after explore refer to the help guide if unclear")
		return fmt.Errorf("Area to explore not passed in")
	}

	var locationName string = appConfig.CliInput[1]

	res, err := api.GetAreaEncounterResponse(appConfig, cache, locationName)
	if err != nil {
		return fmt.Errorf("Could not get area information, %w", err)
	}

	fmt.Printf("Exploring %s ...\n", locationName)

	if len(res) < 1 {
		fmt.Printf("No Pokemon found in %s\n", locationName)
		return nil

	}

	fmt.Printf("Pokemon in %s:\n", locationName)

	for _, pokemonEncounter := range res {
		fmt.Printf(" - %s\n", pokemonEncounter.Name)
	}

	return nil
}
