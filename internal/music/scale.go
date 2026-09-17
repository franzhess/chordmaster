package music

import "strings"

type ScaleStep struct {
	Name      string
	Semitones int
}

var (
	HalfStep        = ScaleStep{Name: "H", Semitones: 1}
	WholeStep       = ScaleStep{Name: "W", Semitones: 2}
	MinorThirdStep  = ScaleStep{Name: "m3", Semitones: 3}
	AugmentedSecond = ScaleStep{Name: "WH", Semitones: 3}
)

type ScalePattern struct {
	Name  string
	Steps []ScaleStep
}

type Scale struct {
	Root         PitchClass
	Pattern      ScalePattern
	PitchClasses []PitchClass
	NoteNames    []string
}

func NewScale(root PitchClass, pattern ScalePattern) Scale {
	chromatic := ChromaticPitchClasses()
	rootIndex := pitchClassIndex(chromatic, root)
	if rootIndex < 0 {
		return Scale{Root: root, Pattern: pattern}
	}

	currentIndex := rootIndex
	pitchClasses := []PitchClass{chromatic[currentIndex]}
	for _, step := range pattern.Steps[:len(pattern.Steps)-1] {
		currentIndex = (currentIndex + step.Semitones) % len(chromatic)
		pitchClasses = append(pitchClasses, chromatic[currentIndex])
	}

	return Scale{
		Root:         root,
		Pattern:      pattern,
		PitchClasses: pitchClasses,
		NoteNames:    spellScale(root, pitchClasses),
	}
}

func (s Scale) PitchClassNamesString() string {
	if len(s.NoteNames) > 0 {
		return strings.Join(s.NoteNames, ", ")
	}

	return PitchClassNamesString(s.PitchClasses)
}

func (p ScalePattern) StepString() string {
	steps := make([]string, 0, len(p.Steps))
	for _, step := range p.Steps {
		steps = append(steps, step.Name)
	}

	return strings.Join(steps, " ")
}

func (p ScalePattern) SemitoneCount() int {
	total := 0
	for _, step := range p.Steps {
		total += step.Semitones
	}

	return total
}

func SupportedScalePatterns() []ScalePattern {
	return []ScalePattern{
		{Name: "Major (Ionian)", Steps: steps(WholeStep, WholeStep, HalfStep, WholeStep, WholeStep, WholeStep, HalfStep)},
		{Name: "Natural Minor (Aeolian)", Steps: steps(WholeStep, HalfStep, WholeStep, WholeStep, HalfStep, WholeStep, WholeStep)},
		{Name: "Harmonic Minor", Steps: steps(WholeStep, HalfStep, WholeStep, WholeStep, HalfStep, AugmentedSecond, HalfStep)},
		{Name: "Melodic Minor (ascending)", Steps: steps(WholeStep, HalfStep, WholeStep, WholeStep, WholeStep, WholeStep, HalfStep)},
		{Name: "Dorian", Steps: steps(WholeStep, HalfStep, WholeStep, WholeStep, WholeStep, HalfStep, WholeStep)},
		{Name: "Phrygian", Steps: steps(HalfStep, WholeStep, WholeStep, WholeStep, HalfStep, WholeStep, WholeStep)},
		{Name: "Lydian", Steps: steps(WholeStep, WholeStep, WholeStep, HalfStep, WholeStep, WholeStep, HalfStep)},
		{Name: "Mixolydian", Steps: steps(WholeStep, WholeStep, HalfStep, WholeStep, WholeStep, HalfStep, WholeStep)},
		{Name: "Locrian", Steps: steps(HalfStep, WholeStep, WholeStep, HalfStep, WholeStep, WholeStep, WholeStep)},
		{Name: "Major Pentatonic", Steps: steps(WholeStep, WholeStep, MinorThirdStep, WholeStep, MinorThirdStep)},
		{Name: "Minor Pentatonic", Steps: steps(MinorThirdStep, WholeStep, WholeStep, MinorThirdStep, WholeStep)},
		{Name: "Blues Scale", Steps: steps(MinorThirdStep, WholeStep, HalfStep, HalfStep, MinorThirdStep, WholeStep)},
		{Name: "Chromatic", Steps: steps(HalfStep, HalfStep, HalfStep, HalfStep, HalfStep, HalfStep, HalfStep, HalfStep, HalfStep, HalfStep, HalfStep, HalfStep)},
		{Name: "Whole Tone", Steps: steps(WholeStep, WholeStep, WholeStep, WholeStep, WholeStep, WholeStep)},
		{Name: "Diminished (Half-Whole)", Steps: steps(HalfStep, WholeStep, HalfStep, WholeStep, HalfStep, WholeStep, HalfStep, WholeStep)},
		{Name: "Diminished (Whole-Half)", Steps: steps(WholeStep, HalfStep, WholeStep, HalfStep, WholeStep, HalfStep, WholeStep, HalfStep)},
	}
}

func steps(values ...ScaleStep) []ScaleStep {
	return values
}

func pitchClassIndex(pitchClasses []PitchClass, target PitchClass) int {
	for i, pitchClass := range pitchClasses {
		if pitchClass.SharpName == target.SharpName || pitchClass.FlatName == target.FlatName && target.FlatName != "" {
			return i
		}
	}

	return -1
}

func spellScale(root PitchClass, pitchClasses []PitchClass) []string {
	rootName := root.PreferredName(PreferSharps)
	if root.SharpName == "" && root.FlatName != "" {
		rootName = root.FlatName
	}

	preference := spellingPreference(rootName)
	if len(pitchClasses) != 7 {
		return preferredPitchClassNames(pitchClasses, preference)
	}

	rootLetter := rootName[0]
	rootLetterIndex := letterIndex(rootLetter)
	if rootLetterIndex < 0 {
		return preferredPitchClassNames(pitchClasses, preference)
	}

	names := make([]string, 0, len(pitchClasses))
	for i, pitchClass := range pitchClasses {
		letter := scaleLetters[(rootLetterIndex+i)%len(scaleLetters)]
		names = append(names, spellWithLetter(letter, pitchClassIndex(ChromaticPitchClasses(), pitchClass)))
	}

	return names
}

func preferredPitchClassNames(pitchClasses []PitchClass, preference SpellingPreference) []string {
	names := make([]string, 0, len(pitchClasses))
	for _, pitchClass := range pitchClasses {
		names = append(names, pitchClass.PreferredName(preference))
	}

	return names
}

var scaleLetters = []byte{'C', 'D', 'E', 'F', 'G', 'A', 'B'}

var naturalSemitones = map[byte]int{
	'C': 0,
	'D': 2,
	'E': 4,
	'F': 5,
	'G': 7,
	'A': 9,
	'B': 11,
}

func letterIndex(letter byte) int {
	for i, candidate := range scaleLetters {
		if candidate == letter {
			return i
		}
	}

	return -1
}

func spellWithLetter(letter byte, pitchClassIndex int) string {
	diff := (pitchClassIndex - naturalSemitones[letter] + 12) % 12
	switch diff {
	case 0:
		return string(letter)
	case 1:
		return string(letter) + "#"
	case 2:
		return string(letter) + "##"
	case 10:
		return string(letter) + "bb"
	case 11:
		return string(letter) + "b"
	default:
		return ChromaticPitchClasses()[pitchClassIndex].PreferredName(PreferSharps)
	}
}

func spellingPreference(rootName string) SpellingPreference {
	if strings.Contains(rootName, "b") {
		return PreferFlats
	}

	return PreferSharps
}
