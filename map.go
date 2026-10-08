package main

import (
	"fmt"

	"github.com/alberto077/pokedexcli/internal/api"
	"github.com/alberto077/pokedexcli/internal/pokecache"
)

func commandMap(a *appConfig, cache *pokecache.Cache) error {

	formattedRequest := fmt.Sprintf("https://pokeapi.co/api/v2/location-area?offset=%d&limit=%d", a.LocationRequestState.Offset, a.LocationRequestState.Page_limit)
	res, err := api.Locations_Request(formattedRequest, cache)

	if err != nil {
		fmt.Print("Error getting the map location")
		return err
	}

	for _, location := range res.Results {
		fmt.Println(location.Name)

	}

	a.LocationRequestState.Offset += a.LocationRequestState.Page_limit

	return nil

}

func commandMapB(a *appConfig, cache *pokecache.Cache) error {
	if offset, page_limit := a.LocationRequestState.Offset, a.LocationRequestState.Page_limit; offset-page_limit*2 < 0 {
		fmt.Println("Warning: At first page already")
		return nil
	}

	a.LocationRequestState.Offset -= a.LocationRequestState.Page_limit * 2

	formattedRequest := fmt.Sprintf("https://pokeapi.co/api/v2/location-area?offset=%d&limit=%d", a.LocationRequestState.Offset, a.LocationRequestState.Page_limit)
	res, err := api.Locations_Request(formattedRequest, cache)

	if err != nil {
		fmt.Print("Error getting the map location")
		return err
	}

	for _, location := range res.Results {
		fmt.Println(location.Name)

	}

	return nil

}
