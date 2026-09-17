# Suggested Commands

Development:
- `go run .`: run the TUI app from the project root.
- `go test ./...`: run all tests.
- `gofmt -w .`: format Go source files.
- `go mod tidy`: clean module dependency metadata after dependency changes.
- `go list -f '{{.ImportPath}} -> {{join .Imports ","}}' ./...`: inspect package dependency direction.

Verification:
- Preferred full verification after code/doc cleanup touching Go files: `gofmt -w . && go mod tidy && go test ./...`.
- Confirm `internal/music` purity with package imports; it should list only stdlib imports such as `fmt`, `strconv`, and `strings`.

Notes:
- The project path is `/Users/franz/Projects/chordmaster`.
- The project directory is not currently a Git repository.
- Prefer project-aware search/read tools. If using shell commands, keep `workdir` at the project root and avoid broad searches outside the project.