package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	return strings.Split(strings.TrimSpace(strings.ToLower(text)), " ")

}

func startRepl(sharedState *config){

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

				command, ok := sharedState.allCliCommads[input]

				if !ok {
					fmt.Println("Invalid command, use the help command to discover what you can do")
				} else {
					command.command(sharedState)
				}
			}
	


	}

}
