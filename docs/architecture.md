# Architecture

Chordmaster is split into small internal packages with explicit responsibilities.

## Packages

- `cmd root` (`main.go`): starts the app. It creates the single audio engine and passes it into the UI model.
- `internal/music`: pure music-domain logic. It owns pitch classes, octave-bearing notes, spelling, scales, chord formulas, chord progressions, and generic riff notation/timing. It should not import Bubble Tea, Lip Gloss, Oto, or instrument-specific packages.
- `internal/instruments/guitar`: guitar-specific rendering, including fretboard rendering and compiled original guitar riff exercises.
- `internal/instruments/guitar/chordshapes`: guitar chord shape templates, transposition, parsing, compact frets/fingers rendering, lookup, and validation.
- `internal/instruments/piano`: piano instrument stubs and standard keyboard range metadata.
- `internal/instruments/ukulele`: ukulele instrument stubs and standard GCEA tuning metadata.
- `internal/synth`: low-level sound generation and playback. It should not know about UI screens or domain navigation state.
- `internal/ui`: Bubble Tea presentation logic. It can depend on `music`, `instruments`, and `synth`, but those packages should not depend on it.

## Music Responsibilities

- Pitch classes and enharmonic names.
- Octave-bearing notes.
- Scale patterns and concrete scale construction.
- Scale spelling rules.
- Chord formulas, chord construction, and chord lookup by notes.
- Scale-derived chord schemas such as triads and 7th chords.
- Chord progression parsing and resolution.
- Generic riff notation parsing and timing conversion.

## Guitar Responsibilities

- Standard guitar tuning and fretboard rendering.
- Compiled original guitar riff exercises.
- Movable CAGED/barre chord shape templates.
- Open chord shape templates.
- Shape transposition, parsing, rendering, and validation.

## Instrument Stub Responsibilities

- Standard piano range metadata.
- Standard ukulele tuning metadata.
- No UI or synth dependencies.

## UI Responsibilities

- Screen state, cursor movement, and responsive layout composition.
- Splash, menu, chord browser, scale browser, progressions, and riffs rendering.
- Guitar chord tab diagrams in note or fingering mode. The chord browser shows every available voicing for the selected chord; scale and progression cards use the lowest-position voicing.
- Highlighted step playback for scale notes and scale-derived chords.
- Keyboard shortcuts.
- Calling `synth.Engine` when the user presses playback shortcuts.

## Synth Responsibilities

- One Oto context per process.
- Sine-wave PCM generation.
- Domain note-name parsing for playback, including enharmonic double sharps and double flats.
- Envelope handling.
- Tone, chord, and scale playback.
- Cancellable tone/chord playback used by the UI for highlighted step playback.

## Current UI Layout Rules

- The dashboard uses the full terminal width/height rather than a centered fixed viewport.
- Main menu options are centered in the dashboard body.
- Scale, chord, and progression screens keep selection metadata in compact context sections with counts next to selectable values.
- Chord cards render as chord name, note list, blank separator row, and a four-fret tab diagram.
- Chord tab diagrams render high-to-low strings (`e B G D A E`). Open-position shapes use `||`; shifted shapes show the starting fret below the first displayed fret.
- Flow layouts measure rendered Lip Gloss width, so rows wrap correctly for wider shifted chord shapes.

## Dependency Direction

`main` -> `ui` -> `music`

`ui` -> `instruments/guitar`

`instruments/guitar` -> `music`

`instruments/guitar/chordshapes` -> `music`

`instruments/piano` -> `music`

`instruments/ukulele` -> `music`

`main` -> `synth`

`synth` -> `music`

`ui` -> `synth`

`music` has no application-layer or instrument-specific dependencies.
