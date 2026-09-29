package main

import (
	"fmt"
	"github.com/alberto077/pokedexcli/configuration"
	"github.com/alberto077/pokedexcli/internal/api"
)



func commandMap(sharedState *configuration.Config) error {
	if sharedState.LocationRequestState.Offset != 0 {
		sharedState.LocationRequestState.Offset += sharedState.LocationRequestState.Page_limit

	}

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

func commandMapB(sharedState *configuration.Config) error {

	if offset, page_limit := sharedState.LocationRequestState.Offset, sharedState.LocationRequestState.Page_limit; offset - page_limit * 2< 0 {
		fmt.Println("Warning: At first page already")
		return nil
	}

	sharedState.LocationRequestState.Offset -= sharedState.LocationRequestState.Page_limit * 2

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
