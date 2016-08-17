package gogame

import "github.com/gopherjs/gopherjs/js"

var _ Surface = &Display{}

// Display is the main canvas in the web page and implements the Surface interface.
// It implements double buffering by having a back Surface. All draw operations will
// go to the back Surface until Flip is called. The behavior of a Display is undefined
// until SetMode is called.
type Display struct {
	surface
	frontSurface Surface
}

func newDisplay(canvas *js.Object) *Display {
	d := &Display{}
	d.frontSurface = NewSurfaceFromCanvas(canvas)
	return d
}

// SetMode initalizes the Display.
func (d *Display) SetMode(width, height int) {
	d.canvas = jq("<canvas>").Get(0)
	d.canvas.Set("width", width)
	d.canvas.Set("height", height)
	d.ctx = d.canvas.Call("getContext", "2d")
	d.frontSurface.GetCanvas().Set("width", width)
	d.frontSurface.GetCanvas().Set("height", height)
}

// Flip draws the back Surface onto the Display.
func (d *Display) Flip() {
	d.frontSurface.Blit(d, 0, 0)
}
