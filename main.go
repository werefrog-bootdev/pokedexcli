package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	// Build a stable order: "help" first, then alphabetical
	names := make([]string, 0, len(commands))
	for name := range commands {
		names = append(names, name)
	}
	sort.Slice(names, func(i, j int) bool {
		if names[i] == "help" && names[j] != "help" {
			return true
		}
		if names[j] == "help" && names[i] != "help" {
			return false
		}
		return names[i] < names[j]
	})

	for _, name := range names {
		cmd := commands[name]
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// Declare first (no init expression referencing commandHelp)
var commands map[string]cliCommand

// Populate later to avoid the init cycle
func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			fmt.Println()
			return
		}

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		if cmd, ok := commands[words[0]]; ok {
			if err := cmd.callback(); err != nil {
				fmt.Println("Error:", err)
			}
			continue
		}

		fmt.Println("Unknown command")
	}
}
