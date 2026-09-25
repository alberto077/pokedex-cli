package main

import (
	"io"
	"os"
	"reflect"
	"testing"
	"strings"
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

// Could be helpful to just split the commands up into a exit and help test and then allow the exit to run in a subprocess and execute its own 
// Exit out of the subprocess where as the help would need something to force it out the subprocess
// Easiest test would be to just see if the exit one exits with a 0 but we just care about whats being displayed right or do we just want the commands to execute not sure

// It would be useful to test for a case where something might hang as well

// Probably wouldve been easier to just make a global variable and then have it be set to true when the exit function is called
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


			allCliCommands[tc.input].command()

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


			allCliCommands[tc.input].command()

			writer.Close()
			os.Stdout = originalStdOut
			
			io.Copy(&got, reader)

			if !reflect.DeepEqual(tc.want, got.String()) {
				t.Fatalf("expected: %q, got: %q", tc.want, got.String())
			}
		})
	}
}
