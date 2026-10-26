package api

import (
	"encoding/json"
	"fmt"
	"github.com/alberto077/pokedexcli/internal/config"
	"github.com/alberto077/pokedexcli/internal/pokecache"
	"io"
)


func GetAreaEncounterResponse(appConifg *config.AppConfig, cache *pokecache.Cache, area string) ([]config.PokemonBaseInfo, error) {
	var encounterResponse config.AreaEncounterResponse

	// side note if we intialize a capcity to 1- and size to - will appends jsut write to the next non intiialized space but what if the space if purposelly allocated already with a 0 value how does append handle this is it looked for non inserted positions
	// The append reads the length data member of the slice to determine where the next value will go along side that is the standard capcacity check to know if the we need to allocate more memory to the underlying array
	var pokemonSeen []config.PokemonBaseInfo = make([]config.PokemonBaseInfo, 0, 10)
	formattedRequest := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", area)

	cachedRes, ok := cache.ThreadSafeGet(formattedRequest)
	if !ok {
		config.LogAction("Cache miss for area request")
		res, err := appConifg.Client.Get(formattedRequest)
		if err != nil {
			return pokemonSeen, fmt.Errorf("Could not fetch area information, %w", err)
		}
		defer res.Body.Close()

		cacheBuffer, err := io.ReadAll(res.Body)
		if err != nil {
			return pokemonSeen, fmt.Errorf("Could not read location response body, %w", err)
		}

		cache.Add(formattedRequest, cacheBuffer)
		cachedRes, ok = cache.ThreadSafeGet(formattedRequest)
		if !ok {
			return pokemonSeen, fmt.Errorf("Could not cache result properly, %w", err)
		}

	}

	if err := json.Unmarshal(cachedRes, &encounterResponse); err != nil {
		return pokemonSeen, fmt.Errorf("Could not Unmarshal raw bytes into pokemon struct, %w", err)
	}

	for _, encounter := range encounterResponse.PokemonEncounters {
		pokemonSeen = append(pokemonSeen, encounter.Pokemon)
	}

	return pokemonSeen, nil

}
