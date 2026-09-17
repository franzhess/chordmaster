# Chord Progressions

Chordmaster resolves Roman-numeral progression patterns against the currently selected scale. It supports diatonic, modal, borrowed-chord, secondary-dominant, blues, and cadence examples.

The supported patterns are defined in `internal/music/progression.go`. The selected rhythm pattern cycles across the resolved chords during playback.

Examples:

- `I - V - vi - IV`: common major-key loop
- `ii - V - I`: jazz cadence
- `i - bVII - bVI - V`: Andalusian cadence
- `V/V - V - I`: secondary dominant resolution
