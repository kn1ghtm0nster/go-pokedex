package cli

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/kn1ghtm0nster/go-pokedex/internal/pokeapi"
)

var supportedCommands map[string]cliCommand

func commandExit(conf *Config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *Config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")

	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range supportedCommands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

func commandMap(conf *Config, args []string) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if conf.Next != "" {
		url = conf.Next
	} else if conf.Next == "" && conf.Previous != "" {
		fmt.Println("No more pages, returning to the first page.")
		conf.Next = ""
		conf.Previous = ""
		return nil
	}

	info, exists := conf.Cache.Get(url)
	if exists {
		serializedInfo := pokeapi.PokeAPILocationAreaResponse{}
		if err := json.Unmarshal(info, &serializedInfo); err != nil {
			fmt.Println("Error unmarshalling cached data:", err)
		}

		fmt.Println("[INFO] - Cached data found. Using existing data.")
		for _, location := range serializedInfo.Results {
			fmt.Println(location.Name)
		}
		return nil
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data := pokeapi.PokeAPILocationAreaResponse{}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&data); err != nil {
		return err
	}

	for _, location := range data.Results {
		fmt.Println(location.Name)
	}

	byteData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling data for cache:", err)
		return nil
	}
	conf.Cache.Add(url, byteData)

	conf.Next = data.Next
	conf.Previous = data.Previous

	return nil
}

func commandPreviousMap(conf *Config, args []string) error {
	url := "https://pokeapi.co/api/v2/location-area/"

	if conf.Previous != "" {
		url = conf.Previous
	} else if conf.Previous == "" && conf.Next == "" {
		fmt.Println("No more pages, returning to the first page.")
		conf.Next = ""
		conf.Previous = ""
		return nil
	}

	info, exists := conf.Cache.Get(url)
	if exists {
		serializedInfo := pokeapi.PokeAPILocationAreaResponse{}
		if err := json.Unmarshal(info, &serializedInfo); err != nil {
			fmt.Println("Error unmarshalling cached data:", err)
		}

		fmt.Println("[INFO] - Cached data found. Using existing data.")
		for _, location := range serializedInfo.Results {
			fmt.Println(location.Name)
		}
		return nil
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	data := pokeapi.PokeAPILocationAreaResponse{}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&data); err != nil {
		return err
	}

	for _, location := range data.Results {
		fmt.Println(location.Name)
	}

	byteData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling data for cache:", err)
		return nil
	}
	conf.Cache.Add(url, byteData)

	conf.Next = data.Next
	conf.Previous = data.Previous

	return nil
}

func commandExplore(conf *Config, args []string) error {
	if len(args) == 0 {
		fmt.Println("Please provide a location area name. Usage: explore <location-area-name>")
		return nil
	}
	locationName := args[0]
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", locationName)

	info, exists := conf.Cache.Get(url)
	if exists {
		serializedInfo := pokeapi.PokeAPILocationAreaDetail{}
		if err := json.Unmarshal(info, &serializedInfo); err != nil {
			fmt.Println("Error unmarshalling cached data:", err)
			return nil
		}

		fmt.Println("[INFO] - Cached data found. Using existing data.")
		fmt.Println()
		fmt.Println("Exploring", locationName, "...")
		fmt.Println("Found Pokemon:")
		for _, encounter := range serializedInfo.PokemonEncounters {
			fmt.Println(" -", encounter.Pokemon.Name)
		}
		return nil
	}

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data from PokeAPI:", err)
		return nil
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Printf("Error: Received status code %d from PokeAPI. Please check the location area name.\n", res.StatusCode)
		return nil
	}

	data := pokeapi.PokeAPILocationAreaDetail{}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&data); err != nil {
		fmt.Println("Error decoding response from PokeAPI:", err)
		return nil
	}
	
	fmt.Println("Exploring", locationName, "...")
	fmt.Println("Found Pokemon:")
	for _, encounter := range data.PokemonEncounters {
		fmt.Println(" -", encounter.Pokemon.Name)
	}

	byteData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling data for cache:", err)
		return nil
	}
	conf.Cache.Add(url, byteData)

	return nil
}

func commandCatch(conf *Config, args []string) error {
	if len(args) == 0 {
		fmt.Println("Please provide a pokemon name. Usage: catch <pokemon-name>")
		return nil
	}
	pokemonName := strings.ToLower(args[0])

	_, alreadyCaught := conf.Caught[pokemonName]
	if alreadyCaught {
		fmt.Printf("You have already caught %s!\n", pokemonName)
		return nil
	}

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonName)

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data from PokeAPI:", err)
		return nil
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Printf("Error: Received status code %d from PokeAPI. Please check the pokemon name.\n", res.StatusCode)
		return nil
	}

	pokemonDetail := pokeapi.PokemonDetail{}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&pokemonDetail); err != nil {
		fmt.Println("Error decoding response from PokeAPI:", err)
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	time.Sleep(2 * time.Second)

	playerLuck := rand.Intn(351)

	if playerLuck > pokemonDetail.BaseExperience {
		fmt.Printf("%s was caught!\n", pokemonName)
		conf.Caught[pokemonName] = pokemonDetail
		return nil
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}
}

func commandInspect(conf *Config, args []string) error {
	if len(args) == 0 {
		fmt.Println("Please provide a pokemon name. Usage: inspect <pokemon-name>")
		return nil
	}

	pokemonName := strings.ToLower(args[0])
	pokemonDetail, exists := conf.Caught[pokemonName]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	var sb strings.Builder
	fmt.Fprintf(&sb, "Name: %s\n", pokemonDetail.Name)
	fmt.Fprintf(&sb, "Height: %d\n", pokemonDetail.Height)
	fmt.Fprintf(&sb, "Weight: %d\n", pokemonDetail.Weight)

	fmt.Fprintf(&sb, "Stats:\n")
	for _, stat := range pokemonDetail.Stats {
		fmt.Fprintf(&sb, "  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Fprintf(&sb, "Types:\n")
	for _, t := range pokemonDetail.Types {
		fmt.Fprintf(&sb, "  -%s\n", t.Type.Name)
	}

	fmt.Print(sb.String())
	return nil
}

func init() {
	supportedCommands = map[string]cliCommand{
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
		"explore": {
			name: "explore",
			description: "Displays pokemon in a specific location area.",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Simulates catching a specific pokemon.",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "Displays detailed information about a specific caught pokemon.",
			callback: commandInspect,
		},
	}
}