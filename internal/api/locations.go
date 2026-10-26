package api

import (
	"encoding/json"
	"fmt"
	"github.com/alberto077/pokedexcli/internal/config"
	"github.com/alberto077/pokedexcli/internal/pokecache"
	"io"
)

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationResponse struct {
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Location `json:"results"`
}

func Locations_Request(appConfig *config.AppConfig, cache *pokecache.Cache, url string) (LocationResponse, error) {

	var locationRes LocationResponse
	var cacheBuffer []byte

	res, ok := cache.ThreadSafeGet(url)
	if !ok {
		// cache miss so make the network request

		config.LogAction(fmt.Sprintf("cache miss for locations request, location url: %s", url))
		apiResponse, err := appConfig.Client.Get(url)
		if err != nil {
			return locationRes, fmt.Errorf("Could not make request: %w", err)
		}

		defer apiResponse.Body.Close()

		cacheBuffer, err = io.ReadAll(apiResponse.Body)
		if err != nil {
			return locationRes, fmt.Errorf("Error reading api response, %w", err)
		}

		cache.Add(url, cacheBuffer)
		res, ok = cache.ThreadSafeGet(url)
		if !ok {
			return locationRes, fmt.Errorf("Failed to get cache request")
		}

	}

	// just read from the cache o(1) operation
	if err := json.Unmarshal(res, &locationRes); err != nil {
		return locationRes, fmt.Errorf("Could not unmarshal cache read, %w\n cache read: %v", err, res)
	}

	return locationRes, nil

}
