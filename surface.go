package gogame

import "github.com/gopherjs/gopherjs/js"

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
	s.canvas.Call("drawImage", source.GetCanvas(), x, y)
}

// BlitArea draws the given portion of the source surface defined by the Rect to this
// one at the given position
func (s *surface) BlitArea(source Surface, area *Rect, x, y float64) {
	s.canvas.Call("drawImage", source.GetCanvas(), area.X, area.Y, area.W, area.H, x, y, area.W, area.H)
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

// DrawRect draws a rectangle on the surface. The thickness of the outer edge is determined
// by the width parameter, if it is <= zero then the rectangle will be filled.
func (s *surface) DrawRect(r *Rect, c Color, width float64) {
	s.ctx.Call("save")
	f := ""
	if width <= 0 {
		f = "fillRect"
		s.ctx.Set("fillStyle", c.String())
	} else {
		f = "strokeRect"
		s.ctx.Set("strokeStyle", c.String())
		s.ctx.Set("lineWidth", width)
	}
	s.ctx.Call(f, r.X, r.Y, r.W, r.H)
	s.ctx.Call("restore")
}

// func (s *surface) DrawCircle(pos Point, radius float64, c Color, width float64)
// func (s *surface) DrawElipse(r *Rect, c Color, width float64)
// func (s *surface) DrawArc(r *Rect, startAngle, stopAngle float64, c Color, width float64)
// func (s *surface) DrawLine(startPos, endPos Point, c Color, width float64)
// func (s *surface) DrawLines(pointList []Point, closed bool, c Color, width float64)
