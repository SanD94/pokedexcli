package main

import (
	"fmt"
	"bufio"
	"os"
)

func main()  {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		lines := scanner.Text()
		word := cleanInput(lines)[0]
		fmt.Printf("Your command was: %s\n", word)
	}
}
