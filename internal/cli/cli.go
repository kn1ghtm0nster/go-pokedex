package cli

import (
	"github.com/kn1ghtm0nster/go-pokedex/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, []string) error
}

type Config struct {
	Next     string
	Previous string
	Cache    *pokecache.Cache
}

