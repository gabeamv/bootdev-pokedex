package repl

import (
	"bootdev-pokedex/internal/pokecache"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	CACHE_INT = 10 * time.Second // reap the cache every 10 seconds
)

func Start() {
	commandMap := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	c := config{Previous: "", Next: fmt.Sprintf(DOMAIN + PATH_AREA_START)} // config struct to keep track of the previous and next locations to read.
	cache := pokecache.NewCache(CACHE_INT)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		args := CleanInput(text)
		command, ok := commandMap[args[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		var arg1 string
		if command.name == "explore" || command.name == "catch" || command.name == "inspect" {
			if len(args) < 2 {
				fmt.Println("missing second argument")
				continue
			}
			arg1 = args[1]
		}
		err := command.callback(&c, cache, arg1)
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
