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

var (
	commandCallers []string // ordered callers
	allCliCommands map[string]cliCommand
	exitSignal chan struct{}
)

func cleanInput(text string) []string {
	return strings.Split(strings.TrimSpace(strings.ToLower(text)), " ")

}

// Send a stop signal which makes the for loop return which will make it so the loop stops and the program exits
func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	close(exitSignal)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _ , caller := range commandCallers {


		callerInfo, ok := allCliCommands[caller]

		if !ok {
			return fmt.Errorf("Could not find command %s in command list", caller)
		}
		
		fmt.Printf("%s: %s\n", callerInfo.name, callerInfo.description)
	}

	return nil

}

func init() {

	exitSignal = make(chan struct{})

	commandCallers = []string{"exit", "help"}

	allCliCommands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			command:     commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			command:     commandHelp,
		},
	}

}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the pokedex CLI")

	for {
		select {
			case <-exitSignal:
				return 
			default:
				
				fmt.Print("Pokedex > ")
				scanner.Scan()
				if err := scanner.Err(); err != nil {
					fmt.Printf("Scanner eskror: %v\n", err)
				}
				input := cleanInput(scanner.Text())[0]

				command, ok := allCliCommands[input]

				if !ok {
					fmt.Println("Invalid command, use the help command to discover what you can do")
				} else {
					command.command()
				}
			}
	


	}


}
