package synth

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"

	"chordmaster/internal/music"

	"github.com/hajimehoshi/oto/v2"
)

const (
	SampleRate          = 44100
	ChannelCount        = 1
	DefaultVolume       = 0.35
	ScaleNoteTime       = 400 * time.Millisecond
	ChordNoteTime       = 1200 * time.Millisecond
	ProgressionBeatTime = 300 * time.Millisecond
	fadeInTime          = 10 * time.Millisecond
	fadeOutTime         = 20 * time.Millisecond
	maxInt16Sample      = 32767
	pickNoiseTime       = 2 * time.Millisecond
)

const (
	strumDirectionDown = "down"
	strumDirectionUp   = "up"
)

var ErrPlaybackCanceled = errors.New("playback canceled")

type Engine struct {
	ctx *oto.Context
}

type TimedNote struct {
	Note     string
	Duration time.Duration
}

type StrumString struct {
	StringIndex      int
	StartSample      int
	Gain             float64
	Brightness       float64
	AttackSamples    int
	PreContactSample int
	PreContactGain   float64
}

type pluckedString struct {
	buffer []float64
	idx    int
	decay  float64
	damp   float64
	low    float64
}

type renderedString struct {
	stringModel pluckedString
	event       StrumString
}

func newPluckedString(freq float64, decay float64, damp float64, seed int) pluckedString {
	length := int(float64(SampleRate) / freq)
	if length < 2 {
		length = 2
	}

	buffer := make([]float64, length)
	for i := range buffer {
		buffer[i] = deterministicNoise(seed, i) * 0.72
	}

	return pluckedString{
		buffer: buffer,
		decay:  decay,
		damp:   damp,
	}
}

func (s *pluckedString) sample() float64 {
	out := s.buffer[s.idx]
	next := s.decay * 0.5 * (s.buffer[s.idx] + s.buffer[(s.idx+1)%len(s.buffer)])
	s.low += s.damp * (next - s.low)
	s.buffer[s.idx] = s.low
	s.idx = (s.idx + 1) % len(s.buffer)
	return out
}

func NewEngine() (*Engine, error) {
	ctx, ready, err := oto.NewContext(SampleRate, ChannelCount, oto.FormatSignedInt16LE)
	if err != nil {
		return nil, err
	}
	<-ready

	return &Engine{ctx: ctx}, nil
}

func (e *Engine) PlayScale(noteNames []string) error {
	for _, noteName := range noteNamesWithOctaves(noteNames, 4) {
		freq, err := NoteFrequency(noteName)
		if err != nil {
			return err
		}

		if err := e.PlayTone(freq, ScaleNoteTime); err != nil {
			return err
		}
	}

	return nil
}

func (e *Engine) PlayChord(noteNames []string) error {
	freqs := make([]float64, 0, len(noteNames))
	for _, noteName := range noteNamesWithOctaves(noteNames, 4) {
		freq, err := NoteFrequency(noteName)
		if err != nil {
			return err
		}
		freqs = append(freqs, freq)
	}

	return e.PlayChordFrequencies(freqs, ChordNoteTime)
}

func (e *Engine) PlayChordSequence(chords [][]string) error {
	for _, chord := range chords {
		if err := e.PlayChord(chord); err != nil {
			return err
		}
	}

	return nil
}

func (e *Engine) PlayTimedChordSequence(chords [][]string, durations []time.Duration) error {
	for i, chord := range chords {
		duration := ChordNoteTime
		if i < len(durations) {
			duration = durations[i]
		}

		freqs := make([]float64, 0, len(chord))
		for _, noteName := range noteNamesWithOctaves(chord, 4) {
			freq, err := NoteFrequency(noteName)
			if err != nil {
				return err
			}
			freqs = append(freqs, freq)
		}

		if err := e.PlayChordFrequencies(freqs, duration); err != nil {
			return err
		}
	}

	return nil
}

func (e *Engine) PlayToneCancellable(freq float64, duration time.Duration, cancel <-chan struct{}) error {
	player := e.ctx.NewPlayer(bytes.NewReader(generateSine(freq, duration, DefaultVolume)))
	return playToneCancellable(player, duration, cancel)
}

// PlayChordCancellable is used by the UI for step-by-step highlighted playback;
// closing cancel stops the current player before its nominal duration expires.
func (e *Engine) PlayChordCancellable(noteNames []string, duration time.Duration, cancel <-chan struct{}) error {
	freqs := make([]float64, 0, len(noteNames))
	for _, noteName := range noteNamesWithOctaves(noteNames, 4) {
		freq, err := NoteFrequency(noteName)
		if err != nil {
			return err
		}
		freqs = append(freqs, freq)
	}

	player := e.ctx.NewPlayer(bytes.NewReader(generateChord(freqs, duration, DefaultVolume)))
	return playToneCancellable(player, duration, cancel)
}

func (e *Engine) PlayTimedNotes(notes []TimedNote) error {
	for _, note := range notes {
		if note.Note == "R" {
			time.Sleep(note.Duration)
			continue
		}

		freq, err := NoteFrequency(note.Note)
		if err != nil {
			return err
		}

		if err := e.PlayTone(freq, note.Duration); err != nil {
			return err
		}
	}

	return nil
}

func (e *Engine) PlayTone(freq float64, duration time.Duration) error {
	player := e.ctx.NewPlayer(bytes.NewReader(generateSine(freq, duration, DefaultVolume)))
	return playTone(player, freq, duration)
}

func (e *Engine) PlayChordFrequencies(freqs []float64, duration time.Duration) error {
	player := e.ctx.NewPlayer(bytes.NewReader(generateChord(freqs, duration, DefaultVolume)))
	return playChord(player, freqs, duration)
}

func NoteFrequency(note string) (float64, error) {
	return noteFrequency(note)
}

func GenerateSine(freq float64, duration time.Duration, volume float64) []byte {
	return generateSine(freq, duration, volume)
}

func ScheduleStrum(chordStrings int, direction string, sampleRate int, durationMs int) []StrumString {
	return ScheduleStrumWithSource(chordStrings, direction, sampleRate, durationMs, rand.NewSource(1))
}

func ScheduleStrumWithSeed(chordStrings int, direction string, sampleRate int, durationMs int, seed int64) []StrumString {
	return ScheduleStrumWithSource(chordStrings, direction, sampleRate, durationMs, rand.NewSource(seed))
}

func ScheduleStrumWithSource(chordStrings int, direction string, sampleRate int, durationMs int, source rand.Source) []StrumString {
	if chordStrings <= 0 || sampleRate <= 0 || durationMs < 0 {
		return nil
	}

	if source == nil {
		source = rand.NewSource(1)
	}

	if direction != strumDirectionUp {
		direction = strumDirectionDown
	}

	rng := rand.New(source)
	events := make([]StrumString, chordStrings)
	durationSamples := int(float64(sampleRate) * float64(durationMs) / 1000)
	for order := 0; order < chordStrings; order++ {
		stringIndex := order
		if direction == strumDirectionUp {
			stringIndex = chordStrings - 1 - order
		}

		position := 0.0
		if chordStrings > 1 {
			position = float64(order) / float64(chordStrings-1)
		}

		// Spread string contact points across one continuous pick gesture, then add
		// small timing jitter so the brush does not feel machine-quantized.
		jitterMs := rng.Float64()*4 - 2
		start := int(position*float64(durationSamples) + jitterMs*float64(sampleRate)/1000)
		if start < 0 || chordStrings == 1 {
			start = 0
		}
		if start > durationSamples {
			start = durationSamples
		}

		// Velocity follows the pick stroke: early strings are slightly stronger,
		// then each string gets small deterministic velocity and brightness jitter.
		velocityCurve := 1.08 - 0.16*position
		velocityJitter := 0.85 + rng.Float64()*0.3
		brightnessJitter := 0.9 + rng.Float64()*0.18
		attackMs := 8 + rng.Float64()*17
		preContactMs := 2 + rng.Float64()*4
		preContactJitterMs := rng.Float64()*2 - 1
		preContactSample := start - int((preContactMs+preContactJitterMs)*float64(sampleRate)/1000)
		if preContactSample < 0 {
			preContactSample = 0
		}
		events[order] = StrumString{
			StringIndex:      stringIndex,
			StartSample:      start,
			Gain:             velocityCurve * velocityJitter,
			Brightness:       brightnessJitter,
			AttackSamples:    int(attackMs * float64(sampleRate) / 1000),
			PreContactSample: preContactSample,
			PreContactGain:   (0.005 + rng.Float64()*0.007) * velocityJitter,
		}
	}

	return events
}

func SlowDownStrum(chordStrings int) []StrumString {
	return ScheduleStrumWithSource(chordStrings, strumDirectionDown, SampleRate, presetDurationMs(150, 220, rand.NewSource(11)), rand.NewSource(12))
}

func NormalDownStrum(chordStrings int) []StrumString {
	return ScheduleStrumWithSource(chordStrings, strumDirectionDown, SampleRate, presetDurationMs(80, 130, rand.NewSource(21)), rand.NewSource(22))
}

func FastUpStrum(chordStrings int) []StrumString {
	return ScheduleStrumWithSource(chordStrings, strumDirectionUp, SampleRate, presetDurationMs(45, 70, rand.NewSource(31)), rand.NewSource(32))
}

// RenderAlternatingStrumExample renders a simple C-G-Am-F guitar progression
// with alternating down/up strums. It is intentionally small so callers can use
// it as a reference for scheduling custom strums without going through Engine.
func RenderAlternatingStrumExample() ([]byte, error) {
	progression := []struct {
		notes       []string
		direction   string
		durationMs  int
		strumSeed   int64
		sustainTime time.Duration
	}{
		{notes: []string{"C3", "E3", "G3", "C4", "E4"}, direction: strumDirectionDown, durationMs: 95, strumSeed: 101, sustainTime: ChordNoteTime},
		{notes: []string{"G2", "B2", "D3", "G3", "B3", "G4"}, direction: strumDirectionUp, durationMs: 55, strumSeed: 102, sustainTime: ChordNoteTime},
		{notes: []string{"A2", "E3", "A3", "C4", "E4"}, direction: strumDirectionDown, durationMs: 105, strumSeed: 103, sustainTime: ChordNoteTime},
		{notes: []string{"F2", "C3", "F3", "A3", "C4", "F4"}, direction: strumDirectionUp, durationMs: 60, strumSeed: 104, sustainTime: ChordNoteTime},
	}

	var rendered []byte
	for _, chord := range progression {
		freqs := make([]float64, 0, len(chord.notes))
		for _, note := range chord.notes {
			freq, err := NoteFrequency(note)
			if err != nil {
				return nil, err
			}
			freqs = append(freqs, freq)
		}

		strum := ScheduleStrumWithSeed(len(freqs), chord.direction, SampleRate, chord.durationMs, chord.strumSeed)
		rendered = append(rendered, generatePluckedChordWithStrum(freqs, chord.sustainTime, DefaultVolume, strum)...)
	}

	return rendered, nil
}

func presetDurationMs(minMs int, maxMs int, source rand.Source) int {
	if maxMs <= minMs {
		return minMs
	}

	return minMs + rand.New(source).Intn(maxMs-minMs+1)
}

func noteFrequency(note string) (float64, error) {
	name, octave, err := parseNote(note)
	if err != nil {
		return 0, err
	}

	pitchClass, err := music.ParsePitchClassName(name)
	if err != nil {
		return 0, err
	}

	semitone, ok := pitchClass.Semitone()
	if !ok {
		return 0, fmt.Errorf("unknown note %q", note)
	}

	midi := (octave+1)*12 + semitone
	return 440 * math.Pow(2, float64(midi-69)/12), nil
}

func generateSine(freq float64, duration time.Duration, volume float64) []byte {
	return generatePluckedNote(freq, duration, volume)
}

func generateChord(freqs []float64, duration time.Duration, volume float64) []byte {
	return generatePluckedChord(freqs, duration, volume)
}

func playTone(player oto.Player, _ float64, duration time.Duration) error {
	player.Play()
	time.Sleep(duration)
	err := player.Err()
	closeErr := player.Close()
	if err != nil {
		return err
	}
	return closeErr
}

func playToneCancellable(player oto.Player, duration time.Duration, cancel <-chan struct{}) error {
	player.Play()
	select {
	case <-time.After(duration):
		err := player.Err()
		closeErr := player.Close()
		if err != nil {
			return err
		}
		return closeErr
	case <-cancel:
		_ = player.Close()
		return ErrPlaybackCanceled
	}
}

func playChord(player oto.Player, _ []float64, duration time.Duration) error {
	player.Play()
	time.Sleep(duration)
	err := player.Err()
	closeErr := player.Close()
	if err != nil {
		return err
	}
	return closeErr
}

func generatePCM(freqs []float64, duration time.Duration, volume float64) []byte {
	if len(freqs) == 0 || duration <= 0 || volume <= 0 {
		return nil
	}

	if volume > 1 {
		volume = 1
	}

	sampleCount := int(float64(SampleRate) * duration.Seconds())
	data := make([]byte, sampleCount*2)
	for sample := 0; sample < sampleCount; sample++ {
		t := float64(sample) / SampleRate
		mixed := 0.0
		for _, freq := range freqs {
			mixed += math.Sin(2 * math.Pi * freq * t)
		}
		mixed /= float64(len(freqs))
		mixed *= volume * envelope(sample, sampleCount)

		value := int16(math.Round(mixed * maxInt16Sample))
		binary.LittleEndian.PutUint16(data[sample*2:], uint16(value))
	}

	return data
}

func generatePluckedNote(freq float64, duration time.Duration, volume float64) []byte {
	return generatePluckedChord([]float64{freq}, duration, volume)
}

func generatePluckedChord(freqs []float64, duration time.Duration, volume float64) []byte {
	return generatePluckedChordWithStrum(freqs, duration, volume, NormalDownStrum(len(freqs)))
}

func generatePluckedChordWithStrum(freqs []float64, duration time.Duration, volume float64, strum []StrumString) []byte {
	if len(freqs) == 0 || duration <= 0 || volume <= 0 {
		return nil
	}

	if volume > 1 {
		volume = 1
	}

	if len(strum) == 0 {
		strum = ScheduleStrum(len(freqs), strumDirectionDown, SampleRate, 0)
	}

	strings := make([]renderedString, 0, len(freqs))
	for _, event := range strum {
		if event.StringIndex < 0 || event.StringIndex >= len(freqs) {
			continue
		}

		freq := freqs[event.StringIndex]
		if freq <= 0 {
			continue
		}

		detune := 1 + (float64((event.StringIndex%3)-1) * 0.0015)
		stringFreq := freq * detune
		decay := 0.9987 - math.Min(0.0014, stringFreq/300000)
		damp := (0.42 + math.Min(0.18, stringFreq/3000)) * event.Brightness
		if damp > 0.72 {
			damp = 0.72
		}
		if damp < 0.32 {
			damp = 0.32
		}
		if event.Gain <= 0 {
			event.Gain = 1
		}
		if event.AttackSamples <= 0 {
			sampleRate := float64(SampleRate)
			event.AttackSamples = int(0.014 * sampleRate)
		}
		strings = append(strings, renderedString{
			stringModel: newPluckedString(stringFreq, decay, damp, event.StringIndex+1),
			event:       event,
		})
	}

	if len(strings) == 0 {
		return nil
	}

	sampleCount := int(float64(SampleRate) * duration.Seconds())
	data := make([]byte, sampleCount*2)
	brushStart, brushEnd := strumWindow(strings)
	body := 0.0
	mixLow := 0.0
	scrapePrevIn := 0.0
	scrapeHigh := 0.0
	scrapeBand := 0.0
	for sample := 0; sample < sampleCount; sample++ {
		mixed := 0.0
		active := 0
		for i := range strings {
			mixed += mutedPreContactNoise(strings[i].event, sample)

			localSample := sample - strings[i].event.StartSample
			if localSample < 0 {
				continue
			}

			active++
			attack := curvedAttack(localSample, strings[i].event.AttackSamples)
			stringSample := strings[i].stringModel.sample() * strings[i].event.Gain * attack
			mixed += stringSample
		}

		mixed += brushScrapeNoise(sample, brushStart, brushEnd, &scrapePrevIn, &scrapeHigh, &scrapeBand)

		if active > 0 {
			mixed /= math.Sqrt(float64(active))
		}

		body += 0.018 * (mixed - body)
		mixed += body * 0.28
		mixLow += 0.58 * (mixed - mixLow)
		mixed = mixLow
		mixed *= volume * envelope(sample, sampleCount)
		mixed = math.Tanh(mixed * 1.15)

		value := int16(math.Round(mixed * maxInt16Sample))
		binary.LittleEndian.PutUint16(data[sample*2:], uint16(value))
	}

	return data
}

func strumWindow(strings []renderedString) (int, int) {
	first := strings[0].event.StartSample
	last := strings[0].event.StartSample
	for _, s := range strings[1:] {
		if s.event.StartSample < first {
			first = s.event.StartSample
		}
		if s.event.StartSample > last {
			last = s.event.StartSample
		}
	}

	sampleRate := float64(SampleRate)
	pre := int(0.012 * sampleRate)
	post := int(0.028 * sampleRate)
	first -= pre
	if first < 0 {
		first = 0
	}

	return first, last + post
}

func curvedAttack(localSample int, attackSamples int) float64 {
	if attackSamples <= 0 || localSample >= attackSamples {
		return 1
	}

	x := float64(localSample) / float64(attackSamples)
	return x * x * (3 - 2*x)
}

func brushScrapeNoise(sample int, start int, end int, prevIn *float64, high *float64, band *float64) float64 {
	if sample < start || sample > end || end <= start {
		return 0
	}

	position := float64(sample-start) / float64(end-start)
	shape := math.Sin(math.Pi * position)
	if shape < 0 {
		shape = 0
	}

	input := deterministicNoise(211, sample)
	*high = 0.94 * (*high + input - *prevIn)
	*prevIn = input
	*band += 0.22 * (*high - *band)
	return (*high - *band*0.35) * shape * 0.005
}

func mutedPreContactNoise(event StrumString, sample int) float64 {
	if sample < event.PreContactSample || sample >= event.StartSample || event.StartSample <= event.PreContactSample {
		return 0
	}

	position := float64(sample-event.PreContactSample) / float64(event.StartSample-event.PreContactSample)
	shape := math.Sin(math.Pi * position)
	noise := deterministicNoise(event.StringIndex+97, sample)
	filtered := noise - deterministicNoise(event.StringIndex+97, sample-1)*0.82
	return filtered * shape * event.PreContactGain
}

func deterministicNoise(seed int, index int) float64 {
	x := uint32(seed*374761393 + index*668265263)
	x = (x ^ (x >> 13)) * 1274126177
	x ^= x >> 16
	return (float64(x)/float64(^uint32(0)))*2 - 1
}

func envelope(sample int, sampleCount int) float64 {
	fadeInSamples := int(float64(SampleRate) * fadeInTime.Seconds())
	fadeOutSamples := int(float64(SampleRate) * fadeOutTime.Seconds())

	if fadeInSamples > 0 && sample < fadeInSamples {
		return float64(sample) / float64(fadeInSamples)
	}

	if fadeOutSamples > 0 && sample >= sampleCount-fadeOutSamples {
		return float64(sampleCount-sample-1) / float64(fadeOutSamples)
	}

	return 1
}

func noteNamesWithOctaves(noteNames []string, octave int) []string {
	notes := make([]string, 0, len(noteNames))
	previous := -1
	currentOctave := octave
	for _, noteName := range noteNames {
		name := strings.TrimSpace(noteName)
		pitchClass, err := music.ParsePitchClassName(name)
		semitone := 0
		ok := err == nil
		if ok {
			semitone, ok = pitchClass.Semitone()
		}
		if !ok {
			notes = append(notes, name+strconv.Itoa(currentOctave))
			continue
		}

		if previous >= 0 && semitone < previous {
			currentOctave++
		}
		previous = semitone

		notes = append(notes, name+strconv.Itoa(currentOctave))
	}

	return notes
}

func NoteNamesWithOctaves(noteNames []string, octave int) []string {
	return noteNamesWithOctaves(noteNames, octave)
}

func parseNote(note string) (string, int, error) {
	note = strings.TrimSpace(note)
	if note == "" {
		return "", 0, errors.New("empty note")
	}

	split := len(note)
	for split > 0 && unicode.IsDigit(rune(note[split-1])) {
		split--
	}

	if split == len(note) {
		return "", 0, fmt.Errorf("missing octave in %q", note)
	}

	octave, err := strconv.Atoi(note[split:])
	if err != nil {
		return "", 0, err
	}

	return note[:split], octave, nil
}
