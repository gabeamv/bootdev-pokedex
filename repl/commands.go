package repl

import (
	"bootdev-pokedex/internal/pokeapi"
	"bootdev-pokedex/internal/pokecache"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var Pokedex = make(map[string]pokeapi.Pokemon)

func commandExit(c *config, cache *pokecache.Cache, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config, cache *pokecache.Cache, args ...string) error {
	message := "Welcome to the Pokedex!\nUsage:\n\n"
	for _, cmd := range getCommands() {
		message += fmt.Sprintf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Print(message)
	return nil
}

func commandMap(c *config, cache *pokecache.Cache, args ...string) error {
	if c.Next == "" {
		return fmt.Errorf("you are already at the end of the location areas")
	}
	fullUrl := c.Next
	var locationAreas pokeapi.LocationAreas
	bytes, ok := cache.Get(fullUrl) // attempt to get the cached location areas
	if !ok {                        // if we cannot get the cached location areas, make request
		var err error
		bytes, err = pokeapi.Get(fullUrl)
		if err != nil {
			return fmt.Errorf("error getting location areas: %w", err)
		}
		cache.Add(fullUrl, bytes) // cache the data
	} else {
		fmt.Print("Cache is being used.\n\n")
	}
	err := json.Unmarshal(bytes, &locationAreas)
	if err != nil {
		return fmt.Errorf("error unmarshalling location areas: %w", err)
	}
	if fullUrl == DOMAIN+PATH_AREA_START {
		fmt.Printf("You are on the first page.\n\n")
	}

	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}
	c.Next = locationAreas.Next
	c.Previous = locationAreas.Previous
	fmt.Println(fullUrl)
	return nil
}

func commandMapb(c *config, cache *pokecache.Cache, args ...string) error {
	if c.Previous == "" {
		return fmt.Errorf("error. you are already at the start of the location areas")
	}
	fullUrl := c.Previous
	var locationAreas pokeapi.LocationAreas
	bytes, ok := cache.Get(fullUrl) // attempt to get the cached location areas
	if !ok {                        // if we cannot get the cached location areas, make request
		var err error
		bytes, err = pokeapi.Get(fullUrl)
		if err != nil {
			return fmt.Errorf("error getting location areas: %w", err)
		}
		cache.Add(fullUrl, bytes) // cache the data
	} else {
		fmt.Print("Cache is being used.\n\n")
	}
	err := json.Unmarshal(bytes, &locationAreas)
	if err != nil {
		return fmt.Errorf("error unmarshalling bytes to pokeapi.LocationAreas: %w", err)
	}
	if fullUrl == DOMAIN+PATH_AREA_START {
		fmt.Printf("You are on the first page.\n\n")
	}
	// Iterate through the list of maps and display the names of location areas
	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}
	c.Next = locationAreas.Next
	c.Previous = locationAreas.Previous
	fmt.Println(fullUrl)
	return nil
}

func commandExplore(c *config, cache *pokecache.Cache, args ...string) error {
	if len(args) == 1 {
		return fmt.Errorf("error. missing location area argument.")
	}
	area := args[1]
	fullUrl := DOMAIN + PATH_AREA + area
	bytes, ok := cache.Get(fullUrl)
	if !ok {
		var err error
		bytes, err = pokeapi.Get(fullUrl)
		if err != nil {
			return fmt.Errorf("error getting the specifc location area: %w", err)
		}
		cache.Add(fullUrl, bytes)
	} else {
		fmt.Print("Cache is being used.\n\n")
	}
	var locationArea pokeapi.LocationArea
	err := json.Unmarshal(bytes, &locationArea)
	if err != nil {
		return fmt.Errorf("error unmarshalling bytes to pokeapi.LocationArea: %w", err)
	}
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}

func commandCatch(c *config, cache *pokecache.Cache, args ...string) error {
	if len(args) == 1 {
		return fmt.Errorf("error: missing pokemon name argument")
	}
	pokemonName := args[1]
	fullUrl := DOMAIN + PATH_POKEMON + pokemonName
	bytes, ok := cache.Get(fullUrl)
	if !ok {
		var err error
		bytes, err = pokeapi.Get(fullUrl)
		if err != nil {
			return fmt.Errorf("error requesting the pokemon '%s': %w", pokemonName, err)
		}
		cache.Add(fullUrl, bytes)
	} else {
		fmt.Print("Cache being used.\n\n")
	}
	var pokemon pokeapi.Pokemon
	err := json.Unmarshal(bytes, &pokemon)
	if err != nil {
		return fmt.Errorf("error unmarshalling bytes to pokeapi.Pokemon")
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	target := pokemon.BaseExperience
	var catch int
	for i := 1; i < 4; i++ {
		catch += rand.Intn(pokeapi.MAX_ROLL + 1)
		waitOutput := strings.Repeat(".", i)
		fmt.Println(waitOutput)
		time.Sleep(time.Second)
	}
	if catch >= target { // if pokemon is caught...
		fmt.Printf("You have captured %s!\n", pokemonName)  // output that pokemon has been caught
		if _, captured := Pokedex[pokemonName]; !captured { // if first time capturing
			Pokedex[pokemonName] = pokemon // store in pokedex
			fmt.Printf("%s's data has been stored in your Pokedex.\n", pokemonName)
		}
		return nil
	}
	fmt.Printf("%s broke free...\n", pokemonName)
	return nil
}

func commandInspect(c *config, cache *pokecache.Cache, args ...string) error {
	if len(args) == 1 {
		return fmt.Errorf("error: missing pokemon name argument")
	}
	pokemonName := args[1]
	if pokemon, captured := Pokedex[pokemonName]; captured {
		output := fmt.Sprintf("Name: %v\nHeight: %v\nWeight: %v\nStats:\n", pokemon.Name, pokemon.Height, pokemon.Weight)
		for _, pokeStat := range pokemon.Stats {
			output += fmt.Sprintf("\t-%v: %v\n", pokeStat.Stat.Name, pokeStat.BaseStat)
		}
		output += "Types:\n"
		for _, pokeType := range pokemon.Types {
			output += fmt.Sprintf("\t-%v\n", pokeType.Type.Name)
		}
		fmt.Print(output)
	} else {
		fmt.Println("You have not caught that Pokemon...")
	}
	return nil
}

func commandPokedex(c *config, cache *pokecache.Cache, args ...string) error {
	output := "Your Pokedex:\n"
	for _, pokemon := range Pokedex {
		output += fmt.Sprintf("\t- %v\n", pokemon.Name)
	}
	fmt.Println(output)
	return nil
}

func getCommands() map[string]commands {
	var commandMap = map[string]commands{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Display the pokemon located in a specific area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a specific pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "View the data of a caught pokemon from the Pokedex",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View all the pokemon that have been caught",
			callback:    commandPokedex,
		},
	}
	return commandMap
}
