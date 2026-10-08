package api

import (
	"encoding/json"
	"fmt"
	"github.com/alberto077/pokedexcli/internal/pokecache"
	"io"
	"net/http"
)

type NamedAPIResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Locations struct {
	Count    int                `json:"count"`
	Next     string             `json:"next"`
	Previous string             `json:"previous"`
	Results  []NamedAPIResource `json:"results"`
}

func Locations_Request(formattedRequest string, cache *pokecache.Cache) (Locations, error) {

	var locationRes Locations
	var cacheBuffer []byte

	res, ok := cache.ThreadSafeGet(formattedRequest)
	if !ok {
		// cache miss so make the network request
		apiResponse, err := http.Get(formattedRequest)
		if err != nil {
			return locationRes, fmt.Errorf("Could not make request: %w", err)
		}

		defer apiResponse.Body.Close()

		cacheBuffer, err = io.ReadAll(apiResponse.Body)
		if err != nil {
			return locationRes, fmt.Errorf("Error reading api response, %w", err)
		}

		cache.Add(formattedRequest, cacheBuffer)
		res, ok = cache.ThreadSafeGet(formattedRequest)
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
