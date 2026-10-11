package api

import (
	"encoding/json"
	"fmt"
	"github.com/alberto077/pokedexcli/internal/config"
	"github.com/alberto077/pokedexcli/internal/pokecache"
	"io"
	"net/http"
)

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationResponse struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Location `json:"results"`
}

func Locations_Request(appConfig *config.AppConfig, cache *pokecache.Cache) ([]Location, error) {

	var locationRes LocationResponse
	var cacheBuffer []byte

	formattedRequest := fmt.Sprintf("https://pokeapi.co/api/v2/location-area?offset=%d&limit=%d", appConfig.LocationRequestState.Offset, appConfig.LocationRequestState.Page_limit)

	res, ok := cache.ThreadSafeGet(formattedRequest)
	if !ok {
		// cache miss so make the network request
		apiResponse, err := http.Get(formattedRequest)
		if err != nil {
			return locationRes.Results, fmt.Errorf("Could not make request: %w", err)
		}

		defer apiResponse.Body.Close()

		cacheBuffer, err = io.ReadAll(apiResponse.Body)
		if err != nil {
			return locationRes.Results, fmt.Errorf("Error reading api response, %w", err)
		}

		cache.Add(formattedRequest, cacheBuffer)
		res, ok = cache.ThreadSafeGet(formattedRequest)
		if !ok {
			return locationRes.Results, fmt.Errorf("Failed to get cache request")
		}

	}

	// just read from the cache o(1) operation
	if err := json.Unmarshal(res, &locationRes); err != nil {
		return locationRes.Results, fmt.Errorf("Could not unmarshal cache read, %w\n cache read: %v", err, res)
	}

	return locationRes.Results, nil

}
