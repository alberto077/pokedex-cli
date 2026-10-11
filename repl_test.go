package main

import (
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/alberto077/pokedexcli/internal/config"
	"github.com/alberto077/pokedexcli/internal/pokecache"
	"github.com/alberto077/pokedexcli/internal/service"
)

func TestCleanInput(t *testing.T) {
	tests := map[string]struct {
		input string
		want  []string
	}{
		"whiteSpace": {input: " testing whiteSpace ", want: []string{"testing", "whitespace"}},
		"lowerCase":  {input: "LOWER CASE TEST", want: []string{"lower", "case", "test"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := cleanInput(tc.input)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: %#v, got: %#v", tc.want, got)
			}
		})

	}

}

func TestHelp(t *testing.T) {

	cfg := config.InitConfig()

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
		Command:     &service.Map{},
	})

	cfg.CommandRegistry.RegisterCommand(cfg, config.CliCommand{

		Name:        "explore",
		Description: "Expore the location specified after the command which will show pokemon encounterd",
		Command:     &service.ExploreLocation{},
	})

	cache := pokecache.NewCache(5 * time.Second)
	test := map[string]struct {
		input string
		want  string
	}{

		"help": {input: "help", want: "Welcome to the Pokedex!\nUsage:\n\nexit: Exit the Pokedex\nhelp: Displays a help message\nmap: Displays a list of locations and moves foward on every request\nmapb: Displays a list of locations and moves backwards on every request\nexplore: Expore the location specified after the command which will show pokemon encounterd\n"},
	}

	for name, tc := range test {

		t.Run(name, func(t *testing.T) {

			var builder strings.Builder
			originalStdOut := os.Stdout

			reader, writer, _ := os.Pipe()
			os.Stdout = writer

			cfg.CommandRegistry.AllCliCommands[tc.input].Execute(cfg, cache)

			writer.Close()
			os.Stdout = originalStdOut

			io.Copy(&builder, reader)

			got := builder.String()

			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: %s, got: %s", tc.want, got)
			}
		})

	}
}

// since we are closing the channel that means that only one function should be executing it when we execite help
func TestExit(t *testing.T) {
	cfg := config.InitConfig()

	cfg.CommandRegistry = config.NewCommandRegistry()

	cfg.CommandRegistry.RegisterCommand(cfg, config.CliCommand{
		Name:        "exit",
		Description: "Exit the Pokedex",
		Command:     &service.Exit{},
	})

	cache := pokecache.NewCache(5 * time.Second)
	tests := map[string]struct {
		input string
		want  string
	}{
		"exit": {input: "exit", want: "Closing the Pokedex... Goodbye!\n"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			var builder strings.Builder
			originalStdOut := os.Stdout

			reader, writer, _ := os.Pipe()
			os.Stdout = writer

			cfg.CommandRegistry.AllCliCommands[tc.input].Execute(cfg, cache)

			writer.Close()
			os.Stdout = originalStdOut

			io.Copy(&builder, reader)

			got := builder.String()

			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: %s, got: %s", tc.want, got)
			}
		})
	}
}

/*
func TestExploreAreaApi(t *testing.T) {
	cache := pokecache.NewCache(5 * time.Second)
	tests := map[string]struct {
		input string
		want  []string
	}{
		"explore canalave-city-area": {input: "canalave-city-area", want: []string{"tentacool", "tentacruel", "staryu", "magikarp", "gyarados", "wingull", "pelipper", "shellos", "gastrodon", "finneon", "lumineon"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			res, err := api.GetAreaEncounterResponse(cache, tc.input)
			if err != nil  {
				t.Fatalf("API ERROR, %v", err)
			}


			var got []string

			for _, pokemon := range res {
				got = append(got, pokemon.Name)
			}



			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: %s, got: %s", tc.want, got)
			}
		})
	}
}


func TestBaseLocationsApiRequest(t *testing.T) {
	cfg := GetCommands()
	cache := pokecache.NewCache(5 * time.Second)
	cfg.LocationRequestState.Offset = 0
	cfg.LocationRequestState.Page_limit = 1

	tests := map[string]struct {
		want string
	}{
		"base": {
			want: "canalave-city-area\n",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			var builder strings.Builder
			originalStdOut := os.Stdout

			reader, writer, _ := os.Pipe()
			os.Stdout = writer

			if err := commandMap(cfg, cache); err != nil {
				t.Fatalf("API ERROR: %v", err)
			}

			writer.Close()
			os.Stdout = originalStdOut

			io.Copy(&builder, reader)

			got := builder.String()

			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: %s, got %s", tc.want, got)
			}

		})
	}
}

func TestSubsequentLocationsApiRequest(t *testing.T) {

	cfg := GetCommands()
	cache := pokecache.NewCache(5 * time.Second)
	cfg.LocationRequestState.Offset = 0
	cfg.LocationRequestState.Page_limit = 1

	tests := map[string]struct {
		want string
	}{
		"base": {
			want: "canalave-city-area\n",
		},

		"request2": {
			want: "eterna-city-area\n",
		},
	}

	for _, name := range slices.Sorted(maps.Keys(tests)) {
		tc := tests[name]
		t.Run(name, func(t *testing.T) {

			var builder strings.Builder
			originalStdOut := os.Stdout

			reader, writer, _ := os.Pipe()
			os.Stdout = writer

			if err := commandMap(cfg, cache); err != nil {
				t.Fatalf("API ERROR: %v", err)
			}

			writer.Close()
			os.Stdout = originalStdOut

			io.Copy(&builder, reader)

			got := builder.String()

			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: %q, got: %q", tc.want, got)
			}

		})
	}
}
*/
