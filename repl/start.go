package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Start() {
	commandMap := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	c := config{Previous: "", Next: fmt.Sprintf(DOMAIN + "location-area")} // config struct to keep track of the previous and next locations to read.
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		command, ok := commandMap[text]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(&c)
		if err != nil {
			fmt.Println(err)
		}
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
