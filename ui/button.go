// Package ui provides some simple common user interface components.
package ui

import (
	"github.com/Bredgren/gogame"
	"github.com/Bredgren/gogame/geo"
)

// ButtonState is a state that a button can be in. Specific states are left for the user
// to define.
type ButtonState int

// BasicButton is a selectable surface. It has multiple states and optionally has different
// surfaces associated with each one. Collision and state changes are left to the user.
type BasicButton struct {
	// Rect is the position and size of the button.
	Rect geo.Rect
	// DefaultSurf is the Surface to use when there is no Surface in StateSurfs for the current state.
	DefaultSurf gogame.Surface
	// StateSurfs map a ButtonState to a surface to use for that state.
	StateSurfs map[ButtonState]gogame.Surface
	// Select is the function to call when the Button is selected.
	Select func()
	// State is the Button's current state.
	State ButtonState
}

// Surface returns the Surface for the current State. If there is no surface for the current
// state then the DefaultSurf is returned. If that is not set either then a blank Surface
// of size 0 is returned.
func (b *BasicButton) Surface() gogame.Surface {
	s, ok := b.StateSurfs[b.State]
	if !ok {
		if b.DefaultSurf != nil {
			return b.DefaultSurf
		}
		return gogame.NewSurface(0, 0)
	}
	return s
}
