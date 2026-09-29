package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/alberto077/pokedexcli/configuration"
)

func cleanInput(text string) []string {
	return strings.Split(strings.TrimSpace(strings.ToLower(text)), " ")

}

func startRepl(sharedState *configuration.Config){

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the pokedex CLI")

	for {
		select {
			case <-configuration.ExitSignal:
				return 
			default:
				
				fmt.Print("Pokedex > ")
				scanner.Scan()
				if err := scanner.Err(); err != nil {
					fmt.Printf("Scanner eskror: %v\n", err)
					return
				}
				input := cleanInput(scanner.Text())[0]

				command, ok := sharedState.AllCliCommands[input]

				if !ok {
					fmt.Println("Invalid command, use the help command to discover what you can do")
				} else {
					command.Command(sharedState)
				}
			}
	


	}

}
