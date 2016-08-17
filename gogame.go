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

var display *Display

// Ready returns a channel that will send one item before closing, signaling that the
// page has loaded and gogame is ready to be used. There must be a canvas element in the
// DOM for Ready to succeed. It will use this canvas as the main Display. A different
// canvas may be specified afterward, if desired, using the SetMainDisplay function.
func Ready() chan struct{} {
	ch := make(chan struct{}, 1)
	jq("body").SetAttr("onload", func() {
		d, err := NewDisplay(jq("canvas").Get(0))
		if err != nil {
			panic("gogame requires there to be a canvas in the DOM")
		}
		SetMainDisplay(d)
		log.Println("gogame ready")
		ch <- struct{}{}
		close(ch)
	})
	return ch
}

// SetMainDisplay changes the main canvas being used. If unset then gogame will default
// to the first canvas in the DOM.
func SetMainDisplay(d *Display) {
	display = d
}

// GetDisplay returns the main Display being used.
func GetDisplay() *Display {
	return display
}

// SetFullscreen sets or unsetd fullscreen mode.
func SetFullscreen(fullscreen bool) {
	//canvas.canvas.Call("requestFullScreen")
	display.canvas.Call("webkitRequestFullScreen")
	//canvas.canvas.Call("mozRequestFullScreen")
}

// Log prints to the console.
func Log(args ...interface{}) {
	console.Call("log", args)
}
