package chordshapes

import (
	"fmt"
	"strconv"
	"strings"

	"chordmaster/internal/music"
)

func ParseNoteName(input string) (Note, error) {
	return music.ParsePitchClassName(input)
}

func ParseFrets(input string) ([6]int, error) {
	var frets [6]int
	input = strings.TrimSpace(input)

	if strings.Contains(input, " ") {
		parts := strings.Fields(input)
		if len(parts) != 6 {
			return frets, fmt.Errorf("expected 6 frets, got %d", len(parts))
		}
		for i, part := range parts {
			fret, err := parseFret(part)
			if err != nil {
				return frets, err
			}
			frets[i] = fret
		}
		return frets, nil
	}

	if len(input) != 6 {
		return frets, fmt.Errorf("expected 6 fret characters, got %d", len(input))
	}

	for i, part := range input {
		fret, err := parseFret(string(part))
		if err != nil {
			return frets, err
		}
		frets[i] = fret
	}

	return frets, nil
}

func parseFret(input string) (int, error) {
	if input == "x" || input == "X" {
		return -1, nil
	}

	fret, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	if fret < 0 || fret > 24 {
		return 0, fmt.Errorf("invalid fret %d", fret)
	}

	return fret, nil
}
