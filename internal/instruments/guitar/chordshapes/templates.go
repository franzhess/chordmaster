package chordshapes

var shapeTemplates = []ShapeTemplate{
	// Open non-movable shapes go first since we want to prioritize them
	{Name: "C open", BaseRoot: C, Quality: Major, Frets: [6]int{-1, 3, 2, 0, 1, 0}, Fingers: [6]int{0, 3, 2, 0, 1, 0}, RootStrings: []int{1}, Movable: false, Tags: []string{"open", "beginner"}},
	{Name: "D open", BaseRoot: D, Quality: Major, Frets: [6]int{-1, -1, 0, 2, 3, 2}, Fingers: [6]int{0, 0, 0, 1, 3, 2}, RootStrings: []int{3}, Movable: false, Tags: []string{"open", "beginner"}},
	{Name: "E open", BaseRoot: E, Quality: Major, Frets: [6]int{0, 2, 2, 1, 0, 0}, Fingers: [6]int{0, 2, 3, 1, 0, 0}, RootStrings: []int{0, 5}, Movable: false, Tags: []string{"open", "beginner"}},
	{Name: "F open", BaseRoot: F, Quality: Major, Frets: [6]int{-1, -1, 3, 2, 1, 1}, Fingers: [6]int{0, 0, 3, 2, 1, 1}, RootStrings: []int{3, 5}, Movable: false, Tags: []string{"open"}},
	{Name: "G open", BaseRoot: G, Quality: Major, Frets: [6]int{3, 2, 0, 0, 0, 3}, Fingers: [6]int{2, 1, 0, 0, 0, 3}, RootStrings: []int{0, 5}, Movable: false, Tags: []string{"open", "beginner"}},
	{Name: "A open", BaseRoot: A, Quality: Major, Frets: [6]int{-1, 0, 2, 2, 2, 0}, Fingers: [6]int{0, 0, 1, 2, 3, 0}, RootStrings: []int{1}, Movable: false, Tags: []string{"open", "beginner"}},

	{Name: "Am open", BaseRoot: A, Quality: Minor, Frets: [6]int{-1, 0, 2, 2, 1, 0}, Fingers: [6]int{0, 0, 2, 3, 1, 0}, RootStrings: []int{1}, Movable: false, Tags: []string{"open", "beginner"}},
	{Name: "Dm open", BaseRoot: D, Quality: Minor, Frets: [6]int{-1, -1, 0, 2, 3, 1}, Fingers: [6]int{0, 0, 0, 2, 3, 1}, RootStrings: []int{3}, Movable: false, Tags: []string{"open", "beginner"}},
	{Name: "Em open", BaseRoot: E, Quality: Minor, Frets: [6]int{0, 2, 2, 0, 0, 0}, Fingers: [6]int{0, 2, 3, 0, 0, 0}, RootStrings: []int{0, 5}, Movable: false, Tags: []string{"open", "beginner"}},

	{Name: "C7 open", BaseRoot: C, Quality: Dominant7, Frets: [6]int{-1, 3, 2, 3, 1, 0}, Fingers: [6]int{0, 3, 2, 4, 1, 0}, RootStrings: []int{1}, Movable: false, Tags: []string{"open"}},
	{Name: "D7 open", BaseRoot: D, Quality: Dominant7, Frets: [6]int{-1, -1, 0, 2, 1, 2}, Fingers: [6]int{0, 0, 0, 2, 1, 3}, RootStrings: []int{3}, Movable: false, Tags: []string{"open"}},
	{Name: "E7 open", BaseRoot: E, Quality: Dominant7, Frets: [6]int{0, 2, 0, 1, 0, 0}, Fingers: [6]int{0, 2, 0, 1, 0, 0}, RootStrings: []int{0, 5}, Movable: false, Tags: []string{"open"}},
	{Name: "G7 open", BaseRoot: G, Quality: Dominant7, Frets: [6]int{3, 2, 0, 0, 0, 1}, Fingers: [6]int{3, 2, 0, 0, 0, 1}, RootStrings: []int{0, 5}, Movable: false, Tags: []string{"open"}},
	{Name: "A7 open", BaseRoot: A, Quality: Dominant7, Frets: [6]int{-1, 0, 2, 0, 2, 0}, Fingers: [6]int{0, 0, 2, 0, 3, 0}, RootStrings: []int{1}, Movable: false, Tags: []string{"open"}},
	{Name: "B7 open", BaseRoot: B, Quality: Dominant7, Frets: [6]int{-1, 2, 1, 2, 0, 2}, Fingers: [6]int{0, 2, 1, 3, 0, 4}, RootStrings: []int{1, 4}, Movable: false, Tags: []string{"open"}},

	{Name: "Cmaj7 open", BaseRoot: C, Quality: Major7, Frets: [6]int{-1, 3, 2, 0, 0, 0}, Fingers: [6]int{0, 3, 2, 0, 0, 0}, RootStrings: []int{1}, Movable: false, Tags: []string{"open"}},
	{Name: "Dmaj7 open", BaseRoot: D, Quality: Major7, Frets: [6]int{-1, -1, 0, 2, 2, 2}, Fingers: [6]int{0, 0, 0, 1, 1, 1}, RootStrings: []int{3}, Movable: false, Tags: []string{"open"}},
	{Name: "Emaj7 open", BaseRoot: E, Quality: Major7, Frets: [6]int{0, 2, 1, 1, 0, 0}, Fingers: [6]int{0, 3, 1, 2, 0, 0}, RootStrings: []int{0, 5}, Movable: false, Tags: []string{"open"}},
	{Name: "Fmaj7 open", BaseRoot: F, Quality: Major7, Frets: [6]int{-1, -1, 3, 2, 1, 0}, Fingers: [6]int{0, 0, 3, 2, 1, 0}, RootStrings: []int{3}, Movable: false, Tags: []string{"open"}},
	{Name: "Gmaj7 open", BaseRoot: G, Quality: Major7, Frets: [6]int{3, 2, 0, 0, 0, 2}, Fingers: [6]int{3, 2, 0, 0, 0, 1}, RootStrings: []int{0, 3}, Movable: false, Tags: []string{"open"}},
	{Name: "Amaj7 open", BaseRoot: A, Quality: Major7, Frets: [6]int{-1, 0, 2, 1, 2, 0}, Fingers: [6]int{0, 0, 2, 1, 3, 0}, RootStrings: []int{1}, Movable: false, Tags: []string{"open"}},

	{Name: "Am7 open", BaseRoot: A, Quality: Minor7, Frets: [6]int{-1, 0, 2, 0, 1, 0}, Fingers: [6]int{0, 0, 2, 0, 1, 0}, RootStrings: []int{1}, Movable: false, Tags: []string{"open"}},
	{Name: "Dm7 open", BaseRoot: D, Quality: Minor7, Frets: [6]int{-1, -1, 0, 2, 1, 1}, Fingers: [6]int{0, 0, 0, 2, 1, 1}, RootStrings: []int{3}, Movable: false, Tags: []string{"open"}},
	{Name: "Em7 open", BaseRoot: E, Quality: Minor7, Frets: [6]int{0, 2, 0, 0, 0, 0}, Fingers: [6]int{0, 2, 0, 0, 0, 0}, RootStrings: []int{0, 5}, Movable: false, Tags: []string{"open"}},

	{Name: "Bdim open", BaseRoot: B, Quality: Diminished, Frets: [6]int{-1, 2, 3, 4, 3, -1}, Fingers: [6]int{0, 1, 2, 4, 3, 0}, RootStrings: []int{1, 3}, Movable: false, Tags: []string{"open"}},

	{Name: "Csus2 open", BaseRoot: C, Quality: Sus2, Frets: [6]int{-1, 3, 0, 0, 1, 0}, Fingers: [6]int{0, 3, 0, 0, 1, 0}, RootStrings: []int{1}, Movable: false, Tags: []string{"open"}},
	{Name: "Dsus2 open", BaseRoot: D, Quality: Sus2, Frets: [6]int{-1, -1, 0, 2, 3, 0}, Fingers: [6]int{0, 0, 0, 1, 3, 0}, RootStrings: []int{3}, Movable: false, Tags: []string{"open"}},
	{Name: "Asus2 open", BaseRoot: A, Quality: Sus2, Frets: [6]int{-1, 0, 2, 2, 0, 0}, Fingers: [6]int{0, 0, 1, 2, 0, 0}, RootStrings: []int{1}, Movable: false, Tags: []string{"open"}},

	{Name: "Csus4 open", BaseRoot: C, Quality: Sus4, Frets: [6]int{-1, 3, 3, 0, 1, 1}, Fingers: [6]int{0, 3, 4, 0, 1, 1}, RootStrings: []int{1}, Movable: false, Tags: []string{"open"}},
	{Name: "Dsus4 open", BaseRoot: D, Quality: Sus4, Frets: [6]int{-1, -1, 0, 2, 3, 3}, Fingers: [6]int{0, 0, 0, 1, 3, 4}, RootStrings: []int{3}, Movable: false, Tags: []string{"open"}},
	{Name: "Esus4 open", BaseRoot: E, Quality: Sus4, Frets: [6]int{0, 2, 2, 2, 0, 0}, Fingers: [6]int{0, 1, 2, 3, 0, 0}, RootStrings: []int{0, 5}, Movable: false, Tags: []string{"open"}},
	{Name: "Asus4 open", BaseRoot: A, Quality: Sus4, Frets: [6]int{-1, 0, 2, 2, 3, 0}, Fingers: [6]int{0, 0, 1, 2, 3, 0}, RootStrings: []int{1}, Movable: false, Tags: []string{"open"}},

	// Movable CAGED/barre templates.
	{Name: "E-shape major barre", BaseRoot: F, Quality: Major, Frets: [6]int{1, 3, 3, 2, 1, 1}, Fingers: [6]int{1, 3, 4, 2, 1, 1}, RootStrings: []int{0, 5}, Movable: true, Tags: []string{"barre", "caged", "e-shape"}},
	{Name: "A-shape major barre", BaseRoot: As, Quality: Major, Frets: [6]int{-1, 1, 3, 3, 3, 1}, Fingers: [6]int{0, 1, 3, 3, 3, 1}, RootStrings: []int{1}, Movable: true, Tags: []string{"barre", "caged", "a-shape"}},
	{Name: "C-shape major movable", BaseRoot: Cs, Quality: Major, Frets: [6]int{-1, 4, 3, 1, 2, 1}, Fingers: [6]int{0, 4, 3, 1, 2, 1}, RootStrings: []int{1}, Movable: true, Tags: []string{"caged", "c-shape"}},
	{Name: "G-shape major movable", BaseRoot: Gs, Quality: Major, Frets: [6]int{4, 3, 1, 1, 1, 4}, Fingers: [6]int{4, 3, 1, 1, 1, 4}, RootStrings: []int{0, 5}, Movable: true, Tags: []string{"caged", "g-shape"}},
	{Name: "D-shape major movable", BaseRoot: Ds, Quality: Major, Frets: [6]int{-1, -1, 1, 3, 4, 3}, Fingers: [6]int{0, 0, 1, 3, 4, 2}, RootStrings: []int{3}, Movable: true, Tags: []string{"caged", "d-shape"}},

	{Name: "Em-shape minor barre", BaseRoot: F, Quality: Minor, Frets: [6]int{1, 3, 3, 1, 1, 1}, Fingers: [6]int{1, 3, 4, 1, 1, 1}, RootStrings: []int{0, 5}, Movable: true, Tags: []string{"barre", "caged", "em-shape"}},
	{Name: "Am-shape minor barre", BaseRoot: As, Quality: Minor, Frets: [6]int{-1, 1, 3, 3, 2, 1}, Fingers: [6]int{0, 1, 3, 4, 2, 1}, RootStrings: []int{1}, Movable: true, Tags: []string{"barre", "caged", "am-shape"}},
	{Name: "Dm-shape minor movable", BaseRoot: Ds, Quality: Minor, Frets: [6]int{-1, -1, 1, 3, 4, 2}, Fingers: [6]int{0, 0, 1, 3, 4, 2}, RootStrings: []int{3}, Movable: true, Tags: []string{"caged", "dm-shape"}},

	{Name: "E7-shape barre", BaseRoot: F, Quality: Dominant7, Frets: [6]int{1, 3, 1, 2, 1, 1}, Fingers: [6]int{1, 3, 1, 2, 1, 1}, RootStrings: []int{0, 5}, Movable: true, Tags: []string{"barre", "caged", "e7-shape"}},
	{Name: "A7-shape barre", BaseRoot: As, Quality: Dominant7, Frets: [6]int{-1, 1, 3, 1, 3, 1}, Fingers: [6]int{0, 1, 3, 1, 4, 1}, RootStrings: []int{1}, Movable: true, Tags: []string{"barre", "caged", "a7-shape"}},
	{Name: "C7-shape movable", BaseRoot: Cs, Quality: Dominant7, Frets: [6]int{-1, 4, 3, 4, 2, 1}, Fingers: [6]int{0, 4, 3, 4, 2, 1}, RootStrings: []int{1}, Movable: true, Tags: []string{"caged", "c7-shape"}},

	{Name: "Emaj7-shape barre", BaseRoot: F, Quality: Major7, Frets: [6]int{1, 3, 2, 2, 1, 1}, Fingers: [6]int{1, 4, 2, 3, 1, 1}, RootStrings: []int{0, 5}, Movable: true, Tags: []string{"barre", "caged", "emaj7-shape"}},
	{Name: "Amaj7-shape barre", BaseRoot: As, Quality: Major7, Frets: [6]int{-1, 1, 3, 2, 3, 1}, Fingers: [6]int{0, 1, 3, 2, 4, 1}, RootStrings: []int{1}, Movable: true, Tags: []string{"barre", "caged", "amaj7-shape"}},
	{Name: "Cmaj7-shape movable", BaseRoot: Cs, Quality: Major7, Frets: [6]int{-1, 4, 3, 1, 1, 1}, Fingers: [6]int{0, 4, 3, 1, 1, 1}, RootStrings: []int{1}, Movable: true, Tags: []string{"caged", "cmaj7-shape"}},

	{Name: "Em7-shape barre", BaseRoot: F, Quality: Minor7, Frets: [6]int{1, 3, 1, 1, 1, 1}, Fingers: [6]int{1, 3, 1, 1, 1, 1}, RootStrings: []int{0, 5}, Movable: true, Tags: []string{"barre", "caged", "em7-shape"}},
	{Name: "Am7-shape barre", BaseRoot: As, Quality: Minor7, Frets: [6]int{-1, 1, 3, 1, 2, 1}, Fingers: [6]int{0, 1, 3, 1, 2, 1}, RootStrings: []int{1}, Movable: true, Tags: []string{"barre", "caged", "am7-shape"}},

	{Name: "Asus2-shape movable", BaseRoot: As, Quality: Sus2, Frets: [6]int{-1, 1, 3, 3, 1, 1}, Fingers: [6]int{0, 1, 3, 4, 1, 1}, RootStrings: []int{1}, Movable: true, Tags: []string{"caged", "asus2-shape"}},
	{Name: "Dsus2-shape movable", BaseRoot: Ds, Quality: Sus2, Frets: [6]int{-1, -1, 1, 3, 4, 1}, Fingers: [6]int{0, 0, 1, 3, 4, 1}, RootStrings: []int{3}, Movable: true, Tags: []string{"caged", "dsus2-shape"}},

	{Name: "Esus4-shape barre", BaseRoot: F, Quality: Sus4, Frets: [6]int{1, 3, 3, 3, 1, 1}, Fingers: [6]int{1, 2, 3, 4, 1, 1}, RootStrings: []int{0, 5}, Movable: true, Tags: []string{"barre", "caged", "esus4-shape"}},
	{Name: "Asus4-shape barre", BaseRoot: As, Quality: Sus4, Frets: [6]int{-1, 1, 3, 3, 4, 1}, Fingers: [6]int{0, 1, 2, 3, 4, 1}, RootStrings: []int{1}, Movable: true, Tags: []string{"barre", "caged", "asus4-shape"}},
	{Name: "Dsus4-shape movable", BaseRoot: Ds, Quality: Sus4, Frets: [6]int{-1, -1, 1, 3, 4, 4}, Fingers: [6]int{0, 0, 1, 3, 4, 4}, RootStrings: []int{3}, Movable: true, Tags: []string{"caged", "dsus4-shape"}},
}

func Templates() []ShapeTemplate {
	templates := make([]ShapeTemplate, len(shapeTemplates))
	copy(templates, shapeTemplates)
	return templates
}
