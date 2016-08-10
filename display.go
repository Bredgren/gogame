package gogame

import "github.com/gopherjs/gopherjs/js"

var _ Surface = &Display{}

// Display is the main canvas in the web page and implements the Surface interface.
type Display struct {
	surface
}

func newDisplay(canvas *js.Object) *Display {
	d := &Display{}
	d.canvas = canvas
	d.ctx = canvas.Call("getContext", "2d")
	return d
}

// SetWidth returns the width of the canvas in pixels
func (c *Display) SetWidth(width float64) {
	c.canvas.Set("width", width)
}

// SetHeight returns the height of the canvas in pixels
func (c *Display) SetHeight(height float64) {
	c.canvas.Set("height", height)
}
