package gogame

import (
	"github.com/Bredgren/gogame/geo"
	"github.com/gopherjs/gopherjs/js"
)

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
	d.frontSurface.Canvas().Set("width", width)
	d.frontSurface.Canvas().Set("height", height)
}

// Flip draws the back Surface onto the Display, making the draw operations since the last
// call to Flip visible.
func (d *Display) Flip() {
	if d == nil {
		return
	}
	d.frontSurface.Blit(d, 0, 0)
}

// Update is like Flip but only updates the portions of the front surface defined by the
// list of Rects.
func (d *Display) Update(rects []geo.Rect) {
	if d == nil {
		return
	}
	for _, r := range rects {
		d.frontSurface.BlitArea(d, r, r.X, r.Y)
	}
}

// SetCursor sets the appearence of the cursor when it is over this Display.
func (d *Display) SetCursor(c Cursor) {
	if d == nil {
		return
	}
	d.frontSurface.Canvas().Get("style").Set("cursor", c)
}

// Cursor returns the current appearence of the cursor when it is over the Display.
func (d *Display) Cursor() Cursor {
	if d == nil {
		return CursorDefault
	}
	c := d.frontSurface.Canvas().Get("style").Get("cursor").String()
	if c == "" {
		return CursorDefault
	}
	return Cursor(c)
}

// Cursor is a style for the cursor. It can be any valid css value for the cursor property,
// most of which are predefined in this package.
type Cursor string

const (
	CursorDefault      Cursor = "default"
	CursorNone         Cursor = "none"
	CursorPointer      Cursor = "pointer"
	CursorText         Cursor = "text"
	CursorVerticalText Cursor = "vertical-text"
	CursorProgress     Cursor = "progress"
	CusorWait          Cursor = "wait"
	CursorAlias        Cursor = "alias"
	CursorAllScroll    Cursor = "all-scroll"
	CursorMove         Cursor = "move"
	CursorCell         Cursor = "cell"
	CursorCopy         Cursor = "copy"
	CursorCrosshair    Cursor = "crosshair"
	CursorNSResize     Cursor = "ns-resize"
	CursorEWResize     Cursor = "ew-resize"
	CursorNESWResize   Cursor = "nesw-resize"
	CursorNWSEResize   Cursor = "nwse-resize"
	CursorRowReszie    Cursor = "row-resize"
	CursorColResize    Cursor = "col-resize"
	CursorHelp         Cursor = "help"
	CursorNoDrop       Cursor = "no-drop"
	CursorNotAllowed   Cursor = "not-allowed"
)
