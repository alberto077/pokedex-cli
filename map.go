package main

import (
	"fmt"
	"github.com/alberto077/pokedexcli/configuration"
	"github.com/alberto077/pokedexcli/internal/api"
)



func commandMap(sharedState *configuration.Config) {
	res, err := api.Locations_Request(sharedState)

	if err != nil {
		fmt.Print("Error getting the map location")
	}


	for _, location := range(res.Results) {
		fmt.Println(location.Name)
		
	}

}
