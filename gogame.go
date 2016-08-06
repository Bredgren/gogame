// Package gogame is wrapper around gopherjs that makes it more convenient to work with for making games.
// It also provides several utilities commonly used in games.
package gogame

import (
	"log"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/jquery"
)

var jq = jquery.NewJQuery
var console = js.Global.Get("console")

var canvas *js.Object
var ctx *js.Object

// Ready returns a channel that will send one item before closing, signaling that the
// page has loaded and gogame is ready to be used.
func Ready() chan struct{} {
	ch := make(chan struct{}, 1)
	jq("body").SetAttr("onload", func() {
		SetCanvas(jq("canvas").Get(0))
		log.Println("gogame ready")
		console.Call("log", canvas)
		ch <- struct{}{}
		close(ch)
	})
	return ch
}

// SetCanvas changes the canvas being used. Only one canvas may be used at a time. If
// unset then gogame will default to the first canvas in the DOM.
func SetCanvas(newCanvas *js.Object) {
	canvas = newCanvas
	ctx = canvas.Call("getContext", "2d")
}

// Canvas returns the canvas object being used
func Canvas() *js.Object {
	return canvas
}

// SetCanvasResolution sets the width and height of the pixels within the canvas
func SetCanvasResolution(width, height int) {
	canvas.Set("width", width)
	canvas.Set("height", height)
}

// SetFullscreen sets or unsetd fullscreen mode
func SetFullscreen(fullscreen bool) {

}

// FillCanvas fills the whole canvas with one color
func FillCanvas(color Color) {
	ctx.Set("fillStyle", color)
	ctx.Call("fillRect", 0, 0, canvas.Get("width"), canvas.Get("height"))
}

// Log prints to the console
func Log(args ...interface{}) {
	console.Call("log", args)
}
