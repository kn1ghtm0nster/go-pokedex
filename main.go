package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main () {
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

func cleanInput(text string) []string {
	// split the input text into a slice of words
	words := strings.Fields(text)

	// remove any punctuation from the words
	for i, word := range words {
		words[i] = strings.Trim(word, ".,!?\"';:-()[]{}")
	}

	// lowercase the words, and trim any whitespace
	for i, word := range words {
		words[i] = strings.ToLower(strings.TrimSpace(word))
	}
	// return the cleaned slice of words
	return words
}

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