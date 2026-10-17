package service

import (
	"fmt"

	"github.com/alberto077/pokedexcli/internal/api"
	"github.com/alberto077/pokedexcli/internal/config"
	"github.com/alberto077/pokedexcli/internal/pokecache"
)

type Map struct{}

const baseLocationUrl string = "https://pokeapi.co/api/v2/location-area"

func (m *Map) Execute(appConfig *config.AppConfig, cache *pokecache.Cache) error {
	var nextRequest string
	if appConfig.LocationRequestState.NextPage == "" {
		nextRequest = baseLocationUrl
		fmt.Println("First Page")
	

	} else {
		nextRequest = appConfig.LocationRequestState.NextPage
	}
	res, err := api.Locations_Request(appConfig, cache, nextRequest)

	if err != nil {
		fmt.Print("Error getting the map location")
		return err
	}

	appConfig.LocationRequestState.PreviousPage = res.Previous
	appConfig.LocationRequestState.NextPage = res.Next

	for _, location := range res.Results {
		fmt.Println(location.Name)

	}

	return nil

}

type MapB struct{}

func (mb *MapB) Execute(appConfig *config.AppConfig, cache *pokecache.Cache) error {
	var prevRequest string
	if appConfig.LocationRequestState.NextPage == "" {
		prevRequest = baseLocationUrl
		fmt.Println("First page")

	} else {
		prevRequest = appConfig.LocationRequestState.PreviousPage
	}

	res, err := api.Locations_Request(appConfig, cache, prevRequest)
	if err != nil {
		fmt.Print("Error getting the map location")
		return err
	}

	appConfig.LocationRequestState.PreviousPage = res.Previous
	appConfig.LocationRequestState.NextPage = res.Next
	for _, location := range res.Results {
		fmt.Println(location.Name)

	}

	return nil

}
