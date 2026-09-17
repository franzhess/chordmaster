package chordshapes

import (
	"strconv"
	"strings"
)

func RenderFrets(frets [6]int) string {
	parts := make([]string, 0, len(frets))
	for _, fret := range frets {
		if fret == -1 {
			parts = append(parts, "x")
			continue
		}
		parts = append(parts, strconv.Itoa(fret))
	}

	return strings.Join(parts, " ")
}

func RenderVoicing(voicing Voicing) string {
	return voicing.Name + " " + RenderFrets(voicing.Frets)
}

func RenderFingers(fingers [6]int) string {
	parts := make([]string, 0, len(fingers))
	for _, finger := range fingers {
		if finger == 0 {
			parts = append(parts, "-")
			continue
		}
		parts = append(parts, strconv.Itoa(finger))
	}

	return strings.Join(parts, " ")
}
