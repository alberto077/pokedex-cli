package main

import (
	"fmt"
)

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	close(exitSignal)
	return nil
}
