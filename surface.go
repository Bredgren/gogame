package gogame

import (
	"fmt"
	"math"

	"github.com/gopherjs/gopherjs/js"
)

// Surface represents an image or drawable sureface.
type Surface interface {
	GetCanvas() *js.Object
	Blit(source Surface, x, y float64)
	BlitArea(source Surface, area *Rect, x, y float64)
	Fill(*FillStyle)
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
	// Scale() (x, y float64)
	// SetScale(x, y float64)
	// Rotation() float64
	// SetRotation(radians float64)
	// Transformed() Surface
	Scaled(x, y float64) Surface
	Rotated(radians float64) Surface
	GetRect() Rect
	DrawRect(*Rect, Styler)
	DrawCircle(posX, posY, radius float64, s Styler)
	DrawEllipse(*Rect, Styler)
	DrawArc(r *Rect, startRadians, stopRadians float64, s Styler)
	DrawLine(startX, startY, endX, endY float64, s Styler)
	DrawLines(pointList [][2]float64, s Styler)
	//DrawQuadraticCurve(startX, startY, endX, endY, cpX, cpY, float64, s Styler)
	//DrawQuadraticCurves(points [][2]float64, cPoints [][2]float64, s Styler)
	//DrawBezierCurve(startX, startY, endX, endY, cpStartX, cpStartY, cpEndX, cpEndY float64, s Styler)
	//DrawBezierCurves(points [][2]float64, cPoints [][2]float64, s Styler)
	DrawText(text string, x, y float64, s *FontStyler)
}

var _ Surface = &surface{}

type surface struct {
	canvas   *js.Object
	ctx      *js.Object
	rotation float64
	scaleX   float64
	scaleY   float64
	// if needed we can cache the Rect for GetRect, which would only need to change when
	// scaleX/Y or rotation change.
	// rect *Rect
}

// NewSurface creates a new Surface.
func NewSurface(width, height int) Surface {
	canvas := jq("<canvas>").Get(0)
	canvas.Set("width", width)
	canvas.Set("height", height)
	return &surface{
		canvas:   canvas,
		ctx:      canvas.Call("getContext", "2d"),
		rotation: 0,
		scaleX:   1,
		scaleY:   1,
	}
}

// GetCanvas returns the surface as an HTML canvas.
func (s *surface) GetCanvas() *js.Object {
	return s.canvas
}

// Blit draws the given surface to this one at the given position. Source's top-left corner
// (according to GetRect) fill be drawn at (x, y).
func (s *surface) Blit(source Surface, x, y float64) {
	// 	s.ctx.Call("save")

	// 	sx, sy := source.Scale()
	// 	s.ctx.Call("scale", sx, sy)

	// 	rect := source.GetRect()
	// 	rect.X = x
	// 	rect.Y = y
	// 	cx, cy := math.Floor(rect.CenterX()), math.Floor(rect.CenterY())
	// 	s.ctx.Call("translate", cx, cy)
	// 	s.ctx.Call("rotate", -source.Rotation())
	// 	s.ctx.Call("translate", -cx, -cy)

	s.ctx.Call("drawImage", source.GetCanvas(), math.Floor(x), math.Floor(y))

	// s.ctx.Call("restore")
}

// BlitArea draws the given portion of the source surface defined by the Rect to this
// one with its top-left corner (according to GetRect) at the given position.
func (s *surface) BlitArea(source Surface, area *Rect, x, y float64) {
	// s.ctx.Call("save")

	// sx, sy := source.Scale()
	// s.ctx.Call("scale", sx, sy)

	// rect := source.GetRect()
	// rect.X = x
	// rect.Y = y
	// cx, cy := math.Floor(rect.CenterX()), math.Floor(rect.CenterY())
	// s.ctx.Call("translate", cx, cy)
	// s.ctx.Call("rotate", -source.Rotation())
	// s.ctx.Call("translate", -cx, -cy)

	s.ctx.Call("drawImage", source.GetCanvas(), math.Floor(area.X), math.Floor(area.Y),
		math.Floor(area.W), math.Floor(area.H), math.Floor(x), math.Floor(y), math.Floor(area.W),
		math.Floor(area.H))

	// s.ctx.Call("restore")
}

// Fill fills the whole surface with the given style.
func (s *surface) Fill(style *FillStyle) {
	s.ctx.Call("save")
	style.Style(s.ctx)
	s.ctx.Call("fillRect", 0, 0, s.canvas.Get("width"), s.canvas.Get("height"))
	s.ctx.Call("restore")
}

// Width returns the unrotated, unscaled width of the surface in pixels. To get the width
// after scaling and rotating use GetRect.
func (s *surface) Width() int {
	return s.canvas.Get("width").Int()
}

// Height returns the unrotated, unscaled height of the surface in pixels. To get the height
// after scaling and rotating use GetRect.
func (s *surface) Height() int {
	return s.canvas.Get("height").Int()
}

// // Scale returns the x and y scales of the this surface.
// func (s *surface) Scale() (x, y float64) {
// 	return s.scaleX, s.scaleY
// }

// // SetScale sets the x and y scales of this surface.
// func (s *surface) SetScale(x, y float64) {
// 	s.scaleX = x
// 	s.scaleY = y
// }

// Scaled returns a new Surface that is equivalent to this one scaled by the given amount.
func (s *surface) Scaled(x, y float64) Surface {
	newS := NewSurface(int(float64(s.Width())*x), int(float64(s.Height())*y))
	ctx := newS.(*surface).ctx
	ctx.Call("save")
	ctx.Call("scale", x, y)
	ctx.Call("drawImage", s.canvas, 0, 0)
	ctx.Call("restore")
	return newS
}

// // Rotation returns the rotation (counter-clockwise) of this surface in radians.
// func (s *surface) Rotation() float64 {
// 	return s.rotation
// }

// // SetRotation sets the rotation (counter-clockwise) of this surface in radians.
// func (s *surface) SetRotation(radians float64) {
// 	s.rotation = radians
// }

// Rotated returns a new Surface that is equivalent to this one but rotated counter-clockwise
// by the given amount.
func (s *surface) Rotated(radians float64) Surface {
	newW, newH := s.getRotatedSize(radians)
	newS := NewSurface(newW, newH)
	ctx := newS.(*surface).ctx
	ctx.Call("save")
	cx, cy := newW/2, newH/2
	s.ctx.Call("translate", cx, cy)
	s.ctx.Call("rotate", -radians)
	s.ctx.Call("translate", -cx, -cy)
	ctx.Call("drawImage", s.canvas, 0, 0)
	ctx.Call("restore")
	return newS
}

// // Transformed returns a new Surface that is equivalent to this one but with the scale and
// // rotation trnsforms applied.
// func (s *surface) Transformed() Surface {
// 	rect := s.GetRect()
// 	newS := NewSurface(int(rect.W), int(rect.H))
// 	newS.Blit(s, 0, 0)
// 	return newS
// }

func (s *surface) getRotatedSize(radians float64) (w, h int) {
	width, height := float64(s.Width()), float64(s.Height())
	cx, cy := width/2, height/2
	cos, sin := math.Cos(s.rotation), math.Sin(s.rotation)
	corners := [][2]float64{
		{0.0, 0.0},
		{width, 0.0},
		{0.0, height},
		{width, height},
	}
	for _, corner := range corners {
		nx := cx + (corner[0]-cx)*cos + (corner[1]-cy)*sin
		ny := cy - (corner[0]-cx)*sin + (corner[1]-cy)*cos
		corner[0] = nx
		corner[1] = ny
	}
	minX, maxX, minY, maxY := corners[0][0], corners[0][0], corners[0][1], corners[0][1]
	for _, corner := range corners[1:] {
		if corner[0] < minX {
			minX = corner[0]
		}
		if corner[0] > maxX {
			maxX = corner[0]
		}
		if corner[1] < minY {
			minY = corner[1]
		}
		if corner[1] > maxY {
			maxY = corner[1]
		}
	}
	return int(maxX - minX), int(maxY - minY)
}

// GetRect returns the bouding rectangle for the surface, taking into acount rotation and scale.
func (s *surface) GetRect() Rect {
	// w := float64(s.Width()) * s.scaleX
	// h := float64(s.Height()) * s.scaleY
	// cx, cy := w/2, h/2
	// cos, sin := math.Cos(s.rotation), math.Sin(s.rotation)
	// corners := [][2]float64{
	// 	{0.0, 0.0},
	// 	{w, 0.0},
	// 	{0.0, h},
	// 	{w, h},
	// }
	// for _, corner := range corners {
	// 	nx := cx + (corner[0]-cx)*cos + (corner[1]-cy)*sin
	// 	ny := cy - (corner[0]-cx)*sin + (corner[1]-cy)*cos
	// 	corner[0] = nx
	// 	corner[1] = ny
	// }
	// minX, maxX, minY, maxY := corners[0][0], corners[0][0], corners[0][1], corners[0][1]
	// for _, corner := range corners[1:] {
	// 	if corner[0] < minX {
	// 		minX = corner[0]
	// 	}
	// 	if corner[0] > maxX {
	// 		maxX = corner[0]
	// 	}
	// 	if corner[1] < minY {
	// 		minY = corner[1]
	// 	}
	// 	if corner[1] > maxY {
	// 		maxY = corner[1]
	// 	}
	// }
	// return Rect{X: 0, Y: 0, W: maxX - minX, H: maxY - minY}
	return Rect{X: 0, Y: 0, W: float64(s.Width()), H: float64(s.Height())}
}

// DrawRect draws a rectangle on the surface.
func (s *surface) DrawRect(r *Rect, style Styler) {
	s.ctx.Call("save")
	style.Style(s.ctx)
	s.ctx.Call("translate", math.Floor(r.X), math.Floor(r.Y))
	s.ctx.Call(fmt.Sprintf("%sRect", style.DrawType()), 0, 0, math.Floor(r.W), math.Floor(r.H))
	s.ctx.Call("restore")
}

// DrawCircle draws a circle on the surface.
func (s *surface) DrawCircle(posX, posY, radius float64, style Styler) {
	s.ctx.Call("save")
	style.Style(s.ctx)
	s.ctx.Call("translate", math.Floor(posX), math.Floor(posY))
	s.ctx.Call("beginPath")
	s.ctx.Call("arc", 0, 0, radius, 0, 2*math.Pi)
	s.ctx.Call(string(style.DrawType()))
	s.ctx.Call("restore")
}

// DrawEllipse draws an ellipse on the canvas within the given Rect.
func (s *surface) DrawEllipse(r *Rect, style Styler) {
	s.ctx.Call("save")
	style.Style(s.ctx)
	s.ctx.Call("translate", math.Floor(r.X), math.Floor(r.Y))
	s.ctx.Call("beginPath")
	s.ctx.Call("ellipse", math.Floor(r.Width()/2), math.Floor(r.Height()/2), math.Floor(r.Width()/2),
		math.Floor(r.Height()/2), 0, 0, 2*math.Pi)
	s.ctx.Call(string(style.DrawType()))
	s.ctx.Call("restore")
}

// DrawArc draws an arc on the canvas within the given Rect between the given angles.
// Angles are counter-clockwise.
func (s *surface) DrawArc(r *Rect, startRadians, stopRadians float64, style Styler) {
	s.ctx.Call("save")
	style.Style(s.ctx)
	s.ctx.Call("translate", math.Floor(r.X), math.Floor(r.Y))
	s.ctx.Call("beginPath")
	s.ctx.Call("ellipse", math.Floor(r.Width()/2), math.Floor(r.Height()/2), math.Floor(r.Width()/2),
		math.Floor(r.Height()/2), 0, 2*math.Pi-startRadians, 2*math.Pi-stopRadians, true)
	s.ctx.Call(string(style.DrawType()))
	s.ctx.Call("restore")
}

// DrawLine draws a line on the canvas.
func (s *surface) DrawLine(startX, startY, endX, endY float64, style Styler) {
	s.ctx.Call("save")
	style.Style(s.ctx)
	s.ctx.Call("translate", startX, startY)
	s.ctx.Call("beginPath")
	s.ctx.Call("moveTo", 0, 0)
	s.ctx.Call("lineTo", endX-startX, endY-startY)
	s.ctx.Call(string(style.DrawType()))
	s.ctx.Call("restore")
}

// DrawLines draws multiple connectd lines to the surface.
func (s *surface) DrawLines(pointList [][2]float64, style Styler) {
	s.ctx.Call("save")
	style.Style(s.ctx)
	s.ctx.Call("beginPath")
	s.ctx.Call("moveTo", pointList[0][0], pointList[0][1])
	for _, p := range pointList[1:] {
		s.ctx.Call("lineTo", p[0], p[1])
	}
	s.ctx.Call(string(style.DrawType()))
	s.ctx.Call("restore")
}

// DrawText draws the given text to the surface.
func (s *surface) DrawText(text string, x, y float64, style *FontStyler) {
	s.ctx.Call("save")
	style.Style(s.ctx)
	s.ctx.Call("translate", math.Floor(x), math.Floor(y))
	s.ctx.Call(fmt.Sprintf("%sText", style.DrawType()), text, 0, 0)
	s.ctx.Call("restore")
}
