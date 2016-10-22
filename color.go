package gogame

// import (
// 	"math"
// 	"strconv"

// 	"github.com/gopherjs/gopherjs/js"
// )

// // DefaultColor is the color used for a nil Colorer.
// var DefaultColor Colorer = Transparent

// // Colorer is a type suitable to be used as a fill or line style.
// type Colorer interface {
// 	Color(ctx *js.Object) interface{}
// }

// var (
// 	Transparent = ColorCSS("rgba(0, 0, 0, 0)")
// 	Black       = ColorCSS("#000")
// 	White       = ColorCSS("#FFF")
// 	Red         = ColorCSS("#F00")
// 	Green       = ColorCSS("#0F0")
// 	Blue        = ColorCSS("#00B")
// )

// var _ Colorer = &Color{}

// // Color is a flat rgba color. Each component ranges from 0.0 to 1.0. An alpha (A) of
// // of 0.0 is fully transparent.
// type Color struct {
// 	R, G, B, A float64
// }

// // Color implements the Colorer interface.
// func (c Color) Color(*js.Object) interface{} {
// 	return c.String()
// }

// // String implements the Stringer interface. The format of the returned string is of a
// // CSS color, i.e. "rgba(r, g, b, a)".
// func (c Color) String() string {
// 	return "rgba(" + strconv.Itoa(clampToInt(255*c.R, 0, 255)) + "," +
// 		strconv.Itoa(clampToInt(255*c.G, 0, 255)) + "," +
// 		strconv.Itoa(clampToInt(255*c.B, 0, 255)) + "," +
// 		strconv.FormatFloat(clampToFloat(c.A, 0.0, 1.0), 'f', -1, 64) + ")"
// }

// func clampToInt(v, min, max float64) int {
// 	return int(math.Min(math.Max(v, min), max))
// }

// func clampToFloat(v, min, max float64) float64 {
// 	return math.Min(math.Max(v, min), max)
// }

// // ColorCSS is CSS style color, e.g. "rgb(1,2,3)", "#FFF". Using this type instead of
// // the Color type can be better for performance in certain cases since converting numbers
// // to strings can be expensive.
// type ColorCSS string

// // Color implements the Colorer interface.
// func (c ColorCSS) Color(*js.Object) interface{} {
// 	return string(c)
// }

// ColorStop is used for gradients to specify at which point it reaches a color. Position
// is from 0.0 to 1.0 and is its relative position in the gradient, 0.0 being the start
// and 1.0 being the end.
// type ColorStop struct {
// 	Position float64
// 	color.Color
// }

// var _ Colorer = &LinearGradient{}

// LinearGradient smoothly transitions between multiple colors in the direction defined
// by two points.
// type LinearGradient struct {
// 	X1, Y1, X2, Y2 float64
// 	ColorStops     []ColorStop
// }

// // Color implements the Colorer interface.
// func (l *LinearGradient) Color(ctx *js.Object) interface{} {
// 	lg := l
// 	if lg == nil {
// 		lg = &LinearGradient{}
// 	}
// 	grad := ctx.Call("createLinearGradient", math.Floor(lg.X1), math.Floor(lg.Y1), math.Floor(lg.X2),
// 		math.Floor(lg.Y2))
// 	for _, stop := range lg.ColorStops {
// 		grad.Call("addColorStop", stop.Position, stop.Colorer.Color(ctx))
// 	}
// 	return grad
// }

// var _ Colorer = &RadialGradient{}

// RadialGradient smoothly transitions between multiple colors from one circle to another.
// type RadialGradient struct {
// 	X1, Y1, R1, X2, Y2, R2 float64
// 	ColorStops             []ColorStop
// }

// // Color implements the Colorer interface.
// func (r *RadialGradient) Color(ctx *js.Object) interface{} {
// 	rg := r
// 	if rg == nil {
// 		rg = &RadialGradient{}
// 	}
// 	grad := ctx.Call("createRadialGradient", math.Floor(rg.X1), math.Floor(rg.Y1), math.Floor(rg.R1),
// 		math.Floor(rg.X2), math.Floor(rg.Y2), math.Floor(rg.R2))
// 	for _, stop := range rg.ColorStops {
// 		grad.Call("addColorStop", stop.Position, stop.Colorer.Color(ctx))
// 	}
// 	return grad
// }

// var _ Colorer = &Pattern{}

// // RepeatType describes how to repeat.
// type RepeatType string

// const (
// 	// Repeat repeats in both horizontal and vertical directions.
// 	Repeat RepeatType = "repeat"
// 	// RepeatX repeats in the horizontal direction.
// 	RepeatX RepeatType = "repeat-x"
// 	// RepeatY repeats in the vertical direction.
// 	RepeatY RepeatType = "repeat-y"
// 	// NoRepeat doesn't repeat.
// 	NoRepeat RepeatType = "no-repeat"
// )

// // Pattern is an image that optionally repeats.
// type Pattern struct {
// 	Source Surface
// 	Type   RepeatType
// }

// // Color implements the Colorer interface.
// func (p *Pattern) Color(ctx *js.Object) interface{} {
// 	if p == nil {
// 		return DefaultColor.Color(ctx)
// 	}
// 	if p.Type == "" {
// 		p.Type = Repeat
// 	}
// 	return ctx.Call("createPattern", p.Source.Canvas(), p.Type)
// }
