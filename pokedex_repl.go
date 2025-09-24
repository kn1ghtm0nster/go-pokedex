package main

import (
	"bufio"
	"fmt"
	"os"
)

func PokedexREPL() {
	inputScanner := bufio.NewScanner(os.Stdin)

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
			cmd.callback()
		} else {
			fmt.Println("Unknown command")
		}
		
	}
}