package cli

import "strings"

func CleanInput(text string) []string {
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
