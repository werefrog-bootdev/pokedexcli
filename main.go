package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ----- CLI plumbing -----

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, []string) error
}

func cleanInput(text string) []string {
	text = strings.TrimSpace(strings.ToLower(text))
	return strings.Fields(text)
}

// Construit la map de commandes et câble help via une closure
func getCommands(cfg *Config) map[string]cliCommand {
	cmds := make(map[string]cliCommand)

	cmds["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback: func(_ *Config, _ []string) error {
			fmt.Println("Closing the Pokedex... Goodbye!")
			os.Exit(0)
			return nil
		},
	}

	cmds["map"] = cliCommand{
		name:        "map",
		description: "Show next 20 location-areas",
		callback:    commandMap, // from commands_map.go
	}

	cmds["mapb"] = cliCommand{
		name:        "mapb",
		description: "Show previous 20 location-areas",
		callback:    commandMapBack, // from commands_map.go
	}

	// help utilise makeHelpCommand (défini dans commands_map.go)
	helpCb := makeHelpCommand(&cmds)
	cmds["help"] = cliCommand{
		name:        "help",
		description: "Show available commands",
		callback:    helpCb,
	}

	return cmds
}

func main() {
	// Pagination state lives here
	cfg := &Config{} // NextURL/PrevURL will be populated after first map call
	commands := getCommands(cfg)

	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !sc.Scan() {
			fmt.Println()
			return
		}

		words := cleanInput(sc.Text())
		if len(words) == 0 {
			continue
		}

		cmdName := words[0]
		args := words[1:]

		cmd, ok := commands[cmdName]
		if !ok {
			fmt.Println("Unknown command. Type 'help'.")
			continue
		}

		if err := cmd.callback(cfg, args); err != nil {
			fmt.Println("Error:", err)
		}
	}
}
