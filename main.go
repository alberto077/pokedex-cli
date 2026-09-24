package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	command     func() error
}

var allCliCommands map[string]cliCommand

func cleanInput(text string) []string {
	return strings.Split(strings.TrimSpace(strings.ToLower(text)), " ")

}

func commandExit() error {
	fmt.Println("Exiting the Pokedex CLI ....")
	fmt.Println("Goodbye")
	os.Exit(1)
	return nil
}

func commandHelp() error {
	fmt.Println("Pokedex CLI")
	fmt.Print("Usage: \n\n\n\n")
	for name, info := range allCliCommands {
		fmt.Printf("%s: %s", name, info.description)
	}

	return nil

}

func init() {

	allCliCommands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit out of the Pokedex CLI",
			command:     commandExit,
		},
		"help": {
			name:        "help",
			description: "Helpful descriptions of all commands",
			command:     commandHelp,
		},
	}

}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the pokedex CLI")

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		words := cleanInput(scanner.Text())

		fmt.Printf("Your command was: %s\n", words[0])

	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %v\n", err)
	}

}
