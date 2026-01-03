package repl

import (
	"strings"
	"fmt"
	"bufio"
	"os"
)

type commands struct {
	name string
	description string
	callback func() error
}

func Start() {
	commandMap := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		command, ok := commandMap[text]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		command.callback()
	}
}

func CleanInput(text string) []string {
	text = strings.ToLower(text)
	words := strings.Split(text, " ")
	var res []string
	for _, word := range words {
		if word != "" {
			res = append(res, word)
		}
	}
	return res
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	message := fmt.Sprintf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range getCommands() {
		message += fmt.Sprintf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Print(message)
	return nil
}

func getCommands() map[string]commands {
	var commandMap = map[string]commands{
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
	}
	return commandMap
}
