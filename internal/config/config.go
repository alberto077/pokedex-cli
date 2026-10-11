package config

import "github.com/alberto077/pokedexcli/internal/pokecache"

type Command interface {
	Execute(appconifg *AppConfig, cache *pokecache.Cache) error
}

type CommandRegistry struct {
	AllCliCommands map[string]Command
}

type CliCommand struct {
	Name        string
	Description string
	Command     Command
}

type LocationData struct {
	Offset     int16
	Page_limit int16
}

type AppConfig struct {
	ExitSignal           chan struct{}
	CommandCallers       []CliCommand
	CommandRegistry      *CommandRegistry
	LocationRequestState LocationData
	CliInput             []string
}

func NewCommandRegistry() *CommandRegistry {
	return &CommandRegistry{
		AllCliCommands: make(map[string]Command),
	}
}

func (cr *CommandRegistry) RegisterCommand(cfg *AppConfig, cliCommand CliCommand) {

	cr.AllCliCommands[cliCommand.Name] = cliCommand.Command
	cfg.CommandCallers = append(cfg.CommandCallers, cliCommand)

}

func InitConfig() *AppConfig {
	var cfg AppConfig
	cfg.ExitSignal = make(chan struct{})
	cfg.CommandCallers = make([]CliCommand, 0)
	cfg.LocationRequestState = LocationData{
		Offset:     0,
		Page_limit: 20, // default state needs to be changable for tests
	}

	return &cfg
}
