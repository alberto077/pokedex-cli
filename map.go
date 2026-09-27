package main

import (
	"fmt"
	"github.com/alberto077/pokedexcli/configuration"
	"github.com/alberto077/pokedexcli/internal/api"
)



func commandMap(sharedState *configuration.Config) error {
	res, err := api.Locations_Request(sharedState)

	if err != nil {
		fmt.Print("Error getting the map location")
		return err
	}


	for _, location := range(res.Results) {
		fmt.Println(location.Name)
		
	}

	return nil

}
