package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, url *string) error {
	locationResp, err := cfg.pokeapiClient.ListLocations(url)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, area := range locationResp.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapf(cfg *config) error {
	err := commandMap(cfg, cfg.nextLocationsURL)
	return err
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}
	err := commandMap(cfg, cfg.prevLocationsURL)
	return err
}
