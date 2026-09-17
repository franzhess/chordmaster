package chordshapes

func NotesForVoicing(voicing Voicing, tuning [6]Note) [6]*Note {
	var notes [6]*Note
	for stringIndex, fret := range voicing.Frets {
		if fret == -1 {
			continue
		}

		note := TransposeNote(tuning[stringIndex], fret)
		notes[stringIndex] = &note
	}

	return notes
}

func ValidateVoicing(voicing Voicing, chordTones []Note, tuning [6]Note) bool {
	noteSet := make(map[Note]bool, len(chordTones))
	for _, tone := range chordTones {
		noteSet[normalizeNote(tone)] = true
	}

	seen := make(map[Note]bool, len(chordTones))
	for _, note := range NotesForVoicing(voicing, tuning) {
		if note == nil {
			continue
		}

		normalized := normalizeNote(*note)
		if !noteSet[normalized] {
			return false
		}
		seen[normalized] = true
	}

	for tone := range noteSet {
		if !seen[tone] {
			return false
		}
	}

	return true
}
