package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const URL = "https://pokeapi.co/api/v2/"

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	callRegister()
	config := Config{next: "", prev: ""}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		lines := scanner.Text()
		word := cleanInput(lines)[0]
		command, ok := registry[word]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		command.callback(&config)
	}

}

func cleanInput(text string) []string {
	smallText := strings.ToLower(text)

	return strings.Fields(smallText)
}
