package service

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/alberto077/pokedexcli/internal/api"
	"github.com/alberto077/pokedexcli/internal/config"
	"github.com/alberto077/pokedexcli/internal/pokecache"
)

var cfg *config.AppConfig

type SequentialMapTransport struct {
	count     int
	responses []*http.Response
}

func (s *SequentialMapTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	resp := s.responses[s.count]
	s.count++
	return resp, nil
}

func TestMain(m *testing.M) {

	cfg = config.InitConfig()

	cfg.CommandRegistry = config.NewCommandRegistry()

	cfg.CommandRegistry.RegisterCommand(cfg, config.CliCommand{
		Name:        "exit",
		Description: "Exit the Pokedex",
		Command:     &Exit{},
	})

	cfg.CommandRegistry.RegisterCommand(cfg, config.CliCommand{
		Name:        "help",
		Description: "Displays a help message",
		Command:     &Help{},
	})

	cfg.CommandRegistry.RegisterCommand(cfg, config.CliCommand{

		Name:        "map",
		Description: "Displays a list of locations and moves foward on every request",
		Command:     &Map{},
	})

	cfg.CommandRegistry.RegisterCommand(cfg, config.CliCommand{

		Name:        "mapb",
		Description: "Displays a list of locations and moves backwards on every request",
		Command:     &MapB{},
	})

	cfg.CommandRegistry.RegisterCommand(cfg, config.CliCommand{

		Name:        "explore",
		Description: "Expore the location specified after the command which will show pokemon encounterd",
		Command:     &ExploreLocation{},
	})

	cfg.CommandRegistry.RegisterCommand(cfg, config.CliCommand{
		Name:        "catch",
		Description: "Throw a Pokeball for a chance to catch a pokemon",
		Command:     &Catch{},
	})

	cfg.CommandRegistry.RegisterCommand(cfg, config.CliCommand{

		Name:        "inspect",
		Description: "Get stats about a pokemon youve caught and stored in your pokedex",
		Command:     &Inspect{},
	})

	cfg.CommandRegistry.RegisterCommand(cfg, config.CliCommand{

		Name:        "pokedex",
		Description: "Get a list of all the pokemon you have caught",
		Command:     &OpenPokedex{},
	})

	cfg.CommandRegistry.RegisterCommand(cfg, config.CliCommand{

		Name:        "release",
		Description: "Release one of the pokemon you have caught",
		Command:     &Release{},
	})
	exitCode := m.Run()

	os.Exit(exitCode)
}

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
			got := CleanInput(tc.input)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: %#v, got: %#v", tc.want, got)
			}
		})

	}

}

func TestHelp(t *testing.T) {

	cache := pokecache.NewCache(5 * time.Second)
	test := map[string]struct {
		input string
		want  string
	}{

		"help": {input: "help", want: "Welcome to the Pokedex!\nUsage:\n\nexit: Exit the Pokedex\nhelp: Displays a help message\nmap: Displays a list of locations and moves foward on every request\nmapb: Displays a list of locations and moves backwards on every request\nexplore: Expore the location specified after the command which will show pokemon encounterd\ncatch: Throw a Pokeball for a chance to catch a pokemon\ninspect: Get stats about a pokemon youve caught and stored in your pokedex\npokedex: Get a list of all the pokemon you have caught\nrelease: Release one of the pokemon you have caught\n"},
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

func TestMap(t *testing.T) {

	cache := pokecache.NewCache(5 * time.Second)

	locationResponse := api.LocationResponse{
		Next:     "next.com",
		Previous: "",
		Results:  []api.Location{{Name: "canalave-city-area", Url: "curr.com"}},
	}
	locationResponseSubsequent := api.LocationResponse{
		Next:     "nextnext.com",
		Previous: locationResponse.Results[0].Url,
		Results:  []api.Location{{Name: "next-area", Url: locationResponse.Next}},
	}

	jsonResponse, err := json.Marshal(locationResponse)

	if err != nil {
		t.Fatalf("Could not marshal json")
	}

	jsonResponseSubsequent, err := json.Marshal(locationResponseSubsequent)

	if err != nil {
		t.Fatalf("Could not marshal json")
	}

	transport := &SequentialMapTransport{
		count: 0,
		responses: []*http.Response{
			{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(jsonResponse))},
			{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(jsonResponseSubsequent))},
			{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(jsonResponse))},
			//Could add a cache test instead of creating a base response twice
		},
	}

	cfg.Client = &http.Client{
		Transport: transport,
	}
	tests := map[string]struct {
		want  string
		input string
	}{
		"Base Map": {
			want: "First Page\ncanalave-city-area\n", input: "map",
		},
		"Map Forward": {
			want: "next-area\n", input: "map",
		},
		"Map back": {
			want: "First Page\ncanalave-city-area\n", input: "mapb",
		},
	}

	test_order := []string{"Base Map", "Map Forward", "Map back"}

	for _, name := range test_order {
		tc := tests[name]
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
				t.Fatalf("expected: %s, got %s", tc.want, got)
			}

		})
	}
}

func TestExplore(t *testing.T) {
	cache := pokecache.NewCache(5 * time.Second)

	pokemonEncounter := config.AreaEncounterResponse{
		PokemonEncounters: []config.PokemonEncounters{
			{
				Pokemon: config.PokemonBaseInfo{
					Name: "pikachu",
					URL:  "pikapika.com",
				},
			},
			{

				Pokemon: config.PokemonBaseInfo{
					Name: "charizard",
					URL:  "charichar.com",
				},
			},
		}}
	jsonResponse, err := json.Marshal(pokemonEncounter)

	if err != nil {
		t.Fatalf("Could not marshal json")
	}

	exploreTransport := &SequentialMapTransport{
		count: 0,
		responses: []*http.Response{
			{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(jsonResponse))},
		},
	}

	cfg.Client = &http.Client{
		Transport: exploreTransport,
	}

	tests := map[string]struct {
		want  string
		input string
	}{
		"Explore area": {want: "Exploring canalave-city-area ...\nPokemon in canalave-city-area:\n - pikachu\n - charizard\n",
			input: "explore"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			var builder strings.Builder
			originalStdOut := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			cfg.CliInput = []string{"explore", "canalave-city-area"}
			cfg.CommandRegistry.AllCliCommands[tc.input].Execute(cfg, cache)

			w.Close()

			io.Copy(&builder, r)
			os.Stdout = originalStdOut

			got := builder.String()

			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected %s, got: %s", tc.want, got)
			}

		})
	}
}

func TestCatch(t *testing.T) {
	cache := pokecache.NewCache(5 * time.Second)

	pokemon := config.Pokemon{
		ID:             1,
		Name:           "pikachu",
		BaseExperience: 1,
	}
	jsonResponse, err := json.Marshal(pokemon)
	if err != nil {
		t.Fatalf("Could not marshal json")
	}

	catchTransport := &SequentialMapTransport{
		count: 0,
		responses: []*http.Response{
			{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(jsonResponse))},
		},
	}

	cfg.Client = &http.Client{
		Transport: catchTransport,
	}

	tests := map[string]struct {
		input []string
		want  string
	}{
		"catchPikachu": {want: "Throwing a Pokeball at pikachu...\npikachu was caught!\n", input: []string{"catch", "pikachu"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var builder strings.Builder
			originalStdOut := os.Stdout
			r, w, err := os.Pipe()
			if err != nil {
				t.Fatalf("Could not create pipe to capture stdout")
			}
			os.Stdout = w
			cfg.CliInput = tc.input
			cfg.CommandRegistry.AllCliCommands[tc.input[0]].Execute(cfg, cache)

			w.Close()

			io.Copy(&builder, r)
			os.Stdout = originalStdOut

			got := builder.String()

			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}

			if pokemon, ok := cfg.Pokedex[tc.input[1]]; !ok || pokemon.Name != tc.input[1] {
				t.Fatalf("%s was not stored in the pokedex", tc.input[1])

			}

		})
	}
}
