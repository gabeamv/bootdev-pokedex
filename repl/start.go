package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gabeamv/bootdev-pokedex/internal/pokecache"
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
		err := command.callback(&c, cache, args...)
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
