package ui

import (
	"math/rand"
	"strings"
	"time"

	"chordmaster/internal/instruments/guitar"
	"chordmaster/internal/music"
	"chordmaster/internal/synth"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
)

type splashTimeoutMsg struct{}

type playbackSource int

const (
	playbackGeneral playbackSource = iota
	playbackSplash
)

type playbackFinishedMsg struct {
	err    error
	source playbackSource
}

type playbackKind int

const (
	playbackKindNone playbackKind = iota
	playbackKindScaleNotes
	playbackKindScaleChords
)

type playbackStepFinishedMsg struct {
	err     error
	kind    playbackKind
	session int
	index   int
}

type screen int

const (
	screenSplash screen = iota
	screenMainMenu
	screenChordBrowser
	screenScaleBrowser
	screenChordProgressions
	screenRiffs
)

type model struct {
	screen             screen
	menuItems          []string
	menuCursor         int
	chordCatalog       music.ChordCatalog
	pitchClasses       []music.PitchClass
	scalePatterns      []music.ScalePattern
	progressions       []music.ChordProgression
	rhythmPatterns     []music.RhythmPattern
	riffs              []music.Riff
	rootCursor         int
	variationCursor    int
	scaleRootCursor    int
	scalePatternCursor int
	progressionCursor  int
	rhythmCursor       int
	riffCursor         int
	splashRiff         music.Riff
	selectedInstrument string
	chordTabMode       chordTabMode
	isPlaying          bool
	playbackKind       playbackKind
	playbackSession    int
	playbackCancel     chan struct{}
	activeScaleNote    int
	activeScaleChord   int
	showHelp           bool
	keys               keyMap
	help               help.Model
	mainViewport       viewport.Model
	audio              *synth.Engine
	status             string
	width              int
	height             int
}

func NewModel(audio ...*synth.Engine) tea.Model {
	m := model{
		screen: screenSplash,
		menuItems: []string{
			"Scales",
			"Chords",
			"Progressions",
			"Riffs",
		},
		chordCatalog:       music.DefaultChordCatalog(),
		pitchClasses:       music.ChromaticPitchClasses(),
		scalePatterns:      music.SupportedScalePatterns(),
		progressions:       music.SupportedChordProgressions(),
		rhythmPatterns:     music.SupportedRhythmPatterns(),
		riffs:              guitar.SupportedRiffs(),
		selectedInstrument: "Guitar",
		activeScaleNote:    -1,
		activeScaleChord:   -1,
		keys:               defaultKeyMap(),
		help:               help.New(),
		mainViewport:       viewport.New(viewport.WithWidth(1), viewport.WithHeight(1)),
	}
	m.syncViewport()
	if len(m.riffs) > 0 {
		m.splashRiff = m.riffs[rand.Intn(len(m.riffs))]
	}
	if len(audio) > 0 {
		m.audio = audio[0]
		m.isPlaying = len(m.splashRiff.Events) > 0
	}

	return m
}

func (m model) Init() tea.Cmd {
	if m.audio != nil && len(m.splashRiff.Events) > 0 {
		return m.playRiff(m.splashRiff, playbackSplash)
	}

	return splashTimeout
}

func splashTimeout() tea.Msg {
	time.Sleep(3 * time.Second)
	return splashTimeoutMsg{}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return update(m, msg)
}

func pitchClassFromRoot(root string) music.PitchClass {
	if strings.Contains(root, "b") {
		return music.PitchClass{FlatName: root}
	}

	return music.PitchClass{SharpName: root}
}
