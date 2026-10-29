package main

import (
	"github.com/alberto077/pokedexcli/internal/config"
	"github.com/alberto077/pokedexcli/internal/pokecache"
	"github.com/alberto077/pokedexcli/internal/service"
	"time"
)

func main() {
	var cfg *config.AppConfig = config.InitConfig()
	var pcache *pokecache.Cache = pokecache.NewCache(30 * time.Second)

	cfg.CommandRegistry = config.NewCommandRegistry()

	cfg.CommandRegistry.RegisterCommand(cfg, config.CliCommand{
		Name:        "exit",
		Description: "Exit the Pokedex",
		Command:     &service.Exit{},
	})

	cfg.CommandRegistry.RegisterCommand(cfg, config.CliCommand{
		Name:        "help",
		Description: "Displays a help message",
		Command:     &service.Help{},
	})

	cfg.CommandRegistry.RegisterCommand(cfg, config.CliCommand{

		Name:        "map",
		Description: "Displays a list of locations and moves foward on every request",
		Command:     &service.Map{},
	})

	cfg.CommandRegistry.RegisterCommand(cfg, config.CliCommand{

		Name:        "mapb",
		Description: "Displays a list of locations and moves backwards on every request",
		Command:     &service.MapB{},
	})

	cfg.CommandRegistry.RegisterCommand(cfg, config.CliCommand{

		Name:        "explore",
		Description: "Expore the location specified after the command which will show pokemon encounterd",
		Command:     &service.ExploreLocation{},
	})

	cfg.CommandRegistry.RegisterCommand(cfg, config.CliCommand{

		Name:        "catch",
		Description: "throw a pokeball for a chance to catch a pokemon",
		Command:     &service.Catch{},
	})

	cfg.CommandRegistry.RegisterCommand(cfg, config.CliCommand{

		Name:        "inspect",
		Description: "Get stats about a pokemon youve caught and stored in your pokedex",
		Command:     &service.Inspect{},
	})

	cfg.CommandRegistry.RegisterCommand(cfg, config.CliCommand{

		Name:        "pokedex",
		Description: "Get a list of all the pokemon you have caught",
		Command:     &service.OpenPokedex{},
	})

	cfg.CommandRegistry.RegisterCommand(cfg, config.CliCommand{

		Name:        "release",
		Description: "Release one of the pokemon you have caught",
		Command:     &service.Release{},
	})
	startRepl(cfg, pcache)
}
