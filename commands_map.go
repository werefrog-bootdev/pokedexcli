package main

import (
	"errors"
	"fmt"
)

// map: show next 20 location-areas
func commandMap(cfg *Config, _ []string) error {
	resp, err := fetchLocationAreas(cfg, cfg.NextURL) // empty NextURL means "first page"
	if err != nil {
		return err
	}
	// print names
	for _, r := range resp.Results {
		fmt.Println(r.Name)
	}
	// update pagination
	cfg.NextURL = resp.Next
	cfg.PrevURL = resp.Previous
	return nil
}

// mapb: show previous 20 location-areas
func commandMapBack(cfg *Config, _ []string) error {
	if cfg.PrevURL == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	resp, err := fetchLocationAreas(cfg, cfg.PrevURL)
	if err != nil {
		return err
	}
	for _, r := range resp.Results {
		fmt.Println(r.Name)
	}
	cfg.NextURL = resp.Next
	cfg.PrevURL = resp.Previous
	return nil
}

// help command prints available commands (wired with closure over commands map)
func makeHelpCommand(commands *map[string]cliCommand) func(*Config, []string) error {
	return func(_ *Config, _ []string) error {
		if commands == nil {
			return errors.New("commands not initialised")
		}
		fmt.Println("Welcome to the Pokedex!")
		fmt.Println("Available commands:")
		// stable order: help first, then alphabetical
		names := make([]string, 0, len(*commands))
		for name := range *commands {
			names = append(names, name)
		}
		// put "help" first
		// simple selection sort for tiny set
		for i := 0; i < len(names); i++ {
			for j := i + 1; j < len(names); j++ {
				a, b := names[i], names[j]
				if (b == "help" && a != "help") || (a != "help" && b != "help" && b < a) {
					names[i], names[j] = names[j], names[i]
				}
			}
		}
		for _, name := range names {
			fmt.Printf(" - %-6s : %s\n", name, (*commands)[name].description)
		}
		return nil
	}
}
