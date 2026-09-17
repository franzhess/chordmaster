package music

import (
	"strconv"
	"strings"
)

type RootNote string

type ScaleDegreeModifier string

const (
	ModifierNone       ScaleDegreeModifier = ""
	ModifierFlat       ScaleDegreeModifier = "b"
	ModifierSharp      ScaleDegreeModifier = "#"
	ModifierDoubleFlat ScaleDegreeModifier = "bb"
)

type ScaleDegree struct {
	Degree   int
	Modifier ScaleDegreeModifier
}

func (d ScaleDegree) String() string {
	return string(d.Modifier) + strconv.Itoa(d.Degree)
}

type ChordVariation struct {
	Name    string
	Suffix  string
	Formula []ScaleDegree
}

type Chord struct {
	Scale     Scale
	Variation ChordVariation
	NoteNames []string
}

type ScaleChordSchema struct {
	Name    string
	Formula []ScaleDegree
}

type ScaleChord struct {
	Scale      Scale
	RootDegree int
	Schema     ScaleChordSchema
	NoteNames  []string
}

func NewChord(scale Scale, variation ChordVariation) Chord {
	return Chord{
		Scale:     scale,
		Variation: variation,
		NoteNames: chordNoteNames(scale, variation),
	}
}

func (c Chord) NoteNamesString() string {
	return strings.Join(c.NoteNames, ", ")
}

func NewScaleChord(scale Scale, rootDegree int, schema ScaleChordSchema) ScaleChord {
	return ScaleChord{
		Scale:      scale,
		RootDegree: rootDegree,
		Schema:     schema,
		NoteNames:  scaleChordNoteNames(scale, rootDegree, schema),
	}
}

func (c ScaleChord) NoteNamesString() string {
	return strings.Join(c.NoteNames, ", ")
}

func (c ScaleChord) ChordName(catalog ChordCatalog) (string, bool) {
	return catalog.ChordNameByNotes(c.NoteNames)
}

func ScaleChords(scale Scale, schema ScaleChordSchema) []ScaleChord {
	chords := make([]ScaleChord, 0, len(scale.NoteNames))
	for rootDegree := 1; rootDegree <= len(scale.NoteNames); rootDegree++ {
		chords = append(chords, NewScaleChord(scale, rootDegree, schema))
	}

	return chords
}

func SupportedScaleChordSchemas() []ScaleChordSchema {
	return []ScaleChordSchema{
		{Name: "Triad", Formula: formula(degree(1), degree(3), degree(5))},
		{Name: "7th chord", Formula: formula(degree(1), degree(3), degree(5), degree(7))},
		{Name: "9th chord", Formula: formula(degree(1), degree(3), degree(5), degree(7), degree(9))},
		{Name: "11th chord", Formula: formula(degree(1), degree(3), degree(5), degree(7), degree(9), degree(11))},
		{Name: "13th chord", Formula: formula(degree(1), degree(3), degree(5), degree(7), degree(9), degree(11), degree(13))},
	}
}

func (v ChordVariation) FormulaString() string {
	return scaleDegreeFormulaString(v.Formula)
}

func (s ScaleChordSchema) FormulaString() string {
	return scaleDegreeFormulaString(s.Formula)
}

func scaleDegreeFormulaString(formula []ScaleDegree) string {
	degrees := make([]string, 0, len(formula))
	for _, degree := range formula {
		degrees = append(degrees, degree.String())
	}

	return strings.Join(degrees, ",")
}

type ChordCatalog struct {
	Roots      []RootNote
	Variations []ChordVariation
}

func DefaultChordCatalog() ChordCatalog {
	return ChordCatalog{
		Roots: []RootNote{"C", "C#", "D", "Eb", "E", "F", "F#", "G", "Ab", "A", "Bb", "B"},
		Variations: []ChordVariation{
			{Name: "Major", Suffix: "", Formula: formula(degree(1), degree(3), degree(5))},
			{Name: "Minor", Suffix: "m", Formula: formula(degree(1), flat(3), degree(5))},
			{Name: "Dominant7", Suffix: "7", Formula: formula(degree(1), degree(3), degree(5), flat(7))},
			{Name: "Minor7", Suffix: "m7", Formula: formula(degree(1), flat(3), degree(5), flat(7))},
			{Name: "Major7", Suffix: "maj7", Formula: formula(degree(1), degree(3), degree(5), degree(7))},
			{Name: "Sus4", Suffix: "sus4", Formula: formula(degree(1), degree(4), degree(5))},
			{Name: "Sus2", Suffix: "sus2", Formula: formula(degree(1), degree(2), degree(5))},
			{Name: "Power", Suffix: "5", Formula: formula(degree(1), degree(5))},
			{Name: "Major6", Suffix: "6", Formula: formula(degree(1), degree(3), degree(5), degree(6))},
			{Name: "Minor6", Suffix: "m6", Formula: formula(degree(1), flat(3), degree(5), degree(6))},
			{Name: "Diminished", Suffix: "dim", Formula: formula(degree(1), flat(3), flat(5))},
			{Name: "Augmented", Suffix: "aug", Formula: formula(degree(1), degree(3), sharp(5))},
			{Name: "Dominant7Sus4", Suffix: "7sus4", Formula: formula(degree(1), degree(4), degree(5), flat(7))},
			{Name: "HalfDiminished7", Suffix: "m7b5", Formula: formula(degree(1), flat(3), flat(5), flat(7))},
			{Name: "Diminished7", Suffix: "dim7", Formula: formula(degree(1), flat(3), flat(5), doubleFlat(7))},
			{Name: "Dominant9", Suffix: "9", Formula: formula(degree(1), degree(3), degree(5), flat(7), degree(9))},
			{Name: "Major9", Suffix: "maj9", Formula: formula(degree(1), degree(3), degree(5), degree(7), degree(9))},
			{Name: "Minor9", Suffix: "m9", Formula: formula(degree(1), flat(3), degree(5), flat(7), degree(9))},
			{Name: "Dominant13", Suffix: "13", Formula: formula(degree(1), degree(3), degree(5), flat(7), degree(9), degree(11), degree(13))},
			{Name: "Major13", Suffix: "maj13", Formula: formula(degree(1), degree(3), degree(5), degree(7), degree(9), degree(11), degree(13))},
			{Name: "Minor11", Suffix: "m11", Formula: formula(degree(1), flat(3), degree(5), flat(7), degree(9), degree(11))},
			{Name: "Dominant11", Suffix: "11", Formula: formula(degree(1), degree(3), degree(5), flat(7), degree(9), degree(11))},
			{Name: "Major11", Suffix: "maj11", Formula: formula(degree(1), degree(3), degree(5), degree(7), degree(9), degree(11))},
			{Name: "MinorMajor7", Suffix: "mmaj7", Formula: formula(degree(1), flat(3), degree(5), degree(7))},
			{Name: "Augmented7", Suffix: "aug7", Formula: formula(degree(1), degree(3), sharp(5), flat(7))},
			{Name: "PowerOctave", Suffix: "5o", Formula: formula(degree(1), degree(5), flat(8))},
		},
	}
}

func (c ChordCatalog) ChordName(rootIndex, variationIndex int) string {
	if rootIndex < 0 || rootIndex >= len(c.Roots) || variationIndex < 0 || variationIndex >= len(c.Variations) {
		return ""
	}

	return string(c.Roots[rootIndex]) + c.Variations[variationIndex].Suffix
}

func (c ChordCatalog) ChordNameByNotes(noteNames []string) (string, bool) {
	variation, ok := c.ChordVariationByNotes(noteNames)
	if !ok || len(noteNames) == 0 {
		return "", false
	}

	return noteNames[0] + variation.Suffix, true
}

func (c ChordCatalog) ChordVariationByNotes(noteNames []string) (ChordVariation, bool) {
	if len(noteNames) == 0 {
		return ChordVariation{}, false
	}

	offsets, ok := chordPitchClassOffsets(noteNames)
	if !ok {
		return ChordVariation{}, false
	}

	for _, variation := range c.Variations {
		if sameIntSet(offsets, formulaPitchClassOffsets(variation.Formula)) {
			return variation, true
		}
	}

	return ChordVariation{}, false
}

func formula(degrees ...ScaleDegree) []ScaleDegree {
	return degrees
}

func degree(value int) ScaleDegree {
	return ScaleDegree{Degree: value}
}

func flat(value int) ScaleDegree {
	return ScaleDegree{Degree: value, Modifier: ModifierFlat}
}

func sharp(value int) ScaleDegree {
	return ScaleDegree{Degree: value, Modifier: ModifierSharp}
}

func doubleFlat(value int) ScaleDegree {
	return ScaleDegree{Degree: value, Modifier: ModifierDoubleFlat}
}

func chordNoteNames(scale Scale, variation ChordVariation) []string {
	if len(scale.PitchClasses) < 7 || len(scale.NoteNames) < 7 {
		return nil
	}

	chromatic := ChromaticPitchClasses()
	names := make([]string, 0, len(variation.Formula))
	for _, scaleDegree := range variation.Formula {
		degreeIndex := (scaleDegree.Degree - 1) % 7
		basePitchClassIndex := pitchClassIndex(chromatic, scale.PitchClasses[degreeIndex])
		if basePitchClassIndex < 0 {
			continue
		}

		targetPitchClassIndex := (basePitchClassIndex + modifierSemitones(scaleDegree.Modifier) + len(chromatic)) % len(chromatic)
		letter := scale.NoteNames[degreeIndex][0]
		names = append(names, spellWithLetter(letter, targetPitchClassIndex))
	}

	return names
}

func scaleChordNoteNames(scale Scale, rootDegree int, schema ScaleChordSchema) []string {
	if len(scale.NoteNames) == 0 || rootDegree < 1 || rootDegree > len(scale.NoteNames) {
		return nil
	}

	names := make([]string, 0, len(schema.Formula))
	for _, scaleDegree := range schema.Formula {
		degreeIndex := (rootDegree - 1 + scaleDegree.Degree - 1) % len(scale.NoteNames)
		names = append(names, scale.NoteNames[degreeIndex])
	}

	return names
}

func modifierSemitones(modifier ScaleDegreeModifier) int {
	switch modifier {
	case ModifierFlat:
		return -1
	case ModifierSharp:
		return 1
	case ModifierDoubleFlat:
		return -2
	default:
		return 0
	}
}

func chordPitchClassOffsets(noteNames []string) ([]int, bool) {
	rootPitchClass, ok := noteNamePitchClass(noteNames[0])
	if !ok {
		return nil, false
	}

	offsets := make([]int, 0, len(noteNames))
	for _, noteName := range noteNames {
		pitchClass, ok := noteNamePitchClass(noteName)
		if !ok {
			return nil, false
		}

		offsets = append(offsets, (pitchClass-rootPitchClass+12)%12)
	}

	return offsets, true
}

func formulaPitchClassOffsets(formula []ScaleDegree) []int {
	offsets := make([]int, 0, len(formula))
	for _, scaleDegree := range formula {
		offsets = append(offsets, (majorScaleDegreeSemitones(scaleDegree.Degree)+modifierSemitones(scaleDegree.Modifier)+12)%12)
	}

	return offsets
}

func majorScaleDegreeSemitones(degree int) int {
	majorScale := []int{0, 2, 4, 5, 7, 9, 11}
	return majorScale[(degree-1)%len(majorScale)]
}

func noteNamePitchClass(noteName string) (int, bool) {
	if noteName == "" {
		return 0, false
	}

	letter := noteName[0]
	base, ok := naturalSemitones[letter]
	if !ok {
		return 0, false
	}

	modifier := 0
	for _, accidental := range noteName[1:] {
		switch accidental {
		case '#':
			modifier++
		case 'b':
			modifier--
		default:
			return 0, false
		}
	}

	return (base + modifier + 12) % 12, true
}

func sameIntSet(left, right []int) bool {
	if len(left) != len(right) {
		return false
	}

	seen := make(map[int]int, len(left))
	for _, value := range left {
		seen[value]++
	}

	for _, value := range right {
		if seen[value] == 0 {
			return false
		}
		seen[value]--
	}

	return true
}
