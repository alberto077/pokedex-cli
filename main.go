package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func cleanInput(text string) []string {
	return strings.Split(strings.TrimSpace(strings.ToLower(text)), " ")

}

func main(){

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the pokedex CLI")




	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan(){
			break
		}


		words := cleanInput(scanner.Text())

		
		fmt.Printf("Your command was: %s\n", words[0])
		
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %v\n", err)
	}

	
}

