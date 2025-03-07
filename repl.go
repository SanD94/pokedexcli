package main

import (
	"strings"
)

func cleanInput(text string) []string {
	smallText := strings.ToLower(text)

	return strings.Fields(smallText)
}
