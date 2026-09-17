package chordshapes

type ChordQuality string

const (
	Major      ChordQuality = "maj"
	Minor      ChordQuality = "min"
	Dominant7  ChordQuality = "7"
	Major7     ChordQuality = "maj7"
	Minor7     ChordQuality = "m7"
	Diminished ChordQuality = "dim"
	Sus2       ChordQuality = "sus2"
	Sus4       ChordQuality = "sus4"
)

type ShapeTemplate struct {
	Name        string
	BaseRoot    Note
	Quality     ChordQuality
	Frets       [6]int
	Fingers     [6]int
	RootStrings []int
	Movable     bool
	Tags        []string
}

type Voicing struct {
	Name    string
	Root    Note
	Quality ChordQuality
	Frets   [6]int
	Fingers [6]int
	Source  string
	Tags    []string
}

var StandardTuning = [6]Note{E, A, D, G, B, E}
