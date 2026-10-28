package config

import (
	"bufio"
	"log"
	"net/http"
	"os"

	"github.com/alberto077/pokedexcli/internal/pokecache"
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
type PokemonBaseInfo struct {
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
	Pokemon PokemonBaseInfo `json:"pokemon"`
}

type PokedexConfig struct {
	Pokedex map[string]Pokemon
}

type Pokemon struct {
	ID             int         `json:"id"`
	Name           string      `json:"name"`
	BaseExperience int         `json:"base_experience"`
	Abilities      []Abilities `json:"abilities"`
	Forms          []Forms     `json:"forms"`
	Moves          []Moves     `json:"moves"`
	Stats          []Stats     `json:"stats"`
}
type Ability struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Abilities struct {
	Ability Ability `json:"ability"`
}
type Forms struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Move struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type VersionGroup struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type MoveLearnMethod struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type VersionGroupDetails struct {
	LevelLearnedAt  int             `json:"level_learned_at"`
	VersionGroup    VersionGroup    `json:"version_group"`
	MoveLearnMethod MoveLearnMethod `json:"move_learn_method"`
	Order           int             `json:"order"`
}
type Moves struct {
	Move                Move                  `json:"move"`
	VersionGroupDetails []VersionGroupDetails `json:"version_group_details"`
}
type Stat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Stats struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}
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
	NextPage     string
	PreviousPage string
	Offset       int16
	Page_limit   int16
}

type AppConfig struct {
	ExitSignal           chan struct{}
	CommandCallers       []CliCommand
	CommandRegistry      *CommandRegistry
	LocationRequestState LocationData
	CliInput             []string
	Scanner              *bufio.Scanner
	Client               *http.Client
	Pokedex              map[string]Pokemon
	Logger               *os.File
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

func LogAction(input string) {
	log.Println(input)

}

func InitConfig() *AppConfig {
	var cfg AppConfig
	logFile, err := os.OpenFile("repl.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Could not find log file")
	}
	cfg.ExitSignal = make(chan struct{})
	cfg.CommandCallers = make([]CliCommand, 0)
	cfg.LocationRequestState = LocationData{
		Offset:     0,
		Page_limit: 20, // default state needs to be changable for tests
	}
	cfg.Scanner = bufio.NewScanner(os.Stdin)
	cfg.Client = &http.Client{}
	cfg.Pokedex = make(map[string]Pokemon)
	log.SetOutput(logFile)
	cfg.Logger = logFile

	return &cfg
}
