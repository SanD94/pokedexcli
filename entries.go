package main

import (
	"github.com/sand94/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	register         *commandRegister
}

type commandRegister struct {
	registry map[string]cliCommand
}

// registry call for each command
func (cr *commandRegister) callRegister() {
	cr.registry = make(map[string]cliCommand)

	cr.addRegister("help", cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	})

	cr.addRegister("map", cliCommand{
		name:        "map",
		description: "Get the next page of locations",
		callback:    commandMapf,
	})

	cr.addRegister("mapb", cliCommand{
		name:        "mapb",
		description: "Get the previous page of locations",
		callback:    commandMapb,
	})

	cr.addRegister("exit", cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	})
}

func (cr *commandRegister) addRegister(name string, command cliCommand) {
	cr.registry[name] = command
}

func (cr commandRegister) getCommands() map[string]cliCommand {
	return cr.registry
}
