# Style And Conventions

- Standard Go formatting via `gofmt`.
- Use concise lower-case names for internal types/functions unless export is required across packages.
- TUI follows Bubble Tea conventions: `model` type with `Init`, `Update`, and `View` methods.
- Styling is defined with package-level Lip Gloss styles.
- UI layout helpers should measure rendered text with `lipgloss.Width` when arranging chord/tab cards; shifted chord shapes may be wider than open-position shapes.
- Comments should be reserved for non-obvious behavior such as async playback session cancellation, measured flow wrapping, and guitar tab windowing.
- Keep changes minimal and avoid introducing abstraction until there is reuse.
- Preserve package boundaries: `internal/music` is pure music theory and must stay independent of UI, synth, and instrument packages.
- Put shared note, pitch-class, spelling, scale, or chord concepts in `internal/music`.
- Put guitar fretboard and chord-shape infrastructure under `internal/instruments/guitar`.
- Keep `internal/instruments/piano` and `internal/instruments/ukulele` as metadata-only stubs until app behavior needs them; no UI or synth dependencies there.
- Tests use the standard `testing` package and exercise model/domain behavior directly.
- Use git only when the user asks for git-oriented work such as reviewing changes, committing, or opening a PR.