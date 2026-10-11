package service

import (
	"fmt"

	"github.com/alberto077/pokedexcli/internal/api"
	"github.com/alberto077/pokedexcli/internal/config"
	"github.com/alberto077/pokedexcli/internal/pokecache"
)

type Map struct{}

func (m *Map) Execute(appConfig *config.AppConfig, cache *pokecache.Cache) error {

	res, err := api.Locations_Request(appConfig, cache)

	if err != nil {
		fmt.Print("Error getting the map location")
		return err
	}

	for _, location := range res {
		fmt.Println(location.Name)

	}

	appConfig.LocationRequestState.Offset += appConfig.LocationRequestState.Page_limit

	return nil

}

type MapB struct{}

func (mb *MapB) Execute(appConfig *config.AppConfig, cache *pokecache.Cache) error {
	if offset, page_limit := appConfig.LocationRequestState.Offset, appConfig.LocationRequestState.Page_limit; offset-page_limit*2 < 0 {
		fmt.Println("Warning: At first page already")
		return nil
	}

	appConfig.LocationRequestState.Offset -= appConfig.LocationRequestState.Page_limit * 2

	res, err := api.Locations_Request(appConfig, cache)
	if err != nil {
		fmt.Print("Error getting the map location")
		return err
	}

	for _, location := range res {
		fmt.Println(location.Name)

	}

	return nil

}
