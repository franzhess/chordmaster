package music

import "strings"

type ChordProgression struct {
	Name    string
	Symbols []string
}

type ProgressionChord struct {
	Symbol    string
	Name      string
	NoteNames []string
}

func SupportedChordProgressions() []ChordProgression {
	return []ChordProgression{
		{Name: "Authentic Cadence", Symbols: []string{"V", "I"}},
		{Name: "Plagal Cadence", Symbols: []string{"IV", "I"}},
		{Name: "Half Cadence", Symbols: []string{"I", "V"}},
		{Name: "Deceptive Cadence", Symbols: []string{"V", "vi"}},
		{Name: "I-IV-V", Symbols: []string{"I", "IV", "V"}},
		{Name: "I-V-vi-IV", Symbols: []string{"I", "V", "vi", "IV"}},
		{Name: "I-vi-IV-V", Symbols: []string{"I", "vi", "IV", "V"}},
		{Name: "I-IV-vi-V", Symbols: []string{"I", "IV", "vi", "V"}},
		{Name: "I-V-IV", Symbols: []string{"I", "V", "IV"}},
		{Name: "I-IV-I-V", Symbols: []string{"I", "IV", "I", "V"}},
		{Name: "I-ii-V-I", Symbols: []string{"I", "ii", "V", "I"}},
		{Name: "I-vi-ii-V", Symbols: []string{"I", "vi", "ii", "V"}},
		{Name: "vi-IV-I-V", Symbols: []string{"vi", "IV", "I", "V"}},
		{Name: "vi-V-IV-V", Symbols: []string{"vi", "V", "IV", "V"}},
		{Name: "I-iii-IV-V", Symbols: []string{"I", "iii", "IV", "V"}},
		{Name: "I-iii-vi-ii-V-I", Symbols: []string{"I", "iii", "vi", "ii", "V", "I"}},
		{Name: "i-iv-v", Symbols: []string{"i", "iv", "v"}},
		{Name: "i-VI-VII", Symbols: []string{"i", "VI", "VII"}},
		{Name: "i-VII-VI", Symbols: []string{"i", "VII", "VI"}},
		{Name: "i-VI-III-VII", Symbols: []string{"i", "VI", "III", "VII"}},
		{Name: "i-iv-VII", Symbols: []string{"i", "iv", "VII"}},
		{Name: "i-VII-VI-VII", Symbols: []string{"i", "VII", "VI", "VII"}},
		{Name: "i-III-VII-VI", Symbols: []string{"i", "III", "VII", "VI"}},
		{Name: "i-iv-V", Symbols: []string{"i", "iv", "V"}},
		{Name: "i-VI-V", Symbols: []string{"i", "VI", "V"}},
		{Name: "ii°-V-i", Symbols: []string{"ii°", "V", "i"}},
		{Name: "i-V-i", Symbols: []string{"i", "V", "i"}},
		{Name: "Dorian i-IV", Symbols: []string{"i", "IV"}},
		{Name: "Dorian i-VII", Symbols: []string{"i", "VII"}},
		{Name: "Mixolydian I-bVII", Symbols: []string{"I", "bVII"}},
		{Name: "Mixolydian I-bVII-IV", Symbols: []string{"I", "bVII", "IV"}},
		{Name: "Lydian I-II", Symbols: []string{"I", "II"}},
		{Name: "Phrygian i-bII", Symbols: []string{"i", "bII"}},
		{Name: "12-Bar Blues", Symbols: []string{"I", "I", "I", "I", "IV", "IV", "I", "I", "V", "IV", "I", "I"}},
		{Name: "I-IV-I-I", Symbols: []string{"I", "IV", "I", "I"}},
		{Name: "Jazz Blues Turnaround", Symbols: []string{"I7", "IV7", "I7", "vi7", "ii7", "V7", "I7", "V7"}},
		{Name: "vi-ii-V-I", Symbols: []string{"vi", "ii", "V", "I"}},
		{Name: "iii-vi-ii-V-I", Symbols: []string{"iii", "vi", "ii", "V", "I"}},
		{Name: "vii°-iii-vi-ii-V-I", Symbols: []string{"vii°", "iii", "vi", "ii", "V", "I"}},
		{Name: "I-vi-ii-V", Symbols: []string{"I", "vi", "ii", "V"}},
		{Name: "iii-vi-ii-V", Symbols: []string{"iii", "vi", "ii", "V"}},
		{Name: "I-bIII7-ii-bII7", Symbols: []string{"I", "bIII7", "ii", "bII7"}},
		{Name: "Andalusian Cadence", Symbols: []string{"i", "bVII", "bVI", "V"}},
		{Name: "I-bVII-IV", Symbols: []string{"I", "bVII", "IV"}},
		{Name: "I-V-bVII-IV", Symbols: []string{"I", "V", "bVII", "IV"}},
		{Name: "I-bIII-IV", Symbols: []string{"I", "bIII", "IV"}},
		{Name: "IV-I-V-vi", Symbols: []string{"IV", "I", "V", "vi"}},
		{Name: "ii-V-I", Symbols: []string{"ii", "V", "I"}},
		{Name: "Imaj7-iii7-vi7-ii7", Symbols: []string{"Imaj7", "iii7", "vi7", "ii7"}},
		{Name: "IVmaj7-iii7-vi7-V7", Symbols: []string{"IVmaj7", "iii7", "vi7", "V7"}},
		{Name: "i-bVI-bIII-bVII", Symbols: []string{"i", "bVI", "bIII", "bVII"}},
		{Name: "i-V-bVI-IV", Symbols: []string{"i", "V", "bVI", "IV"}},
		{Name: "IV-#iv°-I", Symbols: []string{"IV", "#iv°", "I"}},
		{Name: "V/V-V-I", Symbols: []string{"V/V", "V", "I"}},
		{Name: "V/ii-ii", Symbols: []string{"V/ii", "ii"}},
		{Name: "V/vi-vi", Symbols: []string{"V/vi", "vi"}},
		{Name: "I-iv-I", Symbols: []string{"I", "iv", "I"}},
		{Name: "I-bVI-IV", Symbols: []string{"I", "bVI", "IV"}},
		{Name: "IV-V-I", Symbols: []string{"IV", "V", "I"}},
		{Name: "bII-I", Symbols: []string{"bII", "I"}},
		{Name: "V7-i", Symbols: []string{"V7", "i"}},
		{Name: "Major Key I", Symbols: []string{"I"}},
		{Name: "Major Key ii", Symbols: []string{"ii"}},
		{Name: "Major Key iii", Symbols: []string{"iii"}},
		{Name: "Major Key IV", Symbols: []string{"IV"}},
		{Name: "Major Key V", Symbols: []string{"V"}},
		{Name: "Major Key vi", Symbols: []string{"vi"}},
		{Name: "Major Key vii°", Symbols: []string{"vii°"}},
		{Name: "Natural Minor i", Symbols: []string{"i"}},
		{Name: "Natural Minor ii°", Symbols: []string{"ii°"}},
		{Name: "Natural Minor III", Symbols: []string{"III"}},
		{Name: "Natural Minor iv", Symbols: []string{"iv"}},
		{Name: "Natural Minor v", Symbols: []string{"v"}},
		{Name: "Natural Minor VI", Symbols: []string{"VI"}},
		{Name: "Natural Minor VII", Symbols: []string{"VII"}},
		{Name: "Harmonic Minor i", Symbols: []string{"i"}},
		{Name: "Harmonic Minor ii°", Symbols: []string{"ii°"}},
		{Name: "Harmonic Minor III+", Symbols: []string{"III+"}},
		{Name: "Harmonic Minor iv", Symbols: []string{"iv"}},
		{Name: "Harmonic Minor V", Symbols: []string{"V"}},
		{Name: "Harmonic Minor VI", Symbols: []string{"VI"}},
		{Name: "Harmonic Minor vii°", Symbols: []string{"vii°"}},
		{Name: "Major Key Imaj7", Symbols: []string{"Imaj7"}},
		{Name: "Major Key ii7", Symbols: []string{"ii7"}},
		{Name: "Major Key iii7", Symbols: []string{"iii7"}},
		{Name: "Major Key IVmaj7", Symbols: []string{"IVmaj7"}},
		{Name: "Major Key V7", Symbols: []string{"V7"}},
		{Name: "Major Key vi7", Symbols: []string{"vi7"}},
		{Name: "Major Key viiø7", Symbols: []string{"viiø7"}},
		{Name: "Natural Minor i7", Symbols: []string{"i7"}},
		{Name: "Natural Minor iiø7", Symbols: []string{"iiø7"}},
		{Name: "Natural Minor IIImaj7", Symbols: []string{"IIImaj7"}},
		{Name: "Natural Minor iv7", Symbols: []string{"iv7"}},
		{Name: "Natural Minor v7", Symbols: []string{"v7"}},
		{Name: "Natural Minor VImaj7", Symbols: []string{"VImaj7"}},
		{Name: "Natural Minor VII7", Symbols: []string{"VII7"}},
	}
}

func ResolveChordProgression(scale Scale, progression ChordProgression) []ProgressionChord {
	chords := make([]ProgressionChord, 0, len(progression.Symbols))
	for _, symbol := range progression.Symbols {
		chord, ok := resolveProgressionSymbol(scale, symbol)
		if ok {
			chords = append(chords, chord)
		}
	}

	return chords
}

func resolveProgressionSymbol(scale Scale, symbol string) (ProgressionChord, bool) {
	rootName, variation, ok := progressionSymbolRootAndVariation(scale, symbol)
	if !ok {
		return ProgressionChord{}, false
	}

	notes := chordNoteNamesFromRoot(rootName, variation)
	return ProgressionChord{
		Symbol:    symbol,
		Name:      rootName + variation.Suffix,
		NoteNames: notes,
	}, true
}

func progressionSymbolRootAndVariation(scale Scale, symbol string) (string, ChordVariation, bool) {
	normalized := normalizeProgressionSymbol(symbol)
	if strings.HasPrefix(normalized, "V/") {
		targetRoot, ok := progressionRootName(scale, strings.TrimPrefix(normalized, "V/"))
		if !ok {
			return "", ChordVariation{}, false
		}

		target, err := ParsePitchClassName(targetRoot)
		if err != nil {
			return "", ChordVariation{}, false
		}

		root := TransposePitchClass(target, 7)
		return root.PreferredName(PreferSharps), chordVariation("major"), true
	}

	rootName, ok := progressionRootName(scale, normalized)
	if !ok {
		return "", ChordVariation{}, false
	}

	return rootName, progressionVariation(normalized), true
}

func progressionRootName(scale Scale, symbol string) (string, bool) {
	degree, modifier, ok := progressionDegree(symbol)
	if !ok || len(scale.NoteNames) == 0 {
		return "", false
	}

	if modifier == 0 && len(scale.NoteNames) == 7 {
		return scale.NoteNames[degree-1], true
	}

	rootSemitone, ok := scale.Root.Semitone()
	if !ok {
		return "", false
	}

	semitone := rootSemitone + majorScaleDegreeSemitones(degree) + modifier
	preference := PreferSharps
	if modifier < 0 {
		preference = PreferFlats
	}

	return PitchClassBySemitone(semitone).PreferredName(preference), true
}

func progressionDegree(symbol string) (int, int, bool) {
	modifier := 0
	for strings.HasPrefix(symbol, "b") || strings.HasPrefix(symbol, "#") {
		switch symbol[0] {
		case 'b':
			modifier--
		case '#':
			modifier++
		}
		symbol = symbol[1:]
	}

	for _, candidate := range []struct {
		roman string
		value int
	}{
		{roman: "vii", value: 7},
		{roman: "VII", value: 7},
		{roman: "vi", value: 6},
		{roman: "VI", value: 6},
		{roman: "iv", value: 4},
		{roman: "IV", value: 4},
		{roman: "iii", value: 3},
		{roman: "III", value: 3},
		{roman: "ii", value: 2},
		{roman: "II", value: 2},
		{roman: "v", value: 5},
		{roman: "V", value: 5},
		{roman: "i", value: 1},
		{roman: "I", value: 1},
	} {
		if strings.HasPrefix(symbol, candidate.roman) {
			return candidate.value, modifier, true
		}
	}

	return 0, 0, false
}

func progressionVariation(symbol string) ChordVariation {
	majorQuality := strings.ContainsAny(symbol, "IVX")
	switch {
	case strings.Contains(symbol, "ø7"):
		return chordVariation("half-diminished7")
	case strings.Contains(symbol, "°") && strings.Contains(symbol, "7"):
		return chordVariation("diminished7")
	case strings.Contains(symbol, "°"):
		return chordVariation("diminished")
	case strings.Contains(symbol, "+"):
		return chordVariation("augmented")
	case strings.Contains(symbol, "maj7"):
		return chordVariation("major7")
	case strings.Contains(symbol, "7") && majorQuality:
		return chordVariation("dominant7")
	case strings.Contains(symbol, "7"):
		return chordVariation("minor7")
	case majorQuality:
		return chordVariation("major")
	default:
		return chordVariation("minor")
	}
}

func chordVariation(name string) ChordVariation {
	for _, variation := range DefaultChordCatalog().Variations {
		if strings.EqualFold(variation.Name, strings.ReplaceAll(name, "-", "")) {
			return variation
		}
	}

	return DefaultChordCatalog().Variations[0]
}

func chordNoteNamesFromRoot(rootName string, variation ChordVariation) []string {
	root, err := ParsePitchClassName(rootName)
	if err != nil {
		return nil
	}

	rootSemitone, ok := root.Semitone()
	if !ok || rootName == "" {
		return nil
	}

	rootLetterIndex := letterIndex(rootName[0])
	if rootLetterIndex < 0 {
		return nil
	}

	names := make([]string, 0, len(variation.Formula))
	for _, degree := range variation.Formula {
		semitone := rootSemitone + majorScaleDegreeSemitones(degree.Degree) + modifierSemitones(degree.Modifier)
		letter := scaleLetters[(rootLetterIndex+degree.Degree-1)%len(scaleLetters)]
		names = append(names, spellWithLetter(letter, (semitone%12+12)%12))
	}

	return names
}

func normalizeProgressionSymbol(symbol string) string {
	symbol = strings.TrimSpace(symbol)
	symbol = strings.ReplaceAll(symbol, "♭", "b")
	symbol = strings.ReplaceAll(symbol, "♯", "#")
	return symbol
}
