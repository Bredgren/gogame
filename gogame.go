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

var canvas *Canvas

// Ready returns a channel that will send one item before closing, signaling that the
// page has loaded and gogame is ready to be used.
func Ready() chan struct{} {
	ch := make(chan struct{}, 1)
	jq("body").SetAttr("onload", func() {
		SetCanvas(jq("canvas").Get(0))
		log.Println("gogame ready")
		ch <- struct{}{}
		close(ch)
	})
	return ch
}

// SetCanvas changes the canvas being used. Only one canvas may be used at a time. If
// unset then gogame will default to the first canvas in the DOM.
func SetCanvas(c *js.Object) {
	canvas = newCanvas(c)
}

// GetCanvas returns the canvas object being used
func GetCanvas() *Canvas {
	return canvas
}

// SetFullscreen sets or unsetd fullscreen mode
func SetFullscreen(fullscreen bool) {
	//canvas.canvas.Call("requestFullScreen")
	canvas.canvas.Call("webkitRequestFullScreen")
	//canvas.canvas.Call("mozRequestFullScreen")
}

// Log prints to the console
func Log(args ...interface{}) {
	console.Call("log", args)
}
