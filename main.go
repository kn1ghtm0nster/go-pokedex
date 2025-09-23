package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main () {
	inputScanner := bufio.NewScanner(os.Stdin)

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
		
		fmt.Printf("Your command was: %s\n", command)
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