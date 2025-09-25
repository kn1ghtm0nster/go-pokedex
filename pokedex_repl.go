package main

import (
	"bufio"
	"fmt"
	"os"
)

func PokedexREPL() {
	inputScanner := bufio.NewScanner(os.Stdin)

	config := Config{
		Next:     "",
		Previous: "",
	}

	supportedCommands := map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Displays the names of 20 location areas in the Pokemon world. Can be used again to show the next 20 location area names.",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world. Can be used again to show the previous 20 location area names.",
			callback: commandPreviousMap,
		},
	}

	for {
		fmt.Print("Pokedex > ")
		if !inputScanner.Scan() {
			if err := inputScanner.Err(); err != nil {
				fmt.Println("Error reading input:", err)
			}
			break
		}

		userInput := inputScanner.Text()
		cleanedInput := cleanInput(userInput)

		if len(cleanedInput) == 0 {
			continue
		}

		command := cleanedInput[0]

		if cmd, exists := supportedCommands[command]; exists {
			cmd.callback(&config)
		} else {
			fmt.Println("Unknown command")
		}
		
	}
}