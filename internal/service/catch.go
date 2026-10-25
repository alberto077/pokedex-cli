package service

import (
	"fmt"
	"github.com/alberto077/pokedexcli/internal/config"
	"github.com/alberto077/pokedexcli/internal/pokecache"
)

type Catch struct {}

func (c *Catch) Execute(appconifg *config.AppConfig, cache *pokecache.Cache) error {
	

	if len(appconifg.CliInput) < 2 {
		fmt.Println("Make sure to input the pokemon name that you want to catch")
		return fmt.Errorf("No name passed in")
	}
	pokemon := appconifg.CliInput[1]
	fmt.Printf("Throwing a Pokeball at %s\n", pokemon)

	return nil
}
