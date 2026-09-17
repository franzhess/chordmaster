# Agent Notes

## Project

- Chordmaster is a Go Bubble Tea terminal app for exploring music theory, guitar fretboards/chord shapes, progressions, riffs, and generated audio playback.
- Run from the project root with `go run .`.
- Verify changes with `gofmt -w .`, `go mod tidy`, and `go test ./...`.
- Use git only when the user asks for git-oriented work such as reviewing changes, committing, or opening a PR.

## Architecture Rules

- `internal/music` is the pure domain package for music theory only.
- `internal/music` must not import UI, synth, instrument, Bubble Tea, Lip Gloss, or Oto packages.
- `internal/music` should stay focused on pitch classes, octave-bearing notes, spelling, scale patterns, concrete scales, chord formulas, chord lookup, scale-derived chord schemas, chord progressions, rhythm patterns, and generic riff notation/timing.
- `internal/instruments/guitar` owns standard fretboard rendering, octave-aware fret positions, and compiled guitar riff catalogs.
- `internal/instruments/guitar/chordshapes` owns guitar chord fingering templates, transposition, lookup, parse/render, and validation.
- `internal/instruments/piano` and `internal/instruments/ukulele` are currently metadata stubs; keep them free of UI and synth dependencies.
- `internal/synth` owns PCM generation and Oto playback; it may use `internal/music` for note-name parsing but should not know about UI screens or navigation state.
- `internal/ui` owns Bubble Tea model state, keyboard handling, layout, styles, and screen rendering. It may depend on `music`, guitar packages, and `synth`.

## Current Behavior

- Main menu order is `Scales`, `Chords`, `Progressions`, `Riffs`.
- `p` plays the selected scale, chord, progression, or riff; pressing `p` during highlighted scale playback cancels it.
- `c` plays scale-derived chords on the scale screen.
- `f` toggles guitar chord tabs between note names and finger numbers.
- `[` / `]` changes progression root; `,` / `.` changes progression rhythm.
- Scales play ascending at 400ms per note.
- Chords play simultaneously for 1200ms.
- Audio uses generated 44100 Hz mono signed 16-bit little-endian PCM sine waves with 10ms fade-in and 20ms fade-out.
- The guitar fretboard displays strings high-to-low in the UI (`e B G D A E`) while chord shape fret arrays are low-to-high (`E A D G B e`).
- Chord-shape frets use `-1` for muted strings, `0` for open strings, and positive values for fretted notes.

## Working Notes

- Prefer small, direct refactors over introducing abstractions.
- Keep package dependency direction explicit and update `docs/architecture.md` when package boundaries change.
- If a future change needs shared note, pitch-class, or spelling logic, put it in `internal/music`; if it is guitar fretboard/chord-shape infrastructure, keep it under `internal/instruments/guitar`.
