package ggweb

import (
	"image/color"
	"strconv"
)

// ColorToCSS converts a color to a CSS formatted color string, e.g. "rgba(0, 0, 0, 0)"
func ColorToCSS(c color.Color) string {
	r, g, b, a := c.RGBA()
	return "rgba(" + strconv.Itoa(mapColor(r)) + "," + strconv.Itoa(mapColor(g)) + "," +
		strconv.Itoa(mapColor(b)) + "," +
		strconv.FormatFloat(float64(mapColor(a))/255, 'f', -1, 64) + ")"
}

func mapColor(i uint32) int {
	return int(float64(i) / float64(0xffff) * 0xff)
}

// ColorStop is used for gradients to specify at which point it reaches a color. Position
// is from 0.0 to 1.0 and is its relative position in the gradient, 0.0 being the start
// and 1.0 being the end.
type ColorStop struct {
	Position float64
	color.Color
}

// LinearGradient smoothly transitions between multiple colors in the direction defined
// by two points.
type LinearGradient struct {
	X1, Y1, X2, Y2 float64
	ColorStops     []ColorStop
}

// RadialGradient smoothly transitions between multiple colors from one circle to another.
type RadialGradient struct {
	X1, Y1, R1, X2, Y2, R2 float64
	ColorStops             []ColorStop
}

// RepeatType describes how to repeat.
type RepeatType string

const (
	// RepeatXY repeats in both horizontal and vertical directions.
	RepeatXY RepeatType = "repeat"
	// RepeatX repeats in the horizontal direction.
	RepeatX RepeatType = "repeat-x"
	// RepeatY repeats in the vertical direction.
	RepeatY RepeatType = "repeat-y"
	// NoRepeat doesn't repeat.
	NoRepeat RepeatType = "no-repeat"
)

// Pattern is an image that optionally repeats. The default Type is RepeatXY.
type Pattern struct {
	Source *Surface
	Type   RepeatType
}
