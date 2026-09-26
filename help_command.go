package main

import (
	"fmt"
)

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
