package main

import "fmt"

func commandHelp() error {

	availableCommands := getCommands()

	fmt.Println("Welcome Pokedex \n Usage:")
	for _, cmd := range availableCommands {

		fmt.Printf("-%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Printf("")
	return nil
}
