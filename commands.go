package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")

	fmt.Println("Usage:")
	fmt.Println()

	commands := []string{"exit", "help"}
	descriptions := []string{
		"Exit the Pokedex",
		"Displays a help message",
	}

	for i, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd, descriptions[i])
	}

	return nil
}