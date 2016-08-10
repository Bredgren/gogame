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

// SetMode initalizes the Display.
func (c *Display) SetMode(width, height int) {
	c.canvas.Set("width", width)
	c.canvas.Set("height", height)
}