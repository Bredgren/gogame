package sound

import (
	"math"
	"time"

	"github.com/gopherjs/gopherjs/js"
)

// Interface is the interface that a sound is expected to implement.
type Interface interface {
	// Play plays the sound from its current position. If the sound has ended, or you would
	// like to play it from the beginning the use SetPos(0) to go back to the start or use
	// PlayFromStart.
	Play()
	// PlayFromStart plays the sound from the beginning.
	PlayFromStart()
	// Pause stops playing the sound. It may be played again from where it left off by
	// calling Unpause.
	Pause()
	// Paused returns whether the sound is paused.
	Paused() bool
	// Ended returns whether the sound has ended.
	Ended() bool
	// Loop returns whether the sound will loop.
	Loop() bool
	// SetLoop sets whether the sound will loop.
	SetLoop(bool)
	// Volume returns the volume of the sound in the range 0.0 to 1.0.
	Volume() float64
	// SetVolume sets the volume of the sound. The value is clamped to the range 0.0 to 1.0.
	SetVolume(float64)
	// Length returns the total duration of the sound.
	Length() time.Duration
	// GetPos returns the current playback position relative to the start.
	Pos() time.Duration
	// SetPos sets the current playback position relative to the start.
	SetPos(time.Duration)
}

var _ Interface = &sound{}

type sound struct {
	snd *js.Object
}

// New creates a new sound.
func New(filename string) Interface {
	return &sound{
		snd: js.Global.Get("Audio").New(filename),
	}
}

func (s *sound) Play() {
	s.snd.Call("play")
}

func (s *sound) PlayFromStart() {
	s.snd.Set("currentTime", 0)
	s.snd.Call("play")
}

func (s *sound) Pause() {
	s.snd.Call("pause")
}

func (s *sound) Paused() bool {
	return s.snd.Get("paused").Bool()
}

func (s *sound) Ended() bool {
	return s.snd.Get("ended").Bool()
}

func (s *sound) Loop() bool {
	return s.snd.Get("loop").Bool()
}

func (s *sound) SetLoop(l bool) {
	s.snd.Set("loop", l)
}

func (s *sound) Volume() float64 {
	return s.snd.Get("volume").Float()
}

func (s *sound) SetVolume(v float64) {
	s.snd.Set("volume", math.Min(math.Max(0.0, v), 1.0))
}

func (s *sound) Length() time.Duration {
	return time.Duration(s.snd.Get("duration").Float()) * time.Second
}

func (s *sound) Pos() time.Duration {
	return time.Duration(s.snd.Get("currentTime").Float()) * time.Second
}

func (s *sound) SetPos(pos time.Duration) {
	s.snd.Set("currentTime", pos.Seconds())
}
