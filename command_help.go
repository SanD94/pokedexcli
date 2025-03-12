package main

import (
	"fmt"
)

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")

	for _, v := range registry {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}

	return nil
}
