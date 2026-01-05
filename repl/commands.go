package repl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	message := "Welcome to the Pokedex!\nUsage:\n\n"
	for _, cmd := range getCommands() {
		message += fmt.Sprintf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Print(message)
	return nil
}

func commandMap(c *config) error {
	// Standard get request for location areas.
	if c.Next == "" {
		return fmt.Errorf("you are already at the end of the location areas")
	}
	fullUrl := c.Next
	client := &http.Client{}
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return fmt.Errorf("error creating new get request for location areas: %w", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error getting location areas: %w", err)
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading all the bytes of location areas: %w", err)
	}
	// Get the entire json.
	var locationAreas struct {
		Count    int    `json:"count"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Results  []struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"results"`
	}
	err = json.Unmarshal(bytes, &locationAreas)
	if err != nil {
		return fmt.Errorf("error unmarshalling location areas: %w", err)
	}
	if fullUrl == DOMAIN+"location-area" {
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

func mapb(c *config) error {
	if c.Previous == "" {
		return fmt.Errorf("error. you are already at the start of the location areas")
	}
	fullUrl := c.Previous
	client := &http.Client{}
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return fmt.Errorf("error creating new get request for location areas: %w", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error getting location areas: %w", err)
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading all the bytes of location areas: %w", err)
	}
	// Get the entire json.
	var locationAreas struct {
		Count    int    `json:"count"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Results  []struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"results"`
	}
	err = json.Unmarshal(bytes, &locationAreas)
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
