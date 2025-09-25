package main

import (
	"bufio"
	"fmt"
	"os"
)

// NOTE: to prevent initialization issues, we declare supportedCommands here
// and initialize it in main.go
var supportedCommands map[string]cliCommand

func PokedexREPL() {
	inputScanner := bufio.NewScanner(os.Stdin)

	config := Config{
		Next:     "",
		Previous: "",
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