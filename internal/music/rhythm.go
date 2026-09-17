package music

import "fmt"

type RhythmPattern struct {
	Name  string
	Units []int
}

func SupportedRhythmPatterns() []RhythmPattern {
	return deduplicateRhythmPatterns([]RhythmPattern{
		{Name: "Whole Measure Chords", Units: []int{4, 4, 4, 4}},
		{Name: "Two-Bar Chords", Units: []int{8, 8, 8, 8}},
		{Name: "Half-Measure Chords", Units: []int{2, 2, 2, 2}},
		{Name: "Quarter-Measure Chords", Units: []int{1, 1, 1, 1}},
		{Name: "Pop Split Bar", Units: []int{2, 2, 4}},
		{Name: "Delayed Resolution", Units: []int{2, 4, 2}},
		{Name: "Long Tonic Ending", Units: []int{2, 3, 3}},
		{Name: "Anticipated Chorus Feel", Units: []int{2, 2, 3, 1}},
		{Name: "Power Chord Sustain", Units: []int{4, 4, 2}},
		{Name: "Driving Rock Pattern", Units: []int{2, 2, 2, 4}},
		{Name: "Heavy Ending", Units: []int{2, 2, 4, 8}},
		{Name: "Quick Change Blues", Units: []int{2, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4}},
		{Name: "Shuffle Hold Pattern", Units: []int{8, 8, 4, 4}},
		{Name: "ii-V-I", Units: []int{2, 2, 4}},
		{Name: "Bebop Changes", Units: []int{1, 2, 1, 2, 1, 2, 1, 2}},
		{Name: "Ballad Jazz", Units: []int{4, 4, 8}},
		{Name: "Modal Jazz", Units: []int{16, 16, 16, 16}},
		{Name: "Tension Build", Units: []int{8, 4, 2, 1}},
		{Name: "Suspense Pattern", Units: []int{4, 4, 4, 8}},
		{Name: "Build-Up Compression", Units: []int{8, 4, 2, 1}},
		{Name: "Drop Sustain", Units: []int{8, 8, 8, 4}},
		{Name: "Syncopated Groove", Units: []int{3, 1, 1, 3}},
		{Name: "Standard Waltz", Units: []int{3, 3, 3, 3}},
		{Name: "Long Waltz Phrase", Units: []int{3, 3, 6}},
		{Name: "Delayed Cadence", Units: []int{2, 2, 2, 8}},
		{Name: "Sudden Resolution", Units: []int{6, 1, 1}},
		{Name: "Turnaround Cycle", Units: []int{1, 1, 1, 4}},
		{Name: "Expanding Durations", Units: []int{1, 2, 4, 8}},
		{Name: "Funk Archetype", Units: []int{2, 1, 2, 3}},
	})
}

func deduplicateRhythmPatterns(patterns []RhythmPattern) []RhythmPattern {
	seen := make(map[string]bool, len(patterns))
	unique := make([]RhythmPattern, 0, len(patterns))

	for _, pattern := range patterns {
		units := reduceLoopUnits(pattern.Units)
		key := rhythmPatternKey(units)
		if seen[key] {
			continue
		}

		seen[key] = true
		unique = append(unique, RhythmPattern{Name: pattern.Name, Units: units})
	}

	return unique
}

func reduceLoopUnits(units []int) []int {
	for size := 1; size <= len(units)/2; size++ {
		if len(units)%size != 0 {
			continue
		}

		matches := true
		for i := size; i < len(units); i++ {
			if units[i] != units[i%size] {
				matches = false
				break
			}
		}
		if matches {
			return append([]int(nil), units[:size]...)
		}
	}

	return append([]int(nil), units...)
}

func rhythmPatternKey(units []int) string {
	key := ""
	for _, unit := range units {
		key += fmt.Sprintf("/%d", unit)
	}
	return key
}

func RhythmPatternLabel(pattern RhythmPattern) string {
	label := ""
	for _, unit := range pattern.Units {
		label += fmt.Sprintf("%d", unit)
	}
	return label
}
