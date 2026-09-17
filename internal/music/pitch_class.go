package music

import (
	"fmt"
	"strings"
)

type SpellingPreference int

const (
	PreferSharps SpellingPreference = iota
	PreferFlats
)

type PitchClass struct {
	SharpName string
	FlatName  string
}

type Note struct {
	PitchClass PitchClass
	Octave     int
}

var (
	C  = PitchClass{SharpName: "C"}
	Cs = PitchClass{SharpName: "C#", FlatName: "Db"}
	D  = PitchClass{SharpName: "D"}
	Ds = PitchClass{SharpName: "D#", FlatName: "Eb"}
	E  = PitchClass{SharpName: "E"}
	F  = PitchClass{SharpName: "F"}
	Fs = PitchClass{SharpName: "F#", FlatName: "Gb"}
	G  = PitchClass{SharpName: "G"}
	Gs = PitchClass{SharpName: "G#", FlatName: "Ab"}
	A  = PitchClass{SharpName: "A"}
	As = PitchClass{SharpName: "A#", FlatName: "Bb"}
	B  = PitchClass{SharpName: "B"}
)

func (p PitchClass) Name() string {
	if p.FlatName == "" {
		return p.SharpName
	}

	return p.SharpName + " / " + p.FlatName
}

func (p PitchClass) String() string {
	return p.PreferredName(PreferSharps)
}

func (p PitchClass) PreferredName(preference SpellingPreference) string {
	if preference == PreferFlats && p.FlatName != "" {
		return p.FlatName
	}

	if p.SharpName != "" {
		return p.SharpName
	}

	return p.FlatName
}

func ChromaticPitchClasses() []PitchClass {
	return []PitchClass{
		C,
		Cs,
		D,
		Ds,
		E,
		F,
		Fs,
		G,
		Gs,
		A,
		As,
		B,
	}
}

func ParsePitchClassName(input string) (PitchClass, error) {
	name := strings.TrimSpace(input)
	if pitchClass, ok := pitchClassAliases[name]; ok {
		return pitchClass, nil
	}

	for _, pitchClass := range ChromaticPitchClasses() {
		if name == pitchClass.SharpName || name == pitchClass.FlatName {
			return pitchClass, nil
		}
	}

	if pitchClass, ok := parseAlteredPitchClassName(name); ok {
		return pitchClass, nil
	}

	return PitchClass{}, fmt.Errorf("unknown pitch class %q", input)
}

func (p PitchClass) Semitone() (int, bool) {
	for i, pitchClass := range ChromaticPitchClasses() {
		if samePitchClass(pitchClass, p) {
			return i, true
		}
	}

	return 0, false
}

func PitchClassBySemitone(semitone int) PitchClass {
	chromatic := ChromaticPitchClasses()
	semitone %= len(chromatic)
	if semitone < 0 {
		semitone += len(chromatic)
	}

	return chromatic[semitone]
}

func SemitoneDistance(from PitchClass, to PitchClass) int {
	fromSemitone, ok := from.Semitone()
	if !ok {
		return 0
	}

	toSemitone, ok := to.Semitone()
	if !ok {
		return 0
	}

	return (toSemitone - fromSemitone + 12) % 12
}

func TransposePitchClass(pitchClass PitchClass, semitones int) PitchClass {
	index, ok := pitchClass.Semitone()
	if !ok {
		return pitchClass
	}

	return PitchClassBySemitone(index + semitones)
}

func (n Note) String() string {
	return n.PitchClass.String() + fmt.Sprintf("%d", n.Octave)
}

func samePitchClass(a PitchClass, b PitchClass) bool {
	return a.SharpName == b.SharpName || a.FlatName == b.FlatName && b.FlatName != "" || a.SharpName == b.FlatName && b.FlatName != "" || a.FlatName == b.SharpName && a.FlatName != ""
}

func parseAlteredPitchClassName(name string) (PitchClass, bool) {
	if len(name) < 2 {
		return PitchClass{}, false
	}

	base, ok := naturalSemitones[name[0]]
	if !ok {
		return PitchClass{}, false
	}

	alteration := 0
	for _, accidental := range name[1:] {
		switch accidental {
		case '#':
			alteration++
		case 'b':
			alteration--
		default:
			return PitchClass{}, false
		}
	}

	if alteration < -2 || alteration > 2 {
		return PitchClass{}, false
	}

	return PitchClassBySemitone(base + alteration), true
}

var pitchClassAliases = map[string]PitchClass{
	"B#": C,
	"Fb": E,
	"E#": F,
	"Cb": B,
}

func PitchClassNames(pitchClasses []PitchClass) []string {
	names := make([]string, 0, len(pitchClasses))
	for _, pitchClass := range pitchClasses {
		names = append(names, pitchClass.Name())
	}

	return names
}

func PitchClassNamesString(pitchClasses []PitchClass) string {
	return strings.Join(PitchClassNames(pitchClasses), ", ")
}
