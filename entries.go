package main

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type Config struct {
	next string
	prev string
}

var registry map[string]cliCommand

// registry call for each command
func callRegister() {
	registry = make(map[string]cliCommand)

	addRegister("help", cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	})

	addRegister("map", cliCommand{
		name:        "map",
		description: "List the areas in Pokemon World (next)",
		callback:    commandMap,
	})

	addRegister("mapb", cliCommand{
		name:        "mapb",
		description: "List the areas in Pokemon World (prev)",
		callback:    commandPMap,
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
