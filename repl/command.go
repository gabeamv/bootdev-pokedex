package repl

import (
	"bootdev-pokedex/internal/pokecache"
)

type config struct {
	Previous string
	Next     string
}

type commands struct {
	name        string
	description string
	callback    func(*config, *pokecache.Cache, string) error
}
