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

// Copy returns a new surface equivalent to s in size and content. Context state is not copied.
func (s *Surface) Copy() *Surface {
	r := s.Rect()
	copy := NewSurface(int(r.W), int(r.H))
	copy.Blit(s, 0, 0)
	return copy
}

// Rect returns the rectangle that covers the whole surface.
func (s *Surface) Rect() geo.Rect {
	return geo.Rect{
		W: s.Canvas.Get("width").Float(),
		H: s.Canvas.Get("height").Float(),
	}
}

// SetSize resizes the surface. Resizing the surface clears it's contents and context state.
func (s *Surface) SetSize(w, h int) {
	s.Canvas.Set("width", w)
	s.Canvas.Set("height", h)
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

// SetLineCap sets the style for line end.
func (s *Surface) SetLineCap(cap LineCap) {
	s.Ctx.Set("lineCap", string(cap))
}

// SetLineJoin sets the style for line corners.
func (s *Surface) SetLineJoin(join LineJoin) {
	s.Ctx.Set("lineJoin", string(join))
}

// SetLineWidth sets the width for lines.
func (s *Surface) SetLineWidth(width float64) {
	s.Ctx.Set("lineWidth", width)
}

// SetLineMiterLimit sets the maximum miter length.
func (s *Surface) SetLineMiterLimit(miter float64) {
	s.Ctx.Set("miterLimit", miter)
}

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

func (s *Surface) DrawLine(x1, y1, x2, y2 float64) {
	s.Ctx.Call("beginPath")
	s.Ctx.Call("moveTo", math.Floor(x1), math.Floor(y1))
	s.Ctx.Call("lineTo", math.Floor(x2), math.Floor(y2))
	s.Ctx.Call("stroke")
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

// Translate moves the orgin by the distances given.
func (s *Surface) Translate(x, y float64) {
	s.Ctx.Call("translate", math.Floor(x), math.Floor(y))
}

// Rotate rotates the surface conterclockwise around the current origin.
func (s *Surface) Rotate(radians float64) {
	s.Ctx.Call("rotate", 2*math.Pi-radians)
}

// Scale changes the scale of the surface. 1.0 keeps the current size, smaller values
// shrink and larger grow.
func (s *Surface) Scale(x, y float64) {
	s.Ctx.Call("scale", x, y)
}

// Transform multiplies the current transformation matrix by the one described by the
// parameters:
//  [ a c e ]
//  [ b d f ]
//  [ 0 0 1 ]
func (s *Surface) Transform(a, b, c, d, e, f float64) {
	s.Ctx.Call("transform", a, b, c, d, e, f)
}

// SetTransform resets the transformation matrix then applies the one given.
func (s *Surface) SetTransform(a, b, c, d, e, f float64) {
	s.Ctx.Call("setTransform", a, b, c, d, e, f)
}

// ResetTransform resets the transformation to the identy matrix.
func (s *Surface) ResetTransform() {
	s.Ctx.Call("resetTransform")
}

// SetFont sets the font style.
func (s *Surface) SetFont(f *Font) {
	s.Ctx.Set("font", f.String())
}

// SetTextAlign sets the horizontal alignment of text.
func (s *Surface) SetTextAlign(a TextAlign) {
	s.Ctx.Set("textAlign", string(a))
}

// SetTextBaseline sets the vertical alignment of text.
func (s *Surface) SetTextBaseline(b TextBaseline) {
	s.Ctx.Set("textBaseline", string(b))
}

// DrawText draws the text to the surface at (x, y).
func (s *Surface) DrawText(t DrawType, text string, x, y float64) {
	s.Ctx.Call(string(t)+"Text", text, math.Floor(x), math.Floor(y))
}

// TextWidth returns the width in pixels that the given text will occupy.
func (s *Surface) TextWidth(text string) float64 {
	return s.Ctx.Call("measureText", text).Get("width").Float()
}

// PixelData returns a flat array of colors for each pixel within the given area.
func (s *Surface) PixelData(area geo.Rect) []color.RGBA {
	x, y, w, h := math.Floor(area.X), math.Floor(area.Y), math.Floor(area.W), math.Floor(area.H)
	imgData := s.Ctx.Call("getImageData", x, y, w, h).Get("data")
	data := make([]color.RGBA, int(w*h))
	for i := 0; i < len(data); i++ {
		data[i] = color.RGBA{
			R: uint8(imgData.Index(i * 4).Int()),
			G: uint8(imgData.Index(i*4 + 1).Int()),
			B: uint8(imgData.Index(i*4 + 2).Int()),
			A: uint8(imgData.Index(i*4 + 3).Int()),
		}
	}
	return data
}

// SetPixelData sets the pixels within the given area to the colors in data. The number
// of elemts of data should match area.Area().
func (s *Surface) SetPixelData(data []color.RGBA, area geo.Rect) {
	x, y, w, h := math.Floor(area.X), math.Floor(area.Y), math.Floor(area.W), math.Floor(area.H)
	imgData := s.Ctx.Call("getImageData", x, y, w, h)
	pxData := imgData.Get("data")
	for i := 0; i < len(data); i++ {
		pxData.SetIndex(i*4, data[i].R)
		pxData.SetIndex(i*4+1, data[i].G)
		pxData.SetIndex(i*4+2, data[i].B)
		pxData.SetIndex(i*4+3, data[i].A)
	}
	s.Ctx.Call("putImageData", imgData, x, y)
}

// Alpha returns the global alpha value for the surface. 0.0 is transparent and 1.0 is
// opaque.
func (s *Surface) Alpha() float64 {
	return s.Ctx.Get("globalAlpha").Float()
}

// SetAlpha sets the global alpha value for the surface. 0.0 is transparent and 1.0 is
// opaque.
func (s *Surface) SetAlpha(a float64) {
	s.Ctx.Set("globalAlpha", a)
}

// SetCompositeOp sets the composite operation to use. Default is SourceOver.
func (s *Surface) SetCompositeOp(op CompositeOp) {
	s.Ctx.Set("globalCompositeOperation", string(op))
}

// SetCursor sets the appearence of the cursor when it is over this Display.
func (s *Surface) SetCursor(c Cursor) {
	s.Canvas.Get("style").Set("cursor", c)
}

// Cursor returns the current appearence of the cursor when it is over the Display.
func (s *Surface) Cursor() Cursor {
	c := s.Canvas.Get("style").Get("cursor").String()
	if c == "" {
		return CursorDefault
	}
	return Cursor(c)
}
