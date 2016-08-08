package gogame

import "github.com/gopherjs/gopherjs/js"

var _ Surface = &Canvas{}

// Canvas is the canvas in the web page. It implements the Surface interface.
type Canvas struct {
	canvas *js.Object
	ctx    *js.Object
}

func newCanvas(canvas *js.Object) *Canvas {
	return &Canvas{
		canvas: canvas,
		ctx:    canvas.Call("getContext", "2d"),
	}
}

// Blit draws the given surface to the Canvas at the given position
func (c *Canvas) Blit(source Surface, x, y int) {
	return
}

// Fill fills the whole canvas with one color
func (c *Canvas) Fill(color Color) {
	c.ctx.Set("fillStyle", color.String())
	c.ctx.Call("fillRect", 0, 0, c.canvas.Get("width"), c.canvas.Get("height"))
}

// SetWidth returns the width of the canvas in pixels
func (c *Canvas) SetWidth(width int) {
	c.canvas.Set("width", width)
}

// SetHeight returns the height of the canvas in pixels
func (c *Canvas) SetHeight(height int) {
	c.canvas.Set("height", height)
}

// Width returns the width of the canvas in pixels
func (c *Canvas) Width() int {
	return c.canvas.Get("width").Int()
}

// Height returns the height of the canvas in pixels
func (c *Canvas) Height() int {
	return c.canvas.Get("height").Int()
}
