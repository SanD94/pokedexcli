package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type locArea struct {
	Name string `json:"name"`
	URL  string `json:"URL"`
}

type locAreas struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []locArea `json:"results"`
}

func commandMap(config *Config) error {
	var curURL string
	if config.next == "" {
		curURL = URL + "location-area/"
	} else {
		curURL = config.next
	}
	res, err := http.Get(curURL)

	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var areas locAreas
	if err := json.Unmarshal(body, &areas); err != nil {
		return err
	}

	config.next = areas.Next
	config.prev = areas.Previous

	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}

	return nil
}
