package repl

import (
	"bootdev-pokedex/internal/pokeapi"
	"bootdev-pokedex/internal/pokecache"
	"encoding/json"
	"fmt"
	"os"
)

func commandExit(c *config, cache *pokecache.Cache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config, cache *pokecache.Cache) error {
	message := "Welcome to the Pokedex!\nUsage:\n\n"
	for _, cmd := range getCommands() {
		message += fmt.Sprintf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Print(message)
	return nil
}

func commandMap(c *config, cache *pokecache.Cache) error {
	if c.Next == "" {
		return fmt.Errorf("you are already at the end of the location areas")
	}
	fullUrl := c.Next
	var locationAreas pokeapi.LocationAreas
	bytes, ok := cache.Get(fullUrl) // attempt to get the cached location areas
	if !ok {                        // if we cannot get the cached location areas, make request
		var err error
		bytes, err = pokeapi.GetLocationAreas(fullUrl)
		if err != nil {
			return fmt.Errorf("error getting location areas: %w", err)
		}
		cache.Add(fullUrl, bytes) // cache the data
	}
	err := json.Unmarshal(bytes, &locationAreas)
	if err != nil {
		return fmt.Errorf("error unmarshalling location areas: %w", err)
	}
	if fullUrl == DOMAIN+"location-area" {
		fmt.Printf("You are on the first page.\n\n")
	}

	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}
	c.Next = locationAreas.Next
	c.Previous = locationAreas.Previous
	return nil
}

func mapb(c *config, cache *pokecache.Cache) error {
	if c.Previous == "" {
		return fmt.Errorf("error. you are already at the start of the location areas")
	}
	fullUrl := c.Previous
	var locationAreas pokeapi.LocationAreas
	bytes, ok := cache.Get(fullUrl) // attempt to get the cached location areas
	if !ok {                        // if we cannot get the cached location areas, make request
		var err error
		bytes, err = pokeapi.GetLocationAreas(fullUrl)
		if err != nil {
			return fmt.Errorf("error getting location areas: %w", err)
		}
	}
	err := json.Unmarshal(bytes, &locationAreas)
	if err != nil {
		return fmt.Errorf("error unmarshalling location areas: %w", err)
	}
	if fullUrl == DOMAIN+"location-area?offset=0&limit=20" {
		fmt.Printf("You are on the first page.\n\n")
	}
	// Iterate through the list of maps and display the names of location areas
	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}
	c.Next = locationAreas.Next
	c.Previous = locationAreas.Previous
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
			callback:    mapb,
		},
	}
	return commandMap
}
