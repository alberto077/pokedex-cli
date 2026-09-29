package main

import (
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
	"github.com/alberto077/pokedexcli/internal/api"
	"github.com/alberto077/pokedexcli/configuration"
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

func TestHelp(t *testing.T){
	test := map[string]struct{
		input string
		want string

	}{

		"help": {input: "help", want: "Welcome to the Pokedex!\nUsage:\n\nexit: Exit the Pokedex\nhelp: Displays a help message\n"},
	}

	for name, tc := range test{

	t.Run(name, func(t *testing.T){
			var got strings.Builder

			originalStdOut := os.Stdout

			reader, writer, _ := os.Pipe()
			os.Stdout = writer


			configuration.SharedState.AllCliCommands[tc.input].Command(&configuration.SharedState)

			writer.Close()
			os.Stdout = originalStdOut
			
			io.Copy(&got, reader)


			if !reflect.DeepEqual(tc.want, got.String()) {
				t.Fatalf("expected: %q, got: %q", tc.want, got.String())
			}
		})

	}
}

	
// since we are closing the channel that means that only one function should be executing it when we execite help 
func TestExit(t *testing.T) {
	tests := map[string]struct{
		input string
		want string
	}{
		"exit": {input: "exit", want: "Closing the Pokedex... Goodbye!\n"},
	}
	
	for name, tc := range tests{
		t.Run(name, func(t *testing.T) {
			
			var got strings.Builder
			originalStdOut := os.Stdout

			reader, writer, _ := os.Pipe()
			os.Stdout = writer


			configuration.SharedState.AllCliCommands[tc.input].Command(&configuration.SharedState)

			writer.Close()
			os.Stdout = originalStdOut
			
			io.Copy(&got, reader)

			if !reflect.DeepEqual(tc.want, got.String()) {
				t.Fatalf("expected: %q, got: %q", tc.want, got.String())
			}
		})
	}
}


func TestBaseLocationsApiRequest(t *testing.T) {
	configuration.SharedState.LocationRequestState.Offset = 0
	configuration.SharedState.LocationRequestState.Page_limit = 1

	tests := map[string]struct{
		want api.Locations
	}{
		"base": {
			want: api.Locations{

				Count: 1539, Next: "https://pokeapi.co/api/v2/location-area?offset=1&limit=1",
				Previous:"",
				Results: []api.NamedAPIResource{
					{
						Name:"canalave-city-area",
						Url:"https://pokeapi.co/api/v2/location-area/1/",
					},
				},
			},


	},

	}

	for name, tc := range tests{
		t.Run(name, func(t *testing.T){

			got, err := api.Locations_Request(&configuration.SharedState)


			if err != nil {
				t.Fatalf("API ERROR: %v", err)
			}

			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: %v, got %v", tc.want, got, )
			}

		})
	}
}




func TestSubsequentLocationsApiRequest(t *testing.T) {

	configuration.SharedState.LocationRequestState.Offset = 0
	configuration.SharedState.LocationRequestState.Page_limit = 1

	tests := map[string]struct{
		want api.Locations
	}{
		"base": {
			want: api.Locations{

				Count: 1539, Next: "https://pokeapi.co/api/v2/location-area?offset=1&limit=1",
				Previous:"",
				Results: []api.NamedAPIResource{
					{
						Name:"canalave-city-area",
						Url:"https://pokeapi.co/api/v2/location-area/1/",
					},
				},
			},



		},
		"request2": {
			want: api.Locations{

				Count: 1539, Next: "https://pokeapi.co/api/v2/location-area?offset=2&limit=1",
				Previous:"https://pokeapi.co/api/v2/location-area?offset=0&limit=1",
				Results: []api.NamedAPIResource{
					{
						Name:"eterna-city-area",
						Url:"https://pokeapi.co/api/v2/location-area/2/",
					},
				},
			},



		},

		}


	for name, tc := range tests{
		t.Run(name, func(t *testing.T){

			got, err := api.Locations_Request(&configuration.SharedState)


			configuration.SharedState.LocationRequestState.Offset += configuration.SharedState.LocationRequestState.Page_limit

			if err != nil {
				t.Fatalf("API ERROR: %v", err)
			}

			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: %v, got %v, sharedState status: %d", tc.want, got, configuration.SharedState.LocationRequestState.Offset)
			}

		})
	}
}


























