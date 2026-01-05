package repl

type commands struct {
	name        string
	description string
	callback    func(*config) error
}
