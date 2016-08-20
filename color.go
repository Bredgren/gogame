package gogame

import (
	"fmt"
	"math"

	"github.com/gopherjs/gopherjs/js"
)

// DefaultColor is the color used for a nil Colorer.
var DefaultColor Colorer = Transparent

// Colorer is a type suitable to be used as a fill or line style.
type Colorer interface {
	Color(ctx *js.Object) interface{}
}

var _ Colorer = &Color{}

// Color is a flat rgba color. Each component ranges from 0.0 to 1.0. An alpha (A) of
// of 0.0 is fully transparent.
type Color struct {
	R, G, B, A float64
}

var (
	// Transparent is no color (the default).
	Transparent = Color{}
	// Black is the color black.
	Black = Color{0.0, 0.0, 0.0, 1.0}
	// White is the color white.
	White = Color{1.0, 1.0, 1.0, 1.0}
	// Red is the color red.
	Red = Color{1.0, 0.0, 0.0, 1.0}
	// Green is the color green.
	Green = Color{0.0, 1.0, 0.0, 1.0}
	// Blue is the color blue.
	Blue = Color{0.0, 0.0, 1.0, 1.0}
)

// Color implements the Colorer interface.
func (c Color) Color(*js.Object) interface{} {
	return c.String()
}

// String implements the Stringer interface. The format of the returned string is of a
// CSS color, i.e. "rgba(r, g, b, a)".
func (c Color) String() string {
	return fmt.Sprintf("rgba(%d, %d, %d, %f)",
		clampToInt(255*c.R, 0, 255),
		clampToInt(255*c.G, 0, 255),
		clampToInt(255*c.B, 0, 255),
		clampToFloat(c.A, 0.0, 1.0))
}

func clampToInt(v, min, max float64) int {
	return int(math.Min(math.Max(v, min), max))
}

func clampToFloat(v, min, max float64) float64 {
	return math.Min(math.Max(v, min), max)
}

// ColorStop is used for gradients to specify at which point it reaches a color. Position
// is from 0.0 to 1.0 and is its relative position in the gradient, 0.0 being the start
// and 1.0 being the end.
type ColorStop struct {
	Position float64
	Color
}

var _ Colorer = &LinearGradient{}

// LinearGradient smoothly transitions between multiple colors in the direction defined
// by two points.
type LinearGradient struct {
	X1, Y1, X2, Y2 float64
	ColorStops     []ColorStop
}

// Color implements the Colorer interface.
func (l *LinearGradient) Color(ctx *js.Object) interface{} {
	lg := l
	if lg == nil {
		lg = &LinearGradient{}
	}
	grad := ctx.Call("createLinearGradient", math.Floor(lg.X1), math.Floor(lg.Y1), math.Floor(lg.X2),
		math.Floor(lg.Y2))
	for _, stop := range lg.ColorStops {
		grad.Call("addColorStop", stop.Position, stop.Color.Color(ctx))
	}
	return grad
}

var _ Colorer = &RadialGradient{}

// RadialGradient smoothly transitions between multiple colors from one circle to another.
type RadialGradient struct {
	X1, Y1, R1, X2, Y2, R2 float64
	ColorStops             []ColorStop
}

// Color implements the Colorer interface.
func (r *RadialGradient) Color(ctx *js.Object) interface{} {
	rg := r
	if rg == nil {
		rg = &RadialGradient{}
	}
	grad := ctx.Call("createRadialGradient", math.Floor(rg.X1), math.Floor(rg.Y1), math.Floor(rg.R1),
		math.Floor(rg.X2), math.Floor(rg.Y2), math.Floor(rg.R2))
	for _, stop := range rg.ColorStops {
		grad.Call("addColorStop", stop.Position, stop.Color.Color(ctx))
	}
	return grad
}

var _ Colorer = &Pattern{}

// RepeatType describes how to repeat.
type RepeatType string

const (
	// Repeat repeats in both horizontal and vertical directions.
	Repeat RepeatType = "repeat"
	// RepeatX repeats in the horizontal direction.
	RepeatX RepeatType = "repeat-x"
	// RepeatY repeats in the vertical direction.
	RepeatY RepeatType = "repeat-y"
	// NoRepeat doesn't repeat.
	NoRepeat RepeatType = "no-repeat"
)

// Pattern is an image that optionally repeats.
type Pattern struct {
	Source Surface
	Type   RepeatType
}

// Color implements the Colorer interface.
func (p *Pattern) Color(ctx *js.Object) interface{} {
	if p == nil {
		return DefaultColor.Color(ctx)
	}
	if p.Type == "" {
		p.Type = Repeat
	}
	return ctx.Call("createPattern", p.Source.GetCanvas(), p.Type)
}
