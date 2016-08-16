package gogame

import "github.com/gopherjs/gopherjs/js"

// Styler applies a style to a canvas context.
type Styler interface {
	Style(ctx *js.Object)
	DrawType() DrawType
}

// DrawType is the method used to draw a shape (e.g. fill or stroke).
type DrawType string

const (
	// Fill fills in the shape.
	Fill DrawType = "fill"
	// Stroke draws the outline of the shape.
	Stroke DrawType = "stroke"
)

var _ Styler = &FillStyle{}

// FillStyle specifies a style used for filling shapes.
type FillStyle struct {
	Colorer
}

// Style implements the Styler interface.
func (f *FillStyle) Style(ctx *js.Object) {
	ctx.Set("fillStyle", f.Color(ctx))
}

// DrawType implements the Styler interface.
func (*FillStyle) DrawType() DrawType {
	return Fill
}

// LineCap is a style of line cap.
type LineCap string

const (
	// LineCapButt draws a line with no ends.
	LineCapButt LineCap = "butt"
	// LineCapRound draws a line with rounded ends with radius equal to half its width.
	LineCapRound LineCap = "round"
	// LineCapSquare draws a line with the ends capped with a box that extends by an amount
	// equal to half the lines width.
	LineCapSquare LineCap = "square"
)

// LineJoin is the style for the point where two lines are connected.
type LineJoin string

const (
	// LineJoinRound joins lines with rounded corners.
	LineJoinRound LineJoin = "round"
	// LineJoinBevel joins lines by filling in the triangular gap between them.
	LineJoinBevel LineJoin = "bevel"
	// LineJoinMiter joins lines by extending the edges until they meet.
	LineJoinMiter LineJoin = "miter"
)

var _ Styler = &StrokeStyle{}

// StrokeStyle specifies a style used for lines.
type StrokeStyle struct {
	Colorer
	Width      float64
	Cap        LineCap
	Join       LineJoin
	MiterLimit float64
	Dash       []float64
	DashOffset float64
}

// Style implements the Styler interface.
func (s *StrokeStyle) Style(ctx *js.Object) {
	ctx.Set("strokeStyle", s.Color(ctx))
	ctx.Set("lineWidth", s.Width)
	if s.Cap != "" {
		ctx.Set("lineCap", s.Cap)
	}
	if s.Join != "" {
		ctx.Set("lineJoin", s.Join)
	}
	if s.MiterLimit >= 1.0 {
		ctx.Set("miterLimit", s.MiterLimit)
	}
	if len(s.Dash) > 0 {
		ctx.Call("setLineDash", s.Dash)
		ctx.Set("lineDashOffset", s.DashOffset)
	}
}

// DrawType implements the Styler interface.
func (*StrokeStyle) DrawType() DrawType {
	return Stroke
}
