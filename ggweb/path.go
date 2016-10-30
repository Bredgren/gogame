package ggweb

import (
	"math"

	"github.com/Bredgren/gogame/geo"
	"github.com/gopherjs/gopherjs/js"
)

// Path is equivalent to the Path2D object in js. It holds a reusable sequence of instructions
// for drawing a path. One should use NewPath or NewPathSVG to create it.
type Path struct {
	obj *js.Object
}

// NewPath returns an empty path object.
func NewPath() *Path {
	return &Path{
		obj: js.Global.Get("Path2D").New(),
	}
}

// NewPathSVG returns a new path from SVG data.
func NewPathSVG(data string) *Path {
	return &Path{
		obj: js.Global.Get("Path2D").New(data),
	}
}

// Copy returns a copy of this path.
func (p *Path) Copy() *Path {
	return &Path{
		obj: js.Global.Get("Path2D").New(p.obj),
	}
}

// AddPath appends the given path to p.
func (p *Path) AddPath(p2 *Path) {
	p.obj.Call("addPath", p2.obj)
}

// func (p *Path) AddPathTransform(p2 *Path, t Transform)

// MoveTo moves to the point (x, y) without drawing anything.
func (p *Path) MoveTo(x, y float64) {
	p.obj.Call("moveTo", x, y)
}

// LineTo moves to the point (x, y) and draws a line between the previous position and
// the new position.
func (p *Path) LineTo(x, y float64) {
	p.obj.Call("lineTo", x, y)
}

// Rect adds a rectangle to the path.
func (p *Path) Rect(r geo.Rect) {
	p.obj.Call("rect", r.X, r.Y, r.W, r.H)
}

// Arc draws a circular arc with center (x, y) radius r. The angle are specified counterclockwise
// relative to the +x axis.
func (p *Path) Arc(x, y, r, startRadians, endRadians float64, counterclockwise bool) {
	p.obj.Call("arc", x, y, r, 2*math.Pi-startRadians, 2*math.Pi-endRadians, counterclockwise)
}

// ArcTo draws an arc tangent to the line between (x1, y1), (x2, y2) and connected to the
// previous point by a straight line.
func (p *Path) ArcTo(x1, y1, x2, y2, r float64) {
	p.obj.Call("arcTo", x1, y1, x2, y2, r)
}

// Ellipse draws an ellipse with the given rectangle. The ellipse will be rotated counterclockwise
// by rotateRadians. Other parameters are the same as Arc.
func (p *Path) Ellipse(r geo.Rect, rotateRadians, startRadians, endRadians float64, counterclockwise bool) {
	if p.obj.Get("ellipse") == js.Undefined {
		Warn("ellipse not supported")
		return
	}
	p.obj.Call("ellipse", r.CenterX(), r.CenterY(), r.W/2, r.H/2, 2*math.Pi-rotateRadians,
		2*math.Pi-startRadians, 2*math.Pi-endRadians, counterclockwise)
}

// QuadraticCurveTo adds a quadratic curve to the path from the current point to (x, y)
// with control point (cpx, cpy).
func (p *Path) QuadraticCurveTo(cpx, cpy, x, y float64) {
	p.obj.Call("quadraticCurveTo", cpx, cpy, x, y)
}

// BezierCurveTo adds a bezier curve to the path from the current point to (x, y)
// with control points (cp1x, cp1y) for the start and (cp2x, cp2y) for the end.
func (p *Path) BezierCurveTo(cp1x, cp1y, cp2x, cp2y, x, y float64) {
	p.obj.Call("bezierCurveTo", cp1x, cp1y, cp2x, cp2y, x, y)
}

// Close draws a line to the start of the last continous line.
func (p *Path) Close() {
	p.obj.Call("closePath")
}
