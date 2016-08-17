package gogame

import "github.com/gopherjs/gopherjs/js"

var _ Surface = &Display{}

// Display is a visible canvas in the web page and implements the Surface interface.
// It implements double buffering by having a back Surface. All draw operations will
// go to the back Surface until Flip is called. The behavior of a Display is undefined
// until SetMode is called.
type Display struct {
	surface
	frontSurface Surface
}

// NewDisplay creates a new display for the given canvas. An error is returned if canvas
// is nil or undefined.
func NewDisplay(canvas *js.Object) (*Display, error) {
	d := &Display{}
	var err error
	d.frontSurface, err = NewSurfaceFromCanvas(canvas)
	return d, err
}

// NewDisplayID creates a new display for the canvas with the given ID. An error is returned
// if no canvas was found.
func NewDisplayID(canvasID string) (*Display, error) {
	d := &Display{}
	var err error
	d.frontSurface, err = NewSurfaceFromCanvasID(canvasID)
	return d, err
}

// SetMode initalizes the Display.
func (d *Display) SetMode(width, height int) {
	if d.canvas == nil {
		d.canvas = jq("<canvas>").Get(0)
	}
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
