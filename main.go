package main


func main () {
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
	}

	PokedexREPL()
}
