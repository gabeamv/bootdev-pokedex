package test

import (
	"bootdev-pokedex/repl"
	"testing"

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

func TestCaching(t *testing.T) {

}
