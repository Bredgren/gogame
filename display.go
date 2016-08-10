package gogame

import "github.com/gopherjs/gopherjs/js"

var _ Surface = &Display{}

// Display is the main canvas in the web page and implements the Surface interface.
type Display struct {
	surface
	canvas *js.Object
	ctx    *js.Object
}

func newDisplay(canvas *js.Object) *Display {
	return &Display{
		canvas: canvas,
		ctx:    canvas.Call("getContext", "2d"),
	}
}

// SetWidth returns the width of the canvas in pixels
func (c *Display) SetWidth(width float64) {
	c.canvas.Set("width", width)
}

// SetHeight returns the height of the canvas in pixels
func (c *Display) SetHeight(height float64) {
	c.canvas.Set("height", height)
}
