# Pokemon Pokedex!

A command-line Pokedex application written in Go.  
Explore, catch, and inspect Pokémon using the [PokeAPI](https://pokeapi.co/).

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.20 or newer recommended)

### Installation

1. **Clone the repository:**

   ```sh
   git clone https://github.com/kn1ghtm0nster/go-pokedex.git
   cd go-pokedex
   ```

2. **Install dependencies:**
   ```sh
   go mod tidy
   ```

### Running the Program

From the project root, run:

```sh
go run ./cmd/pokedex
```

You’ll see the REPL prompt:

```
Pokedex >
```

### Features

- `help` — Show available commands
- `exit` — Exit the Pokedex
- `map` / `mapb` — Explore Pokémon locations
- `explore <location>` — List Pokémon in a location
- `catch <pokemon>` — Attempt to catch a Pokémon
- `inspect <pokemon>` — View details of a caught Pokémon
- `pokedex` — List all caught Pokémon

### Notes

- This project is for educational purposes and interacts with the public [PokeAPI](https://pokeapi.co/).
- Make sure you have a stable internet connection to use all features.
- Tests are being added as time goes on. Please take this into consideration if you're planning on contributing to this project!

---

Feel free to open issues or contribute!
