package api

import (
	"encoding/json"
	"fmt"
	"github.com/alberto077/pokedexcli/internal/pokecache"
	"io"
	"net/http"
)

type AreaEncounterResponse struct {
	EncounterMethodRates []EncounterMethodRates `json:"encounter_method_rates"`
	PokemonEncounters    []PokemonEncounters    `json:"pokemon_encounters"`
}
type EncounterMethod struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type EncounterMethodRates struct {
	EncounterMethod EncounterMethod `json:"encounter_method"`
}
type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Method struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type EncounterDetails struct {
	MinLevel        int    `json:"min_level"`
	MaxLevel        int    `json:"max_level"`
	ConditionValues []any  `json:"condition_values"`
	Chance          int    `json:"chance"`
	Method          Method `json:"method"`
}
type PokemonEncounters struct {
	Pokemon Pokemon `json:"pokemon"`
}

func GetAreaEncounterResponse(c *pokecache.Cache, area string) ([]Pokemon, error) {
	var encounterResponse AreaEncounterResponse

	// side note if we intialize a capcity to 1- and size to - will appends jsut write to the next non intiialized space but what if the space if purposelly allocated already with a 0 value how does append handle this is it looked for non inserted positions
	// The append reads the length data member of the slice to determine where the next value will go along side that is the standard capcacity check to know if the we need to allocate more memory to the underlying array
	var pokemonSeen []Pokemon = make([]Pokemon, 0, 10)
	formattedRequest := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", area)

	cachedRes, ok := c.ThreadSafeGet(formattedRequest)
	if !ok {
		res, err := http.Get(formattedRequest)
		if err != nil {
			return pokemonSeen, fmt.Errorf("Could not fetch area information, %w", err)
		}
		defer res.Body.Close()

		cacheBuffer, err := io.ReadAll(res.Body)
		if err != nil {
			return pokemonSeen, fmt.Errorf("Could not read location response body, %w", err)
		}

		c.Add(formattedRequest, cacheBuffer)
		cachedRes, ok = c.ThreadSafeGet(formattedRequest)
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
