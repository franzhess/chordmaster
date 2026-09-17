# Project Overview

Chordmaster is a Go terminal UI for exploring scales, chords, scale-derived triads, chord progressions, rhythm patterns, guitar riffs, guitar fretboard positions, chord fingerings, and generated audio playback.

Tech stack:
- Go 1.26.2
- Bubble Tea for the TUI update/render loop
- Lip Gloss for terminal styling
- Oto v2 for generated PCM audio playback

Current structure:
- `main.go`: process bootstrap only; initializes the single synth engine and starts Bubble Tea with `ui.NewModel`.
- `internal/music`: pure music theory domain: pitch classes, octave-bearing notes, scale patterns, concrete scales, spelling rules, chord formulas, chord lookup, and scale-derived chord schemas. It resolves enharmonic pitch-class names including double sharps and double flats. It must not import UI, synth, instrument, Bubble Tea, Lip Gloss, or Oto packages.
- `internal/instruments/guitar`: standard guitar fretboard rendering and octave-aware fret positions. Depends on `internal/music` for pitch classes and notes.
- `internal/instruments/guitar/chordshapes`: guitar chord fingering templates, transposition, lookup, parse/render, and validation. Depends on `internal/music` for pitch classes. Open/open-position coverage includes common C/D/E/F/G/A major shapes, B7, Fmaj7, Gmaj7, and Bdim for scale-triad coverage.
- `internal/instruments/piano`: metadata stub for standard piano range, A0 through C8. Depends only on `internal/music`.
- `internal/instruments/ukulele`: metadata stub for standard reentrant GCEA ukulele tuning. Depends only on `internal/music`.
- `internal/synth`: sine-wave PCM generation and Oto playback. Uses `internal/music` for note-name parsing so playback supports enharmonic spellings such as `E##` and `Bbb`. Uses 44100 Hz, mono, signed 16-bit little-endian samples, a single Oto context, generated in-memory audio, and cancellable tone/chord playback for UI-highlighted step playback.
- `internal/ui`: Bubble Tea model, update logic, responsive full-screen layouts, styles, and screen rendering. May depend on `music`, guitar packages, and `synth`. It owns guitar chord tab diagrams, chord-card flow layout, tab note/fingering mode, and highlighted scale note/chord playback state.
- `docs/architecture.md`: package responsibilities and dependency direction.
- `docs/spelling.md`: note spelling rules and future chord spelling guidance.
- `AGENTS.md`: persistent project rules and current behavior notes for future agent sessions.

Current verified package imports:
- `internal/music -> fmt, strconv, strings` only.
- `internal/instruments/guitar -> internal/music, fmt, strings`.
- `internal/instruments/guitar/chordshapes -> internal/music, fmt, strconv, strings`.
- `internal/instruments/piano -> internal/music`.
- `internal/instruments/ukulele -> internal/music`.
- `internal/synth -> internal/music` plus stdlib/Oto.
- `internal/ui -> internal/instruments/guitar, internal/instruments/guitar/chordshapes, internal/music, internal/synth, Bubble Tea, Lip Gloss`.

Current UI behavior:
- Dashboard uses full terminal dimensions, not a centered fixed viewport.
- Status bar shows selected instrument plus `Tabs: notes/fingers`; `f` toggles tab mode.
- Chord cards render name, note list, blank separator row, and a four-fret guitar tab.
- Chord tabs render strings high-to-low (`e B G D A E`); open-position shapes use `||`; shifted shapes show the starting fret number below the first displayed fret.
- Chord browser shows all available voicings for the selected chord. Scale/progression cards use the lowest-position voicing.
- On the scale screen, `p` plays/highlights scale notes step-by-step; pressing `p` during playback cancels. `c` plays/highlights scale-derived triads.

Latest verification completed successfully with `gofmt -w . && go mod tidy && go test ./...` plus package dependency listing.