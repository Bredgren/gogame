package ggweb

import (
	"fmt"
	"math"
	"time"

	"github.com/gopherjs/gopherjs/js"
)

// Sound is a music or sound clip.
type Sound struct {
	snd *js.Object
}

// NewSound creates a new sound from the given sound file.
func NewSound(filename string) *Sound {
	// Append timestamp to end of file to prevent caching.
	return &Sound{
		snd: js.Global.Get("Audio").New(fmt.Sprintf("%s?cb=%d", filename, time.Now().Unix())),
	}
}

// func (s *Sound) PlayMode() {
// 	// TODO
// }

// Play plays the sound from its current position. If the sound has ended, or you would
// like to play it from the beginning the use SetPos(0) to go back to the start or use
// PlayFromStart.
func (s *Sound) Play() {
	s.snd.Call("play")
}

// PlayFromStart plays the sound from the beginning.
func (s *Sound) PlayFromStart() {
	s.snd.Set("currentTime", 0)
	s.snd.Call("play")
}

// Pause stops playing the sound. It may be played again from where it left off by
// calling Play.
func (s *Sound) Pause() {
	s.snd.Call("pause")
}

// Paused returns whether the sound is paused.
func (s *Sound) Paused() bool {
	return s.snd.Get("paused").Bool()
}

// Ended returns whether the sound has ended.
func (s *Sound) Ended() bool {
	return s.snd.Get("ended").Bool()
}

// Loop returns whether the sound will loop.
func (s *Sound) Loop() bool {
	return s.snd.Get("loop").Bool()
}

// SetLoop sets whether the sound will loop.
func (s *Sound) SetLoop(l bool) {
	s.snd.Set("loop", l)
}

// Volume returns the volume of the sound in the range 0.0 to 1.0.
func (s *Sound) Volume() float64 {
	return s.snd.Get("volume").Float()
}

// SetVolume sets the volume of the sound. The value is clamped to the range 0.0 to 1.0.
func (s *Sound) SetVolume(v float64) {
	s.snd.Set("volume", math.Min(math.Max(0.0, v), 1.0))
}

// Length returns the total duration of the sound.
func (s *Sound) Length() time.Duration {
	return time.Duration(s.snd.Get("duration").Float()) * time.Second
}

// Pos returns the current playback position relative to the start.
func (s *Sound) Pos() time.Duration {
	return time.Duration(s.snd.Get("currentTime").Float()) * time.Second
}

// SetPos sets the current playback position relative to the start.
func (s *Sound) SetPos(pos time.Duration) {
	s.snd.Set("currentTime", pos.Seconds())
}
