package config

import (
	"bufio"
	"github.com/alberto077/pokedexcli/internal/pokecache"
	"log"
	"net/http"
	"os"
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
	ConditionValues []any  `json:"condition_values"`
	Method          Method `json:"method"`
	MinLevel        int    `json:"min_level"`
	MaxLevel        int    `json:"max_level"`
	Chance          int    `json:"chance"`
}
type PokemonEncounters struct {
	Pokemon PokemonBaseInfo `json:"pokemon"`
}

type Pokemon struct {
	Abilities      []Abilities `json:"abilities"`
	Forms          []Forms     `json:"forms"`
	Moves          []Moves     `json:"moves"`
	Stats          []Stats     `json:"stats"`
	Types          []Types     `json:"types"`
	ID             int         `json:"id"`
	Name           string      `json:"name"`
	Height         int         `json:"height"`
	Weight         int         `json:"weight"`
	BaseExperience int         `json:"base_experience"`
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
	VersionGroup    VersionGroup    `json:"version_group"`
	MoveLearnMethod MoveLearnMethod `json:"move_learn_method"`
	LevelLearnedAt  int             `json:"level_learned_at"`
	Order           int             `json:"order"`
}

type Moves struct {
	VersionGroupDetails []VersionGroupDetails `json:"version_group_details"`
	Move                Move                  `json:"move"`
}
type Stat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Stats struct {
	Stat     Stat `json:"stat"`
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
}

type Type struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Types struct {
	Type Type `json:"type"`
	Slot int  `json:"slot"`
}

type Command interface {
	Execute(appconifg *AppConfig, cache *pokecache.Cache) error
}

type CommandRegistry struct {
	AllCliCommands map[string]Command
}

type CliCommand struct {
	Command     Command
	Name        string
	Description string
}

type LocationData struct {
	NextPage     string
	PreviousPage string
	Offset       int16
	Page_limit   int16
}

type AppConfig struct {
	CommandCallers       []CliCommand
	CliInput             []string
	LocationRequestState LocationData
	ExitSignal           chan struct{}
	CommandRegistry      *CommandRegistry
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
