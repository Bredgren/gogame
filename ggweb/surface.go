package ggweb

import (
	"fmt"
	"image/color"
	"math"

	"github.com/Bredgren/gogame/geo"
	"github.com/gopherjs/gopherjs/js"
)

// Surface is wrapper for a canvas and its context. Generally you'll want to use one of
// the NewSurface* functions to create create the Surface. If initializing manually keep
// in mind that Surface's functions assume Canvas and Ctx are valid. Most of Surface's
// functions deal with float64 for positioning but, unless otherwise stated, values are
// floored before drawing.
type Surface struct {
	Canvas *js.Object
	Ctx    *js.Object
}

// NewSurface creates a new Surface with the given dimensions.
func NewSurface(width, height int) *Surface {
	canvas := js.Global.Get("document").Call("createElement", "canvas")
	canvas.Set("width", width)
	canvas.Set("height", height)
	return &Surface{
		Canvas: canvas,
		Ctx:    canvas.Call("getContext", "2d"),
	}
}

// NewSurfaceFromCanvas creates a new Surface from the given canvas. It panics if given
// a nil or undefined canvas.
func NewSurfaceFromCanvas(canvas *js.Object) *Surface {
	if canvas == nil || canvas == js.Undefined {
		panic(fmt.Sprintf("invalid canvas: %v", canvas))
	}
	return &Surface{
		Canvas: canvas,
		Ctx:    canvas.Call("getContext", "2d"),
	}
}

// NewSurfaceFromID creates a new Surface using the canvas with the given ID. It panics
// if no canvas with the given ID is found. Including the '#' is optional.
func NewSurfaceFromID(canvasID string) *Surface {
	// if !strings.HasPrefix(canvasID, "#") {
	// 	canvasID = "#" + canvasID
	// }
	canvas := js.Global.Get("document").Call("getElementById", canvasID)
	if canvas == nil || canvas == js.Undefined {
		panic(fmt.Sprintf("no canvas found with ID '%s'", canvasID))
	}
	return &Surface{
		Canvas: canvas,
		Ctx:    canvas.Call("getContext", "2d"),
	}
}

// Rect returns the rectangle that covers the whole surface.
func (s *Surface) Rect() geo.Rect {
	return geo.Rect{
		W: s.Canvas.Get("width").Float(),
		H: s.Canvas.Get("height").Float(),
	}
}

// SetSize resizes the surface.
func (s *Surface) SetSize(w, h int) {
	s.Canvas.Set("width", w)
	s.Canvas.Set("height", h)
	// s.Ctx = s.Canvas.Call("getContext", "2d")
}

// Blit draws the source surface to s with source's top left corner at x, y.
func (s *Surface) Blit(source *Surface, x, y float64) {
	s.Ctx.Call("drawImage", source.Canvas, math.Floor(x), math.Floor(y))
}

// BlitArea draws the sub-region of source defined by area to s with the top left corner
// at x, y.
func (s *Surface) BlitArea(source *Surface, area geo.Rect, x, y float64) {
	s.Ctx.Call("drawImage", source.Canvas,
		math.Floor(area.X), math.Floor(area.Y), math.Floor(area.W), math.Floor(area.H),
		math.Floor(x), math.Floor(y), math.Floor(area.W), math.Floor(area.H))
}

// Save saves the current context.
func (s *Surface) Save() {
	s.Ctx.Call("save")
}

// Restore restores the last saved context.
func (s *Surface) Restore() {
	s.Ctx.Call("restore")
}

// StyleColor sets the fill/stoke to a solid color.
func (s *Surface) StyleColor(t DrawType, c color.Color) {
	s.Ctx.Set(string(t)+"Style", ColorToCSS(c))
}

// StyleLinearGradient sets the fill/stroke style to a linear gradient.
func (s *Surface) StyleLinearGradient(t DrawType, g LinearGradient) {
	grad := s.Ctx.Call("createLinearGradient", math.Floor(g.X1), math.Floor(g.Y1), math.Floor(g.X2),
		math.Floor(g.Y2))
	for _, stop := range g.ColorStops {
		grad.Call("addColorStop", stop.Position, ColorToCSS(stop.Color))
	}
	s.Ctx.Set(string(t)+"Style", grad)
}

// StyleRadialGradient sets the fill/stroke style to a radial gradient.
func (s *Surface) StyleRadialGradient(t DrawType, g RadialGradient) {
	grad := s.Ctx.Call("createRadialGradient", math.Floor(g.X1), math.Floor(g.Y1), math.Floor(g.R1),
		math.Floor(g.X2), math.Floor(g.Y2), math.Floor(g.R2))
	for _, stop := range g.ColorStops {
		grad.Call("addColorStop", stop.Position, ColorToCSS(stop.Color))
	}
	s.Ctx.Set(string(t)+"Style", grad)
}

// StylePattern sets the fill/stoke style to a pattern.
func (s *Surface) StylePattern(t DrawType, p Pattern) {
	if p.Type == "" {
		p.Type = RepeatXY
	}
	pat := s.Ctx.Call("createPattern", p.Source.Canvas, p.Type)
	s.Ctx.Set(string(t)+"Style", pat)
}

// // SetLineProps sets the properties of lines and strokes.
// func (s *Surface) SetLineProps(cap, join, width, miter) {
// }

// DrawRect draws a rectangle on the surface.
func (s *Surface) DrawRect(t DrawType, r geo.Rect) {
	s.Ctx.Call(string(t)+"Rect", math.Floor(r.X), math.Floor(r.Y), math.Floor(r.W), math.Floor(r.H))
}

// ClearRect clears the area within the rectangle.
func (s *Surface) ClearRect(r geo.Rect) {
	s.Ctx.Call("clearRect", math.Floor(r.X), math.Floor(r.Y), math.Floor(r.W), math.Floor(r.H))
}

// DrawCircle draws a circle on the surface.
func (s *Surface) DrawCircle(t DrawType, x, y, radius float64) {
	s.Ctx.Call("beginPath")
	s.Ctx.Call("arc", math.Floor(x), math.Floor(y), math.Floor(radius), 0, 2*math.Pi)
	s.Ctx.Call(string(t))
}

// DrawEllipse draws an ellipse on the surface within the given rectangle.
func (s *Surface) DrawEllipse(t DrawType, r geo.Rect) {
	s.Ctx.Call("beginPath")
	s.Ctx.Call("ellipse", math.Floor(r.CenterX()), math.Floor(r.CenterY()),
		math.Floor(r.W/2), math.Floor(r.H/2), 0, 0, 2*math.Pi)
	s.Ctx.Call(string(t))
}

// DrawArc draws an arc on the surface, i.e. any slice of an ellipse. The angles are
// counterclockwise relative to the +x axis. The counterclockwise parameter is for the
// direction to draw in.
func (s *Surface) DrawArc(t DrawType, r geo.Rect, startRadians, endRadians float64, counterclockwise bool) {
	s.Ctx.Call("beginPath")
	s.Ctx.Call("ellipse", math.Floor(r.CenterX()), math.Floor(r.CenterY()), math.Floor(r.W/2),
		math.Floor(r.H/2), 0, 2*math.Pi-startRadians, 2*math.Pi-endRadians, counterclockwise)
	s.Ctx.Call(string(t))
}

// DrawPath draws the given path object to the surface.
func (s *Surface) DrawPath(t DrawType, p *Path) {
	s.Ctx.Call(string(t), p.obj)
}

// ClipPath sets the clipping area of the surface to be within the path. Note that if one
// wishes to only temporarily set the clip area then Save must be called before ClipPath and
// Restore after the desired drawing operations, there is no other way to reset the clip area.
func (s *Surface) ClipPath(p *Path) {
	s.Ctx.Call("clip", p.obj)
}

// func (s *Surface) Scale(x, y)
// func (s *Surface) Rotate(angle)
// func (s *Surface) Translate(x, y)
// func (s *Surface) (Set)Transform(...)
//  + *ed versions for each that return new Surface

// func (s *Surface) SetFontProps(...)
// func (s *Surface) SetTextAlign(align, baseline)
// func (s *Surface) DrawText(FillType, text)
// func (s *Surface) MeasureText(text)

// func (s *Surface) PixeiData() []color.Color?
// func (s *Surface) SetAt(x, y, color.Color)
// func (s *Surface) GetAt(x, y)

// func (s *Surface) (Set)Alpha()
// func (s *Surface) (Set)CompositeOp()

// ... stuff from canvas.go

////////////////////////////////////////////////////////////////////////////////////
// func (s *surface) Fill(style *FillStyle) {
// 	s.ctx.Call("save")
// 	style.Style(s.ctx)
// 	s.ctx.Call("fillRect", 0, 0, s.canvas.Get("width"), s.canvas.Get("height"))
// 	s.ctx.Call("restore")
// }

// func (s *surface) GetAt(x, y int) Color {
// 	data := s.ctx.Call("getImageData", x, y, 1, 1).Get("data")
// 	return Color{R: data.Index(0).Float(), G: data.Index(1).Float(), B: data.Index(2).Float(),
// 		A: data.Index(3).Float()}
// }

// func (s *surface) SetAt(x, y int, c Color) {
// 	imgData := s.ctx.Call("getImageData", x, y, 1, 1)
// 	data := imgData.Get("data")
// 	data.SetIndex(0, clampToInt(255*c.R, 0, 255))
// 	data.SetIndex(1, clampToInt(255*c.G, 0, 255))
// 	data.SetIndex(2, clampToInt(255*c.B, 0, 255))
// 	data.SetIndex(3, clampToInt(255*c.A, 0, 255))
// 	s.ctx.Call("putImageData", imgData, x, y)
// }

// func (s *surface) SetClip(r geo.Rect) {
// 	s.SetClipPath([][2]float64{
// 		{r.X, r.Y}, {r.X + r.W, r.Y}, {r.X + r.W, r.Y + r.H}, {r.X, r.Y + r.H},
// 	})
// }

// func (s *surface) SetClipPath(pointList [][2]float64) {
// 	s.ctx.Call("save")
// 	s.ctx.Call("beginPath")
// 	s.ctx.Call("moveTo", pointList[0][0], pointList[0][1])
// 	for _, p := range pointList[1:] {
// 		s.ctx.Call("lineTo", p[0], p[1])
// 	}
// 	s.ctx.Call("clip")
// }

// func (s *surface) ClearClip() {
// 	s.ctx.Call("restore")
// }

// func (s *surface) Copy() Surface {
// 	copy := NewSurface(s.Width(), s.Height())
// 	copy.Blit(s, 0, 0)
// 	return copy
// }

// func (s *surface) Scaled(x, y float64) Surface {
// 	newS := NewSurface(int(float64(s.Width())*x), int(float64(s.Height())*y))
// 	ctx := newS.(*surface).ctx
// 	ctx.Call("save")
// 	ctx.Call("scale", x, y)
// 	ctx.Call("drawImage", s.canvas, 0, 0)
// 	ctx.Call("restore")
// 	return newS
// }

// func (s *surface) Rotated(radians float64) Surface {
// 	newW, newH := s.getRotatedSize(radians)
// 	newS := NewSurface(newW, newH)
// 	ctx := newS.(*surface).ctx
// 	ctx.Call("save")
// 	cx, cy := newW/2, newH/2
// 	ctx.Call("translate", cx, cy)
// 	ctx.Call("rotate", -radians)
// 	ctx.Call("translate", -s.Width()/2, -s.Height()/2)
// 	ctx.Call("drawImage", s.canvas, 0, 0)
// 	ctx.Call("restore")
// 	return newS
// }

// func (s *surface) getRotatedSize(radians float64) (w, h int) {
// 	width, height := float64(s.Width()), float64(s.Height())
// 	cx, cy := width/2, height/2
// 	cos, sin := math.Cos(radians), math.Sin(radians)

// 	x1 := cx + (0-cx)*cos + (0-cy)*sin
// 	y1 := cy - (0-cx)*sin + (0-cy)*cos
// 	x2 := cx + (width-cx)*cos + (0-cy)*sin
// 	y2 := cy - (width-cx)*sin + (0-cy)*cos
// 	x3 := cx + (0-cx)*cos + (height-cy)*sin
// 	y3 := cy - (0-cx)*sin + (height-cy)*cos
// 	x4 := cx + (width-cx)*cos + (height-cy)*sin
// 	y4 := cy - (width-cx)*sin + (height-cy)*cos

// 	maxX := math.Max(x1, math.Max(x2, math.Max(x3, x4)))
// 	minX := math.Min(x1, math.Min(x2, math.Min(x3, x4)))
// 	maxY := math.Max(y1, math.Max(y2, math.Max(y3, y4)))
// 	minY := math.Min(y1, math.Min(y2, math.Min(y3, y4)))

// 	return int(maxX - minX), int(maxY - minY)
// }

// func (s *surface) DrawRect(r geo.Rect, style Styler) {
// 	if style == nil {
// 		style = &FillStyle{}
// 	}
// 	s.ctx.Call("save")
// 	style.Style(s.ctx)
// 	s.ctx.Call("translate", math.Floor(r.X), math.Floor(r.Y))
// 	s.ctx.Call(fmt.Sprintf("%sRect", style.DrawType()), 0, 0, math.Floor(r.W), math.Floor(r.H))
// 	s.ctx.Call("restore")
// }

// func (s *surface) DrawCircle(posX, posY, radius float64, style Styler) {
// 	if style == nil {
// 		style = &FillStyle{}
// 	}
// 	s.ctx.Call("save")
// 	style.Style(s.ctx)
// 	s.ctx.Call("translate", math.Floor(posX), math.Floor(posY))
// 	s.ctx.Call("beginPath")
// 	s.ctx.Call("arc", 0, 0, radius, 0, 2*math.Pi)
// 	s.ctx.Call(string(style.DrawType()))
// 	s.ctx.Call("restore")
// }

// func (s *surface) DrawEllipse(r geo.Rect, style Styler) {
// 	if style == nil {
// 		style = &FillStyle{}
// 	}
// 	s.ctx.Call("save")
// 	style.Style(s.ctx)
// 	s.ctx.Call("translate", math.Floor(r.X), math.Floor(r.Y))
// 	s.ctx.Call("beginPath")
// 	s.ctx.Call("ellipse", math.Floor(r.Width()/2), math.Floor(r.Height()/2), math.Floor(r.Width()/2),
// 		math.Floor(r.Height()/2), 0, 0, 2*math.Pi)
// 	s.ctx.Call(string(style.DrawType()))
// 	s.ctx.Call("restore")
// }

// func (s *surface) DrawArc(r geo.Rect, startRadians, stopRadians float64, style Styler) {
// 	if style == nil {
// 		style = &StrokeStyle{}
// 	}
// 	s.ctx.Call("save")
// 	style.Style(s.ctx)
// 	s.ctx.Call("translate", math.Floor(r.X), math.Floor(r.Y))
// 	s.ctx.Call("beginPath")
// 	s.ctx.Call("ellipse", math.Floor(r.Width()/2), math.Floor(r.Height()/2), math.Floor(r.Width()/2),
// 		math.Floor(r.Height()/2), 0, 2*math.Pi-startRadians, 2*math.Pi-stopRadians, true)
// 	s.ctx.Call(string(style.DrawType()))
// 	s.ctx.Call("restore")
// }

// func (s *surface) DrawLine(startX, startY, endX, endY float64, style Styler) {
// 	if style == nil {
// 		style = &StrokeStyle{}
// 	}
// 	s.ctx.Call("save")
// 	style.Style(s.ctx)
// 	s.ctx.Call("beginPath")
// 	// Not math.Flooring lines to preserve control with odd width lines.
// 	s.ctx.Call("moveTo", startX, startY)
// 	s.ctx.Call("lineTo", endX, endY)
// 	s.ctx.Call(string(style.DrawType()))
// 	s.ctx.Call("restore")
// }

// func (s *surface) DrawLines(points [][2]float64, style Styler) {
// 	if style == nil {
// 		style = &StrokeStyle{}
// 	}
// 	s.ctx.Call("save")
// 	style.Style(s.ctx)
// 	s.ctx.Call("beginPath")
// 	// Not math.Flooring lines to preserve control with odd width lines.
// 	s.ctx.Call("moveTo", points[0][0], points[0][1])
// 	for _, p := range points[1:] {
// 		s.ctx.Call("lineTo", p[0], p[1])
// 	}
// 	s.ctx.Call(string(style.DrawType()))
// 	s.ctx.Call("restore")
// }

// func (s *surface) DrawText(text string, x, y float64, font *Font, style *TextStyle) {
// 	s.ctx.Call("save")
// 	s.ctx.Set("font", font.String())
// 	style.Style(s.ctx)
// 	s.ctx.Call("translate", math.Floor(x), math.Floor(y))
// 	s.ctx.Call(fmt.Sprintf("%sText", style.DrawType()), text, 0, 0)
// 	s.ctx.Call("restore")
// }

// func (s *surface) DrawQuadraticCurve(startX, startY, endX, endY, cpX, cpY float64, style Styler) {
// 	if style == nil {
// 		style = &StrokeStyle{}
// 	}
// 	s.ctx.Call("save")
// 	style.Style(s.ctx)
// 	s.ctx.Call("beginPath")
// 	// Not math.Flooring lines to preserve control with odd width lines.
// 	s.ctx.Call("moveTo", startX, startY)
// 	s.ctx.Call("quadraticCurveTo", cpX, cpY, endX, endY)
// 	s.ctx.Call(string(style.DrawType()))
// 	s.ctx.Call("restore")
// }

// func (s *surface) DrawQuadraticCurves(points [][2]float64, style Styler) {
// 	if len(points) < 3 {
// 		return // Not enough points for event one curve.
// 	}
// 	if style == nil {
// 		style = &StrokeStyle{}
// 	}
// 	s.ctx.Call("save")
// 	style.Style(s.ctx)
// 	s.ctx.Call("beginPath")
// 	// Not math.Flooring lines to preserve control with odd width lines.
// 	s.ctx.Call("moveTo", points[0][0], points[0][1])
// 	for i := 1; i+1 < len(points); i += 2 {
// 		s.ctx.Call("quadraticCurveTo", points[i][0], points[i][1], points[i+1][0], points[i+1][1])
// 	}
// 	s.ctx.Call(string(style.DrawType()))
// 	s.ctx.Call("restore")
// }

// func (s *surface) DrawBezierCurve(startX, startY, endX, endY, cpStartX, cpStartY, cpEndX, cpEndY float64, style Styler) {
// 	if style == nil {
// 		style = &StrokeStyle{}
// 	}
// 	s.ctx.Call("save")
// 	style.Style(s.ctx)
// 	s.ctx.Call("beginPath")
// 	// Not math.Flooring lines to preserve control with odd width lines.
// 	s.ctx.Call("moveTo", startX, startY)
// 	s.ctx.Call("bezierCurveTo", cpStartX, cpStartY, cpEndX, cpEndY, endX, endY)
// 	s.ctx.Call(string(style.DrawType()))
// 	s.ctx.Call("restore")
// }

// func (s *surface) DrawBezierCurves(points [][2]float64, style Styler) {
// 	if len(points) < 4 {
// 		return // Not enough points for event one curve.
// 	}
// 	if style == nil {
// 		style = &StrokeStyle{}
// 	}
// 	s.ctx.Call("save")
// 	style.Style(s.ctx)
// 	s.ctx.Call("beginPath")
// 	// Not math.Flooring lines to preserve control with odd width lines.
// 	s.ctx.Call("moveTo", points[0][0], points[0][1])
// 	for i := 1; i+2 < len(points); i += 3 {
// 		s.ctx.Call("bezierCurveTo", points[i][0], points[i][1], points[i+1][0], points[i+1][1],
// 			points[i+2][0], points[i+2][1])
// 	}
// 	s.ctx.Call(string(style.DrawType()))
// 	s.ctx.Call("restore")
// }

// var _ Surface = &subsurface{}

// type subsurface struct {
// 	area   geo.Rect
// 	parent Surface
// }

// func (s *subsurface) Canvas() *js.Object {
// 	return s.parent.Canvas()
// }

// func (s *subsurface) Blit(source Surface, x, y float64) {
// 	s.parent.SetClip(s.area)
// 	s.parent.Blit(source, s.area.X+x, s.area.Y+y)
// 	s.parent.ClearClip()
// }

// func (s *subsurface) BlitArea(source Surface, area geo.Rect, x, y float64) {
// 	s.parent.SetClip(s.area)
// 	s.parent.BlitArea(source, area, s.area.X+x, s.area.Y+y)
// 	s.parent.ClearClip()
// }

// func (s *subsurface) BlitComp(source Surface, x, y float64, c composite.Op) {
// 	s.parent.SetClip(s.area)
// 	s.parent.BlitComp(source, s.area.X+x, s.area.Y+y, c)
// 	s.parent.ClearClip()
// }

// func (s *subsurface) BlitAreaComp(source Surface, area geo.Rect, x, y float64, c composite.Op) {
// 	s.parent.SetClip(s.area)
// 	s.parent.BlitAreaComp(source, area, s.area.X+x, s.area.Y+y, c)
// 	s.parent.ClearClip()
// }

// func (s *subsurface) Fill(style *FillStyle) {
// 	s.parent.DrawRect(s.area, style)
// }

// func (s *subsurface) Width() int {
// 	return int(s.area.W)
// }

// func (s *subsurface) Height() int {
// 	return int(s.area.H)
// }

// func (s *subsurface) Copy() Surface {
// 	canvas := js.Global.Get("document").Call("createElement", "canvas")
// 	canvas.Set("width", int(s.area.W))
// 	canvas.Set("height", int(s.area.H))
// 	ctx := canvas.Call("getContext", "2d")
// 	ctx.Call("drawImage", s.Canvas(), math.Floor(s.area.X), math.Floor(s.area.Y),
// 		math.Floor(s.area.W), math.Floor(s.area.H), 0, 0, math.Floor(s.area.W), math.Floor(s.area.H))
// 	// Ignoring impossible error since we know we're giving it a valid canvas
// 	copy, _ := NewSurfaceFromCanvas(canvas)
// 	return copy
// }

// func (s *subsurface) SubSurface(area geo.Rect) Surface {
// 	return &subsurface{
// 		area:   area,
// 		parent: s,
// 	}
// }

// func (s *subsurface) Parent() Surface {
// 	return s.parent
// }

// func (s *subsurface) GetAt(x, y int) Color {
// 	return s.parent.GetAt(int(s.area.X)+x, int(s.area.Y)+y)
// }

// func (s *subsurface) SetAt(x, y int, c Color) {
// 	s.parent.SetClip(s.area)
// 	s.parent.SetAt(int(s.area.X)+x, int(s.area.Y)+y, c)
// 	s.parent.ClearClip()
// }

// func (s *subsurface) SetClip(area geo.Rect) {
// 	s.parent.SetClip(s.area.Intersect(area))
// }

// func (s *subsurface) SetClipPath(pointList [][2]float64) {
// 	for i := range pointList {
// 		pointList[i][0] += s.area.X
// 		pointList[i][1] += s.area.Y
// 	}
// 	s.parent.SetClipPath(pointList)
// }

// func (s *subsurface) ClearClip() {
// 	s.parent.ClearClip()
// }

// func (s *subsurface) Scaled(x, y float64) Surface {
// 	surf := s.Copy()
// 	return surf.Scaled(x, y)
// }

// func (s *subsurface) Rotated(radians float64) Surface {
// 	surf := s.Copy()
// 	return surf.Rotated(radians)
// }

// func (s *subsurface) Rect() geo.Rect {
// 	return geo.Rect{X: 0, Y: 0, W: float64(s.Width()), H: float64(s.Height())}
// }

// func (s *subsurface) DrawRect(r geo.Rect, style Styler) {
// 	s.parent.SetClip(s.area)
// 	r.X += s.area.X
// 	r.Y += s.area.Y
// 	s.parent.DrawRect(r, style)
// 	s.parent.ClearClip()
// }

// func (s *subsurface) DrawCircle(posX, posY, radius float64, style Styler) {
// 	s.parent.SetClip(s.area)
// 	s.parent.DrawCircle(posX+s.area.X, posY+s.area.Y, radius, style)
// 	s.parent.ClearClip()
// }

// func (s *subsurface) DrawEllipse(r geo.Rect, style Styler) {
// 	s.parent.SetClip(s.area)
// 	r.X += s.area.X
// 	r.Y += s.area.Y
// 	s.parent.DrawEllipse(r, style)
// 	s.parent.ClearClip()
// }

// func (s *subsurface) DrawArc(r geo.Rect, startRadians, stopRadians float64, style Styler) {
// 	s.parent.SetClip(s.area)
// 	r.X += s.area.X
// 	r.Y += s.area.Y
// 	s.parent.DrawArc(r, startRadians, stopRadians, style)
// 	s.parent.ClearClip()
// }

// func (s *subsurface) DrawLine(startX, startY, endX, endY float64, style Styler) {
// 	s.parent.SetClip(s.area)
// 	startX += s.area.X
// 	startY += s.area.Y
// 	endX += s.area.X
// 	endY += s.area.Y
// 	s.parent.DrawLine(startX, startY, endX, endY, style)
// 	s.parent.ClearClip()
// }

// func (s *subsurface) DrawLines(points [][2]float64, style Styler) {
// 	s.parent.SetClip(s.area)
// 	for i := range points {
// 		points[i][0] += s.area.X
// 		points[i][1] += s.area.Y
// 	}
// 	s.parent.DrawLines(points, style)
// 	s.parent.ClearClip()
// }

// func (s *subsurface) DrawText(text string, x, y float64, font *Font, style *TextStyle) {
// 	s.parent.SetClip(s.area)
// 	s.parent.DrawText(text, x+s.area.X, y+s.area.Y, font, style)
// 	s.parent.ClearClip()
// }

// func (s *subsurface) DrawQuadraticCurve(startX, startY, endX, endY, cpX, cpY float64, style Styler) {
// 	s.parent.SetClip(s.area)
// 	startX += s.area.X
// 	startY += s.area.Y
// 	endX += s.area.X
// 	endY += s.area.Y
// 	cpX += s.area.X
// 	cpY += s.area.Y
// 	s.parent.DrawQuadraticCurve(startX, startY, endX, endY, cpX, cpY, style)
// 	s.parent.ClearClip()
// }

// func (s *subsurface) DrawQuadraticCurves(points [][2]float64, style Styler) {
// 	s.parent.SetClip(s.area)
// 	for i := range points {
// 		points[i][0] += s.area.X
// 		points[i][1] += s.area.Y
// 	}
// 	s.parent.DrawQuadraticCurves(points, style)
// 	s.parent.ClearClip()
// }

// func (s *subsurface) DrawBezierCurve(startX, startY, endX, endY, cpStartX, cpStartY, cpEndX, cpEndY float64, style Styler) {
// 	s.parent.SetClip(s.area)
// 	startX += s.area.X
// 	startY += s.area.Y
// 	endX += s.area.X
// 	endY += s.area.Y
// 	cpStartX += s.area.X
// 	cpStartY += s.area.Y
// 	cpEndX += s.area.X
// 	cpEndY += s.area.Y
// 	s.parent.DrawBezierCurve(startX, startY, endX, endY, cpStartX, cpStartY, cpEndX, cpEndY, style)
// 	s.parent.ClearClip()
// }

// func (s *subsurface) DrawBezierCurves(points [][2]float64, style Styler) {
// 	s.parent.SetClip(s.area)
// 	for i := range points {
// 		points[i][0] += s.area.X
// 		points[i][1] += s.area.Y
// 	}
// 	s.parent.DrawBezierCurves(points, style)
// 	s.parent.ClearClip()
// }

// // BlitComp is like Blit but uses the given composite operation.
// func (s *Surface) BlitComp(source *Surface, x, y float64, c composite.Op) {
// 	s.Ctx.Set("globalCompositeOperation", c)
// 	s.Blit(source, x, y)
// 	s.Ctx.Set("globalCompositeOperation", composite.SourceOver)
// }

// // BlitAreaComp is like BlitArea but uses the given composite operation.
// func (s *surface) BlitAreaComp(source Surface, area geo.Rect, x, y float64, c composite.Op) {
// 	s.ctx.Set("globalCompositeOperation", c)
// 	s.BlitArea(source, area, x, y)
// 	s.ctx.Set("globalCompositeOperation", composite.SourceOver)
// }

// func (s *surface) SubSurface(area geo.Rect) Surface {
// 	return &subsurface{
// 		area:   area,
// 		parent: s,
// 	}
// }
