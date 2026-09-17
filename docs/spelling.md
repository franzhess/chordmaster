# Note Spelling Rules

## Current Scale Behavior

- If a key spelling is known, scales prefer that spelling family.
- Sharp keys prefer sharps.
- Flat keys prefer flats.
- Ambiguous pitch classes default to sharps until the UI supports selecting a specific enharmonic key spelling.
- Seven-note scales spell one of each letter name in order: A, B, C, D, E, F, G.
- Accidentals are chosen after the target letter is selected so the semitone interval pattern remains correct.
- Pitch-class parsing resolves enharmonic equivalents for natural notes, single accidentals, double sharps, and double flats.
- Non-seven-note scales currently use the key spelling preference rather than forcing one of each letter.

## Future Chord Behavior

- Chords should choose note letters from chord degrees first.
- Chord alterations should then apply `#`, `b`, or `bb` to those selected letters.
- Chord spelling should use the surrounding key context when available.

## No Context Behavior

- Ascending chromatic movement should prefer sharps.
- Descending chromatic movement should prefer flats.
- Otherwise, default to sharps or a configurable preference.

## Enharmonic Parsing Examples

- `E##` resolves to `F# / Gb`.
- `B##` resolves to `C# / Db`.
- `Dbb` resolves to `C`.
- `Bbb` resolves to `A`.
