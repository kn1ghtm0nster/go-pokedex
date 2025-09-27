package cli

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/kn1ghtm0nster/go-pokedex/internal/pokecache"
)


func PokedexREPL() {
	inputScanner := bufio.NewScanner(os.Stdin)

	config := Config{
		Next:     "",
		Previous: "",
		Cache: pokecache.NewCache(5 * time.Second),
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
		cleanedInput := CleanInput(userInput)

		if len(cleanedInput) == 0 {
			continue
		}

		command := cleanedInput[0]
		args := cleanedInput[1:]

		if cmd, exists := supportedCommands[command]; exists {
			cmd.callback(&config, args)
		} else {
			fmt.Println("Unknown command")
		}
		
	}
}