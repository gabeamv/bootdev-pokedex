<!-- .github/copilot-instructions.md for bootdev-pokedex -->
# Quick guide for AI coding agents working on this repository

Purpose: help an AI agent get productive quickly in this very small Go project.

Key facts (from repository):
- Single-file Go project: `main.go` at the repo root. Currently the file is a placeholder.
- No `go.mod` or other Go modules found. No tests or other source directories exist yet.

Immediate goals for an agent
- Treat this as a small Go CLI or service called "bootdev-pokedex".
- Prefer minimal, idiomatic Go: create a `go.mod` before adding dependencies.

Commands and workflows
- Build: `go build -o bin/pokedex ./...` (or `go build -o bin/pokedex main.go` for the single-file case).
- Run locally: `go run main.go`.
- Initialize modules: `go mod init github.com/gabeamv/bootdev-pokedex` (do this before adding external packages).
- Lint/static checks: prefer `go vet` and `golangci-lint run` if the repo later adds a config.
- Tests: `go test ./...` (no tests present yet; add tests under `*_test.go`).

Project-specific patterns and conventions
- Keep things idiomatic and minimal: when this repo grows, follow standard Go layout (recommended):
  - `cmd/pokedex/main.go` for the CLI/service entrypoint
  - `internal/` for non-public packages
  - `pkg/` for public libraries (if needed)
- For now, update `main.go` directly for small changes; for larger features, create `cmd/` and move the entrypoint.
- Use explicit, small functions and prefer returning errors to panicking. Example: an HTTP server helper should return an error from Start() rather than log.Fatal.

Integration points and expectations
- No external APIs or databases are defined in the repo. If adding integrations (e.g., PokeAPI, DB), add clear configuration via environment variables and document them in `README.md`.
- When adding external integrations, include reproducible examples (env files or example flags) and small integration tests or mocks.

Examples from this repository
- `main.go` exists at the repo root; any entrypoint changes should preserve `package main` and implement `func main()`.

What to avoid
- Do not change the repo layout to a heavy framework immediately. Start minimal, keep changes small and well-tested.

Merge guidance (if an existing `.github/copilot-instructions.md` is present)
- Preserve any project-specific notes already present (commands, env vars, CI snippets). Supplement missing parts above rather than replace blindly.

If something is unclear, ask the human owner these quick questions:
1. Intended runtime: CLI, HTTP service, or library?
2. Preferred module path (if you should run `go mod init`)?
3. Any CI/tooling they want (GitHub Actions, golangci-lint config, etc.)?

Next steps for the agent after creating this file
- Run the commands above locally: initialize `go.mod` (with owner confirmation if desired), implement a minimal `main()` that prints "hello" or starts a tiny HTTP handler, add a basic test, and push as a small PR.

Contact/ownership
- Repository owner: `gabeamv` (from repo path). When in doubt, make minimal, reversible edits and request human review.

End of file.
