package api


import (
	"encoding/json"
	"fmt"
	"github.com/alberto077/pokedexcli/internal/config"
	"github.com/alberto077/pokedexcli/internal/pokecache"
	"io"
)

func PokemonInfoExtractor(appConfig *config.AppConfig, cache *pokecache.Cache, pokemonName string) (config.Pokemon, error){
	
	var pokemonResponse config.Pokemon

	var cacheBuffer[]byte
	formattedRequest := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonName)

	cachedRes, ok := cache.ThreadSafeGet(formattedRequest)
	if !ok {

		config.LogAction("Cache miss for Pokemon")
		resp, err := appConfig.Client.Get(formattedRequest)

		if err != nil {
			return pokemonResponse, fmt.Errorf("Request for pokemon info failed")
		}

		defer resp.Body.Close()

		cacheBuffer, err = io.ReadAll(resp.Body)

		if err != nil {
			return pokemonResponse, fmt.Errorf("Could not read Pokemon Info Response")
		}

		cache.Add(formattedRequest, cacheBuffer)
		cachedRes, ok = cache.ThreadSafeGet(formattedRequest)
		if !ok {
			return pokemonResponse, fmt.Errorf("Cache fault")
		}

	}
	if err := json.Unmarshal(cachedRes, &pokemonResponse); err != nil {
		return pokemonResponse, fmt.Errorf("Could not Unmarshal pokemon json response")
	}

	return pokemonResponse, nil
}
