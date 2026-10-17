package main
/*
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
