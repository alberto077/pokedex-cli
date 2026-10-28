package service

import (
	"fmt"

	"github.com/alberto077/pokedexcli/internal/api"
	"github.com/alberto077/pokedexcli/internal/config"
	"github.com/alberto077/pokedexcli/internal/pokecache"
	"math/rand"
)

type Catch struct{}

func (c *Catch) Execute(appconifg *config.AppConfig, cache *pokecache.Cache) error {

	if len(appconifg.CliInput) < 2 {
		fmt.Println("Make sure to input the pokemon name that you want to catch")
		return fmt.Errorf("No name passed in")
	}
	pokemonName := appconifg.CliInput[1]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemonInfo, err := api.PokemonInfoExtractor(appconifg, cache, pokemonName)
	if err != nil {
		return err
	}

	caught, err := ThrowPokeball(pokemonInfo)

	if !caught {
		fmt.Printf("%s escaped!\n", pokemonName)
	} else {
		fmt.Printf("%s was caught!\n", pokemonName)
		appconifg.Pokedex[pokemonName] = pokemonInfo
	}

	return nil
}

func ThrowPokeball(pokemon config.Pokemon) (bool, error) {

	num := rand.Intn(pokemon.BaseExperience)

	if num < 25 {
		return true, nil
	}

	return false, nil
}
