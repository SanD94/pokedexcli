package main

import (
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	smallText := strings.ToLower(text)

	return strings.Fields(smallText)
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")

	for _, v := range registry {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}

	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var registry map[string]cliCommand

func callRegister() {
	registry = make(map[string]cliCommand)
	addRegister("help", cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	})

	addRegister("exit", cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	})
}

func addRegister(name string, command cliCommand) {
	registry[name] = command
}
