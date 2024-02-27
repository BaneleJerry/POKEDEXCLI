package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)
	for {

		fmt.Printf("pokedex >")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanIput(text)
		availableCommands := getCommands()

		cmd, ok := availableCommands[cleaned[0]]

		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		cmd.callback()

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map":{
			name:        "map",
            description: "Displays the map in pokemon World",
            callback:    commandMap,
		},
	}
}

func cleanIput(str string) []string {
	lowerd := strings.ToLower(str)
	words := strings.Fields(lowerd)
	return words
}
