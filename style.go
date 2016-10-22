package gogame

// import "github.com/gopherjs/gopherjs/js"

// // Styler applies a style to a canvas context.
// type Styler interface {
// 	Style(ctx *js.Object)
// 	DrawType() DrawType
// }

// // DrawType is the method used to draw a shape (e.g. fill or stroke).
// type DrawType string

// const (
// 	// Fill fills in the shape.
// 	Fill DrawType = "fill"
// 	// Stroke draws the outline of the shape.
// 	Stroke DrawType = "stroke"
// )

// var _ Styler = &FillStyle{}

// // FillStyle specifies a style used for filling shapes. Using a nil FillStyle will use
// // the default Color.
// type FillStyle struct {
// 	Colorer
// }

// // Style implements the Styler interface.
// func (f *FillStyle) Style(ctx *js.Object) {
// 	color := DefaultColor.Color(ctx)
// 	if f != nil && f.Colorer != nil {
// 		color = f.Color(ctx)
// 	}
// 	ctx.Set("fillStyle", color)
// }

// // DrawType implements the Styler interface.
// func (*FillStyle) DrawType() DrawType {
// 	return Fill
// }

// var (
// 	// FillBlack is a fill style that is solid black.
// 	FillBlack = &FillStyle{Colorer: Black}
// 	// FillWhite is a fill style that is solid white.
// 	FillWhite = &FillStyle{Colorer: White}
// )

// // LineCap is a style of line cap.
// type LineCap string

// const (
// 	// LineCapButt draws a line with no ends.
// 	LineCapButt LineCap = "butt"
// 	// LineCapRound draws a line with rounded ends with radius equal to half its width.
// 	LineCapRound LineCap = "round"
// 	// LineCapSquare draws a line with the ends capped with a box that extends by an amount
// 	// equal to half the lines width.
// 	LineCapSquare LineCap = "square"
// )

// // LineJoin is the style for the point where two lines are connected.
// type LineJoin string

// const (
// 	// LineJoinRound joins lines with rounded corners.
// 	LineJoinRound LineJoin = "round"
// 	// LineJoinBevel joins lines by filling in the triangular gap between them.
// 	LineJoinBevel LineJoin = "bevel"
// 	// LineJoinMiter joins lines by extending the edges until they meet.
// 	LineJoinMiter LineJoin = "miter"
// )

// var _ Styler = &StrokeStyle{}

// // StrokeStyle specifies a style used for lines. A nil StrokeStyle will use default values.
// type StrokeStyle struct {
// 	Colorer
// 	Width      float64
// 	Cap        LineCap
// 	Join       LineJoin
// 	MiterLimit float64
// 	Dash       []float64
// 	DashOffset float64
// }

// // Style implements the Styler interface.
// func (s *StrokeStyle) Style(ctx *js.Object) {
// 	style := s
// 	if style == nil {
// 		style = &StrokeStyle{}
// 	}
// 	color := DefaultColor.Color(ctx)
// 	if style.Colorer != nil {
// 		color = s.Color(ctx)
// 	}
// 	ctx.Set("strokeStyle", color)
// 	ctx.Set("lineWidth", style.Width)
// 	if style.Cap != "" {
// 		ctx.Set("lineCap", style.Cap)
// 	}
// 	if style.Join != "" {
// 		ctx.Set("lineJoin", style.Join)
// 	}
// 	if style.MiterLimit >= 1.0 {
// 		ctx.Set("miterLimit", style.MiterLimit)
// 	}
// 	if len(style.Dash) > 0 {
// 		ctx.Call("setLineDash", style.Dash)
// 		ctx.Set("lineDashOffset", style.DashOffset)
// 	}
// }

// // DrawType implements the Styler interface.
// func (*StrokeStyle) DrawType() DrawType {
// 	return Stroke
// }
