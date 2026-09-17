# Task Completion

Before considering code changes complete:
- Run `gofmt -w .` after editing Go files.
- Run `go mod tidy` and `go test ./...` to verify modules and tests.
- If dependencies changed, include resulting `go.mod` / `go.sum` changes.
- For package-boundary changes, run `go list -f '{{.ImportPath}} -> {{join .Imports ","}}' ./...` and confirm dependency direction.
- Confirm `internal/music` remains pure: no UI, synth, instrument, Bubble Tea, Lip Gloss, or Oto imports.
- Update `README.md`, `docs/architecture.md`, `AGENTS.md`, and Serena memories when package responsibilities or dependency direction change.
- For TUI behavior changes, manually run `go run .` when practical to check rendering and key handling.

Latest cleanup pass verification: `gofmt -w . && go mod tidy && go test ./...` passed for all packages.

When UI behavior changes, also check docs/architecture.md and the `project_overview` memory for updates to current UI behavior notes.