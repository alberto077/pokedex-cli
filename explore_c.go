package main

import (
	"fmt"

	"github.com/alberto077/pokedexcli/internal/pokecache"
	"github.com/alberto077/pokedexcli/internal/api"
)


func commandExploreLocation(appConfig *appConfig, cache *pokecache.Cache) error {
	var locationName string = appConfig.CliInput[1]

	res, err := api.GetAreaEncounterResponse(cache, locationName)
	if err != nil {
		return fmt.Errorf("Could not get area information, %w", err)
	}

	fmt.Println("Pokemon in area:")

	for _, pokemon := range res {
		fmt.Printf("%s\n", pokemon)
	}


	return nil
}
