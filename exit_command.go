package main

import (
	"fmt"
)

func commandExit(sharedState *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	close(exitSignal)
	return nil
}
