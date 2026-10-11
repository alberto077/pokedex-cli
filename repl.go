package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/alberto077/pokedexcli/internal/pokecache"
)

func cleanInput(text string) []string {
	return strings.Split(strings.TrimSpace(strings.ToLower(text)), " ")

}

func startRepl(appConfig *appConfig, cache *pokecache.Cache) {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the pokedex CLI")

	for {
		select {
		case <-appConfig.ExitSignal:
			return
		default:

			fmt.Print("Pokedex > ")
			scanner.Scan()
			if err := scanner.Err(); err != nil {
				fmt.Printf("Scanner eskror: %v\n", err)
				return
			}
			appConfig.CliInput = cleanInput(scanner.Text())

			command, ok := appConfig.AllCliCommands[appConfig.CliInput[0]]

			if !ok {
				fmt.Println("Invalid command, use the help command to discover what you can do")
			} else {
				command.Command(appConfig, cache)
			}
		}

	}

}
