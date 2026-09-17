package ui

import (
	"sort"

	"chordmaster/internal/instruments/guitar/chordshapes"
	"chordmaster/internal/music"
	"strconv"
	"strings"
)

type chordTabMode int

const (
	chordTabNotes chordTabMode = iota
	chordTabFingers
	chordTabVisibleFrets = 4
)

func chordTab(rootName string, variation music.ChordVariation, mode chordTabMode) (string, bool) {
	voicing, ok := chordVoicing(rootName, variation)
	if !ok {
		return "", false
	}

	return renderChordTab(voicing, mode), true
}

func chordVoicing(rootName string, variation music.ChordVariation) (chordshapes.Voicing, bool) {
	voicings, ok := chordVoicings(rootName, variation)
	if !ok {
		return chordshapes.Voicing{}, false
	}

	return voicings[0], true
}

func chordVoicings(rootName string, variation music.ChordVariation) ([]chordshapes.Voicing, bool) {
	root, err := chordshapes.ParseNoteName(rootName)
	if err != nil {
		return nil, false
	}

	quality, ok := chordShapeQuality(variation.Name)
	if !ok {
		return nil, false
	}

	voicings := chordshapes.GetVoicings(root, quality)
	if len(voicings) == 0 {
		return nil, false
	}
	// Non-chord screens show a single shape, so choose the lowest-position voicing.
	// The chord screen uses the full sorted list to expose alternate movable shapes.
	sort.SliceStable(voicings, func(i, j int) bool {
		return chordTabStartFret(voicings[i].Frets) < chordTabStartFret(voicings[j].Frets)
	})

	return voicings, true
}

func renderChordTab(voicing chordshapes.Voicing, mode chordTabMode) string {
	stringLabels := [6]string{"E", "A", "D", "G", "B", "e"}
	startFret := chordTabStartFret(voicing.Frets)
	openPosition := startFret == 1

	var b strings.Builder
	for stringIndex := len(stringLabels) - 1; stringIndex >= 0; stringIndex-- {
		fret := voicing.Frets[stringIndex]
		b.WriteString(mutedStyle.Render(stringLabels[stringIndex] + " "))
		b.WriteString(chordTabMarker(openStringMarker(stringIndex, fret, mode)))
		if openPosition {
			b.WriteString("||")
		} else {
			b.WriteString("|")
		}

		for tabFret := startFret; tabFret < startFret+chordTabVisibleFrets; tabFret++ {
			b.WriteString(chordTabCell(voicing, stringIndex, tabFret, mode))
			b.WriteString("|")
		}

		if stringIndex > 0 || !openPosition {
			b.WriteString("\n")
		}
	}
	if !openPosition {
		b.WriteString("    ")
		b.WriteString(strconv.Itoa(startFret))
	}

	return b.String()
}

func chordTabMarker(marker string) string {
	return marker
}

func chordTabStartFret(frets [6]int) int {
	// Open strings anchor the diagram at the nut; otherwise show the lowest fretted
	// note as the first of four visible frets and print that fret number below.
	startFret := 0
	for _, fret := range frets {
		if fret == 0 {
			return 1
		}
		if fret > 0 && (startFret == 0 || fret < startFret) {
			startFret = fret
		}
	}
	if startFret <= 1 {
		return 1
	}

	return startFret
}

func openStringMarker(stringIndex int, fret int, mode chordTabMode) string {
	if fret == -1 {
		return "X"
	}
	if fret == 0 && mode == chordTabNotes {
		return chordshapes.StandardTuning[stringIndex].String()
	}

	return " "
}

func chordTabCell(voicing chordshapes.Voicing, stringIndex int, tabFret int, mode chordTabMode) string {
	if voicing.Frets[stringIndex] != tabFret {
		return "---"
	}

	if mode == chordTabFingers {
		finger := voicing.Fingers[stringIndex]
		if finger == 0 {
			return "---"
		}
		return "-" + strconv.Itoa(finger) + "-"
	}

	note := chordshapes.TransposeNote(chordshapes.StandardTuning[stringIndex], tabFret).String()
	if len(note) == 1 {
		return "-" + note + "-"
	}

	return note + "-"
}

func chordShapeQuality(variationName string) (chordshapes.ChordQuality, bool) {
	switch variationName {
	case "Major":
		return chordshapes.Major, true
	case "Minor":
		return chordshapes.Minor, true
	case "Dominant7":
		return chordshapes.Dominant7, true
	case "Major7":
		return chordshapes.Major7, true
	case "Minor7":
		return chordshapes.Minor7, true
	case "Diminished":
		return chordshapes.Diminished, true
	case "Sus2":
		return chordshapes.Sus2, true
	case "Sus4", "Dominant7Sus4":
		return chordshapes.Sus4, true
	default:
		return "", false
	}
}
