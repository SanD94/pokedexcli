package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	callRegister()

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
		command.callback()
	}
}
