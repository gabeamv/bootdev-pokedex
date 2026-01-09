package test

import (
	"bootdev-pokedex/internal/pokecache"
	"bootdev-pokedex/repl"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestCleanInput(t *testing.T) {

	cases := map[string]struct { // map of cases, name : testing data
		input string
		want  []string
	}{
		"no-whitespace":       {input: "helloworld", want: []string{"helloworld"}},
		"trailing-whitespace": {input: " helloworld", want: []string{"helloworld"}},
		"leading-whitespace":  {input: "helloworld ", want: []string{"helloworld"}},
		"word":                {input: "hello world", want: []string{"hello", "world"}},
		"capital":             {input: "hello World", want: []string{"hello", "world"}},
	}

	for name, tc := range cases { // iterate through the test cases
		t.Run(name, func(t *testing.T) { // use testing.T.Run to run subtests
			got := repl.CleanInput(tc.input) // test the test case
			diff := cmp.Diff(tc.want, got)   // compare with the actual output
			if diff != "" {                  // if there is a difference
				t.Fatal(diff) // output the difference
			}
		})
	}
}

func TestNewCache(t *testing.T) {
	cases := map[string]struct {
		input time.Duration
		want  bool
	}{
		"new-cache-correct": {input: 10 * time.Second, want: true},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			cache := pokecache.NewCache(tc.input)
			got := reflect.TypeOf(cache) == reflect.TypeOf(&pokecache.Cache{}) // check if the newly created cache is of correct type
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
