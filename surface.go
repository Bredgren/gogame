package gogame

import (
	"fmt"
	"math"

	"github.com/gopherjs/gopherjs/js"
)

// Surface represents an image or drawable sureface
type Surface interface {
	GetCanvas() *js.Object
	Blit(source Surface, x, y float64)
	BlitArea(source Surface, area *Rect, x, y float64)
	Fill(Color)
	Width() int
	Height() int
	//Copy() Surface
	//Scroll(dx, dy int)
	//GetAt(x, y int) Color
	//SetAt(x, y int, Color)
	//SetClip(Rect)
	//GetClip() Rect
	//GetSubsurface(Rect) Surface
	//GetParent() Surface
	//GetRect() Rect
}

var _ Surface = &surface{}

type surface struct {
	canvas *js.Object
	ctx    *js.Object
}

// NewSurface creates a new Surface
func NewSurface(width, height float64) Surface {
	canvas := jq("<canvas>").Get()
	return &surface{
		canvas: canvas,
		ctx:    canvas.Call("getContext", "2d"),
	}
}

// GetCanvas returns the surface as an HTML canvas
func (s *surface) GetCanvas() *js.Object {
	return s.canvas
}

// Blit draws the given surface to this one at the given position
func (s *surface) Blit(source Surface, x, y float64) {
	s.canvas.Call("drawImage", source.GetCanvas(), math.Floor(x), math.Floor(y))
}

// BlitArea draws the given portion of the source surface defined by the Rect to this
// one at the given position
func (s *surface) BlitArea(source Surface, area *Rect, x, y float64) {
	s.canvas.Call("drawImage", source.GetCanvas(), math.Floor(area.X), math.Floor(area.Y),
		math.Floor(area.W), math.Floor(area.H), math.Floor(x), math.Floor(y), math.Floor(area.W),
		math.Floor(area.H))
}

// Fill fills the whole surface with one color
func (s *surface) Fill(color Color) {
	s.ctx.Set("fillStyle", color.String())
	s.ctx.Call("fillRect", 0, 0, s.canvas.Get("width"), s.canvas.Get("height"))
}

// Width returns the width of the surface in pixels
func (s *surface) Width() int {
	return s.canvas.Get("width").Int()
}

// Height returns the height of the surface in pixels
func (s *surface) Height() int {
	return s.canvas.Get("height").Int()
}

// DrawRect draws a rectangle on the surface
func (s *surface) DrawRect(r *Rect, style Styler) {
	s.ctx.Call("save")
	style.Style(s.ctx)
	s.ctx.Call(fmt.Sprintf("%sRect", style.DrawType()), math.Floor(r.X), math.Floor(r.Y),
		math.Floor(r.W), math.Floor(r.H))
	s.ctx.Call("restore")
}

// func (s *surface) DrawCircle(posX, posY, radius, width float64, c Color)
// func (s *surface) DrawElipse(r *Rect, width float64, c Color)
// func (s *surface) DrawArc(r *Rect, startAngle, stopAngle, width float64, c Color)
// func (s *surface) DrawLine(startX startY, endX, endY, width float64, c Color)
// func (s *surface) DrawLines(pointList [][2]float64, closed bool, width float64, c Color)// func (s *surface) DrawQuadraticCurve(startX, startY, endX, endY, cpX, cpY, float64, style Styler)
// func (s *surface) DrawQuadraticCurves(points [][2]float64, cPoints [][2]float64, style Styler)
// func (s *surface) DrawBezierCurve(startX, startY, endX, endY, cpStartX, cpStartY, cpEndX, cpEndY float64, style Styler)
// func (s *surface) DrawBezierCurves(points [][2]float64, cPoints [][2]float64, style Styler)