package gogame

import (
	"fmt"
	"math"
	"strings"

	"github.com/Bredgren/gogame/composite"
	"github.com/Bredgren/gogame/geo"
	"github.com/gopherjs/gopherjs/js"
)

// Surface represents an image or drawable surface.
type Surface interface {
	// Canvas returns the surface as an HTML canvas.
	Canvas() *js.Object
	// Blit draws the given surface to this one at the given position. Source's top-left corner
	// (according to Surface.Rect() fill be drawn at (x, y).
	Blit(source Surface, x, y float64)
	// BlitArea draws the given portion of the source surface defined by the Rect to this
	// one with its top-left corner (according to Surface.Rect()) at the given position.
	BlitArea(source Surface, area geo.Rect, x, y float64)
	// BlitComp is like Blit but takes a composite operation to use.
	BlitComp(source Surface, x, y float64, c composite.Op)
	// BlitAreaComp is like BlitArea but takes a composite operation to use.
	BlitAreaComp(source Surface, area geo.Rect, x, y float64, c composite.Op)
	// Fill fills the whole surface with the given style.
	Fill(*FillStyle)
	// Width returns the unrotated, unscaled width of the surface in pixels. To get the width
	// after scaling and rotating use Surface.Rect.
	Width() int
	// Height returns the unrotated, unscaled height of the surface in pixels. To get the height
	// after scaling and rotating use Surface.Rect.
	Height() int
	// Copy returns a new Surface that is identical to this one.
	Copy() Surface
	// GetAt returns the color of the pixel at (x, y).
	GetAt(x, y int) Color
	// SetAt sets the color of the pixel at (x, y).
	SetAt(x, y int, c Color)
	// SetClip defines a rectangular region of the surface where only the enclosed pixels can
	// be affected by draw operations.
	SetClip(geo.Rect)
	// SetClipPath defines an arbitrary polygon where only the enclosed pixels can be affected
	// by draw operations. The pointList is a list of xy-coordinates.
	SetClipPath(pointList [][2]float64)
	// ClearClip resets the clipping region to the whole canvas.
	ClearClip()
	// Scaled returns a new Surface that is equivalent to this one scaled by the given amount.
	Scaled(x, y float64) Surface
	// Rotated returns a new Surface that is equivalent to this one but rotated counter-clockwise
	// by the given amount.
	Rotated(radians float64) Surface
	// Rect returns the bouding rectangle for the surface, taking into acount rotation and scale.
	Rect() geo.Rect
	// DrawRect draws a rectangle on the surface.
	DrawRect(geo.Rect, Styler)
	// DrawCircle draws a circle on the surface.
	DrawCircle(posX, posY, radius float64, s Styler)
	// DrawEllipse draws an ellipse on the canvas within the given Rect.
	DrawEllipse(geo.Rect, Styler)
	// DrawArc draws an arc on the canvas within the given Rect between the given angles.
	// Angles are counter-clockwise.
	DrawArc(r geo.Rect, startRadians, stopRadians float64, s Styler)
	// DrawLine draws a line on the surface.
	DrawLine(startX, startY, endX, endY float64, s Styler)
	// DrawLines draws multiple connected lines to the surface.
	DrawLines(points [][2]float64, s Styler)
	// DrawText draws the given text to the surface.
	DrawText(text string, x, y float64, font *Font, style *TextStyle)
	// DrawQuadraticCurve draws a quadratic curve to the surface from (startX, startY) to
	// (endX, endY) with the control point (cpX, cpY).
	DrawQuadraticCurve(startX, startY, endX, endY, cpX, cpY float64, s Styler)
	// DrawQuadraticCurves draws multiple connected quadratic curves to the surface. The list
	// of points alternaes between end points and control points, i.e. [startpoint, control1,
	// point2, control2, point3, etc.]. If the list has an even number of points then the last
	// point is ignored. If there are less than 3 points then nothing will drawn.
	DrawQuadraticCurves(points [][2]float64, s Styler)
	// DrawBezierCurve draws a cubic bezier curve to the surface from (startX, startY) to
	// (endX, endY) with the given respective control points..
	DrawBezierCurve(startX, startY, endX, endY, cpStartX, cpStartY, cpEndX, cpEndY float64, s Styler)
	// DrawBezierCurves draws multiple connected cubic quadratic curves to the surface. The list
	// of points alternaes between end points and both control points, i.e. [startpoint, control1,
	// control2, point2, control3, control4, point3, etc.]. If there are not enough points at
	// the end of the list then it will stop at the last possible point. If there are less than
	// 4 points then nothing will drawn.
	DrawBezierCurves(points [][2]float64, s Styler)
}

var _ Surface = &surface{}

type surface struct {
	canvas *js.Object
	ctx    *js.Object
}

// NewSurface creates a new Surface.
func NewSurface(width, height int) Surface {
	canvas := js.Global.Get("document").Call("createElement", "canvas")
	canvas.Set("width", width)
	canvas.Set("height", height)
	return &surface{
		canvas: canvas,
		ctx:    canvas.Call("getContext", "2d"),
	}
}

// NewSurfaceFromCanvas creates a new Surface from the given canvas. An error is returned
// if canvas is nil or undefined.
func NewSurfaceFromCanvas(canvas *js.Object) (Surface, error) {
	if canvas == nil || canvas == js.Undefined {
		return nil, fmt.Errorf("invalid canvas: %v", canvas)
	}
	return &surface{
		canvas: canvas,
		ctx:    canvas.Call("getContext", "2d"),
	}, nil
}

// NewSurfaceFromCanvasID creates a new Surface using the canvas with the given ID. An
// error is returned if no canvas was found.
func NewSurfaceFromCanvasID(canvasID string) (Surface, error) {
	if !strings.HasPrefix(canvasID, "#") {
		canvasID = "#" + canvasID
	}
	canvas := js.Global.Get("document").Call("getElementById", canvasID)
	if canvas == js.Undefined {
		return nil, fmt.Errorf("no canvas found with ID '%s'", canvasID)
	}
	return &surface{
		canvas: canvas,
		ctx:    canvas.Call("getContext", "2d"),
	}, nil
}

func (s *surface) Canvas() *js.Object {
	// TODO handle nil surface
	return s.canvas
}

func (s *surface) Blit(source Surface, x, y float64) {
	s.ctx.Call("drawImage", source.Canvas(), math.Floor(x), math.Floor(y))
}

func (s *surface) BlitArea(source Surface, area geo.Rect, x, y float64) {
	s.ctx.Call("drawImage", source.Canvas(), math.Floor(area.X), math.Floor(area.Y),
		math.Floor(area.W), math.Floor(area.H), math.Floor(x), math.Floor(y), math.Floor(area.W),
		math.Floor(area.H))
}

func (s *surface) BlitComp(source Surface, x, y float64, c composite.Op) {
	s.ctx.Set("globalCompositeOperation", c)
	s.Blit(source, x, y)
	s.ctx.Set("globalCompositeOperation", composite.SourceOver)
}

func (s *surface) BlitAreaComp(source Surface, area geo.Rect, x, y float64, c composite.Op) {
	s.ctx.Set("globalCompositeOperation", c)
	s.BlitArea(source, area, x, y)
	s.ctx.Set("globalCompositeOperation", composite.SourceOver)
}

func (s *surface) Fill(style *FillStyle) {
	s.ctx.Call("save")
	style.Style(s.ctx)
	s.ctx.Call("fillRect", 0, 0, s.canvas.Get("width"), s.canvas.Get("height"))
	s.ctx.Call("restore")
}

func (s *surface) Width() int {
	return s.canvas.Get("width").Int()
}

func (s *surface) Height() int {
	return s.canvas.Get("height").Int()
}

func (s *surface) Copy() Surface {
	copy := NewSurface(s.Width(), s.Height())
	copy.Blit(s, 0, 0)
	return copy
}

func (s *surface) GetAt(x, y int) Color {
	data := s.ctx.Call("getImageData", x, y, 1, 1).Get("data")
	return Color{R: data.Index(0).Float(), G: data.Index(1).Float(), B: data.Index(2).Float(),
		A: data.Index(3).Float()}
}

func (s *surface) SetAt(x, y int, c Color) {
	imgData := s.ctx.Call("getImageData", x, y, 1, 1)
	data := imgData.Get("data")
	data.SetIndex(0, clampToInt(255*c.R, 0, 255))
	data.SetIndex(1, clampToInt(255*c.G, 0, 255))
	data.SetIndex(2, clampToInt(255*c.B, 0, 255))
	data.SetIndex(3, clampToInt(255*c.A, 0, 255))
	s.ctx.Call("putImageData", imgData, x, y)
}

func (s *surface) SetClip(r geo.Rect) {
	s.SetClipPath([][2]float64{
		{r.X, r.Y}, {r.X + r.W, r.Y}, {r.X + r.W, r.Y + r.H}, {r.X, r.Y + r.H},
	})
}

func (s *surface) SetClipPath(pointList [][2]float64) {
	s.ctx.Call("save")
	s.ctx.Call("beginPath")
	s.ctx.Call("moveTo", pointList[0][0], pointList[0][1])
	for _, p := range pointList[1:] {
		s.ctx.Call("lineTo", p[0], p[1])
	}
	s.ctx.Call("clip")
}

func (s *surface) ClearClip() {
	s.ctx.Call("restore")
}

func (s *surface) Scaled(x, y float64) Surface {
	newS := NewSurface(int(float64(s.Width())*x), int(float64(s.Height())*y))
	ctx := newS.(*surface).ctx
	ctx.Call("save")
	ctx.Call("scale", x, y)
	ctx.Call("drawImage", s.canvas, 0, 0)
	ctx.Call("restore")
	return newS
}

func (s *surface) Rotated(radians float64) Surface {
	newW, newH := s.getRotatedSize(radians)
	newS := NewSurface(newW, newH)
	ctx := newS.(*surface).ctx
	ctx.Call("save")
	cx, cy := newW/2, newH/2
	ctx.Call("translate", cx, cy)
	ctx.Call("rotate", -radians)
	ctx.Call("translate", -s.Width()/2, -s.Height()/2)
	ctx.Call("drawImage", s.canvas, 0, 0)
	ctx.Call("restore")
	return newS
}

func (s *surface) getRotatedSize(radians float64) (w, h int) {
	width, height := float64(s.Width()), float64(s.Height())
	cx, cy := width/2, height/2
	cos, sin := math.Cos(radians), math.Sin(radians)

	x1 := cx + (0-cx)*cos + (0-cy)*sin
	y1 := cy - (0-cx)*sin + (0-cy)*cos
	x2 := cx + (width-cx)*cos + (0-cy)*sin
	y2 := cy - (width-cx)*sin + (0-cy)*cos
	x3 := cx + (0-cx)*cos + (height-cy)*sin
	y3 := cy - (0-cx)*sin + (height-cy)*cos
	x4 := cx + (width-cx)*cos + (height-cy)*sin
	y4 := cy - (width-cx)*sin + (height-cy)*cos

	maxX := math.Max(x1, math.Max(x2, math.Max(x3, x4)))
	minX := math.Min(x1, math.Min(x2, math.Min(x3, x4)))
	maxY := math.Max(y1, math.Max(y2, math.Max(y3, y4)))
	minY := math.Min(y1, math.Min(y2, math.Min(y3, y4)))

	return int(maxX - minX), int(maxY - minY)
}

func (s *surface) Rect() geo.Rect {
	return geo.Rect{X: 0, Y: 0, W: float64(s.Width()), H: float64(s.Height())}
}

func (s *surface) DrawRect(r geo.Rect, style Styler) {
	if style == nil {
		style = &FillStyle{}
	}
	s.ctx.Call("save")
	style.Style(s.ctx)
	s.ctx.Call("translate", math.Floor(r.X), math.Floor(r.Y))
	s.ctx.Call(fmt.Sprintf("%sRect", style.DrawType()), 0, 0, math.Floor(r.W), math.Floor(r.H))
	s.ctx.Call("restore")
}

func (s *surface) DrawCircle(posX, posY, radius float64, style Styler) {
	if style == nil {
		style = &FillStyle{}
	}
	s.ctx.Call("save")
	style.Style(s.ctx)
	s.ctx.Call("translate", math.Floor(posX), math.Floor(posY))
	s.ctx.Call("beginPath")
	s.ctx.Call("arc", 0, 0, radius, 0, 2*math.Pi)
	s.ctx.Call(string(style.DrawType()))
	s.ctx.Call("restore")
}

func (s *surface) DrawEllipse(r geo.Rect, style Styler) {
	if style == nil {
		style = &FillStyle{}
	}
	s.ctx.Call("save")
	style.Style(s.ctx)
	s.ctx.Call("translate", math.Floor(r.X), math.Floor(r.Y))
	s.ctx.Call("beginPath")
	s.ctx.Call("ellipse", math.Floor(r.Width()/2), math.Floor(r.Height()/2), math.Floor(r.Width()/2),
		math.Floor(r.Height()/2), 0, 0, 2*math.Pi)
	s.ctx.Call(string(style.DrawType()))
	s.ctx.Call("restore")
}

func (s *surface) DrawArc(r geo.Rect, startRadians, stopRadians float64, style Styler) {
	if style == nil {
		style = &StrokeStyle{}
	}
	s.ctx.Call("save")
	style.Style(s.ctx)
	s.ctx.Call("translate", math.Floor(r.X), math.Floor(r.Y))
	s.ctx.Call("beginPath")
	s.ctx.Call("ellipse", math.Floor(r.Width()/2), math.Floor(r.Height()/2), math.Floor(r.Width()/2),
		math.Floor(r.Height()/2), 0, 2*math.Pi-startRadians, 2*math.Pi-stopRadians, true)
	s.ctx.Call(string(style.DrawType()))
	s.ctx.Call("restore")
}

func (s *surface) DrawLine(startX, startY, endX, endY float64, style Styler) {
	if style == nil {
		style = &StrokeStyle{}
	}
	s.ctx.Call("save")
	style.Style(s.ctx)
	s.ctx.Call("beginPath")
	// Not math.Flooring lines to preserve control with odd width lines.
	s.ctx.Call("moveTo", startX, startY)
	s.ctx.Call("lineTo", endX, endY)
	s.ctx.Call(string(style.DrawType()))
	s.ctx.Call("restore")
}

func (s *surface) DrawLines(points [][2]float64, style Styler) {
	if style == nil {
		style = &StrokeStyle{}
	}
	s.ctx.Call("save")
	style.Style(s.ctx)
	s.ctx.Call("beginPath")
	// Not math.Flooring lines to preserve control with odd width lines.
	s.ctx.Call("moveTo", points[0][0], points[0][1])
	for _, p := range points[1:] {
		s.ctx.Call("lineTo", p[0], p[1])
	}
	s.ctx.Call(string(style.DrawType()))
	s.ctx.Call("restore")
}

func (s *surface) DrawText(text string, x, y float64, font *Font, style *TextStyle) {
	s.ctx.Call("save")
	s.ctx.Set("font", font.String())
	style.Style(s.ctx)
	s.ctx.Call("translate", math.Floor(x), math.Floor(y))
	s.ctx.Call(fmt.Sprintf("%sText", style.DrawType()), text, 0, 0)
	s.ctx.Call("restore")
}

func (s *surface) DrawQuadraticCurve(startX, startY, endX, endY, cpX, cpY float64, style Styler) {
	if style == nil {
		style = &StrokeStyle{}
	}
	s.ctx.Call("save")
	style.Style(s.ctx)
	s.ctx.Call("beginPath")
	// Not math.Flooring lines to preserve control with odd width lines.
	s.ctx.Call("moveTo", startX, startY)
	s.ctx.Call("quadraticCurveTo", cpX, cpY, endX, endY)
	s.ctx.Call(string(style.DrawType()))
	s.ctx.Call("restore")
}

func (s *surface) DrawQuadraticCurves(points [][2]float64, style Styler) {
	if len(points) < 3 {
		return // Not enough points for event one curve.
	}
	if style == nil {
		style = &StrokeStyle{}
	}
	s.ctx.Call("save")
	style.Style(s.ctx)
	s.ctx.Call("beginPath")
	// Not math.Flooring lines to preserve control with odd width lines.
	s.ctx.Call("moveTo", points[0][0], points[0][1])
	for i := 1; i+1 < len(points); i += 2 {
		s.ctx.Call("quadraticCurveTo", points[i][0], points[i][1], points[i+1][0], points[i+1][1])
	}
	s.ctx.Call(string(style.DrawType()))
	s.ctx.Call("restore")
}

func (s *surface) DrawBezierCurve(startX, startY, endX, endY, cpStartX, cpStartY, cpEndX, cpEndY float64, style Styler) {
	if style == nil {
		style = &StrokeStyle{}
	}
	s.ctx.Call("save")
	style.Style(s.ctx)
	s.ctx.Call("beginPath")
	// Not math.Flooring lines to preserve control with odd width lines.
	s.ctx.Call("moveTo", startX, startY)
	s.ctx.Call("bezierCurveTo", cpStartX, cpStartY, cpEndX, cpEndY, endX, endY)
	s.ctx.Call(string(style.DrawType()))
	s.ctx.Call("restore")
}

func (s *surface) DrawBezierCurves(points [][2]float64, style Styler) {
	if len(points) < 4 {
		return // Not enough points for event one curve.
	}
	if style == nil {
		style = &StrokeStyle{}
	}
	s.ctx.Call("save")
	style.Style(s.ctx)
	s.ctx.Call("beginPath")
	// Not math.Flooring lines to preserve control with odd width lines.
	s.ctx.Call("moveTo", points[0][0], points[0][1])
	for i := 1; i+2 < len(points); i += 3 {
		s.ctx.Call("bezierCurveTo", points[i][0], points[i][1], points[i+1][0], points[i+1][1],
			points[i+2][0], points[i+2][1])
	}
	s.ctx.Call(string(style.DrawType()))
	s.ctx.Call("restore")
}
