// Package gogame is wrapper around gopherjs that makes it more convenient to work with for making games.
// It also provides several utilities commonly used in games.
package gogame

import (
	"log"
	"time"

	"github.com/Bredgren/gogame/event"
	"github.com/Bredgren/gogame/key"
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
// to the first canvas in the DOM. This is also the only display that will receive input
// events.
func SetMainDisplay(d *Display) {
	unsetupDisplay()
	display = d
	setupDisplay()
}

// GetDisplay returns the main Display being used.
func GetDisplay() *Display {
	return display
}

// MainLoop is a callback function that returns a time value that can be compared to
// previous calls to determine the elapsed time.
type MainLoop func(time.Duration)

var mainLoop *js.Object

// SetMainLoop sets the callback for the main game loop. The given function will be
// called at a regular interval.
func SetMainLoop(loop MainLoop) {
	var f func(timestamp *js.Object)
	f = func(timestamp *js.Object) {
		mainLoop = js.Global.Call("requestAnimationFrame", f)
		loop(time.Duration(timestamp.Float()) * time.Millisecond)
	}
	f(&js.Object{})
}

// UnsetMainLoop stops calling the main game loop.
func UnsetMainLoop() {
	if mainLoop != nil {
		js.Global.Call("cancelAnimationFrame", mainLoop)
		mainLoop = nil
	}
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

func unsetupDisplay() {
	// Clean up display before we stop using it by removing event listeners.
	if display == nil {
		return
	}
	canvas := display.frontSurface.GetCanvas()
	canvas.Call("removeEventListener")
	js.Global.Call("removeEventListener")
}

func setupDisplay() {
	// Setup envent listeners for the display.
	canvas := display.frontSurface.GetCanvas()
	js.Global.Call("addEventListener", jquery.KEYDOWN, func(e *js.Object) {
		k := key.FromJsEvent(e)
		keyState[k] = true
		err := event.Post(event.Event{
			Type: event.KeyDown,
			Data: event.KeyData{
				Key: k,
				Mod: GetModKeys(),
			},
		})
		if err != nil {
			Log("Warning: event skipped because queue is full", e)
		}
	})

	js.Global.Call("addEventListener", jquery.KEYUP, func(e *js.Object) {
		k := key.FromJsEvent(e)
		keyState[k] = false
		err := event.Post(event.Event{
			Type: event.KeyUp,
			Data: event.KeyData{
				Key: k,
				Mod: GetModKeys(),
			},
		})
		if err != nil {
			Log("Warning: event skipped because queue is full", e)
		}
	})

	canvas.Call("addEventListener", jquery.MOUSEMOVE, func(e *js.Object) {
		Log("mousemove", e)
	})
}

var keyState = map[key.Key]bool{}

// GetPressedKeys returns the current state of every key. If a Key maps to true then it is pressed.
func GetPressedKeys() map[key.Key]bool {
	m := make(map[key.Key]bool)
	for k, press := range keyState {
		if press {
			m[k] = true
		}
	}
	return m
}

// GetModKeys returns just the stat for the modifier keys.
func GetModKeys() map[key.Key]bool {
	m := make(map[key.Key]bool)
	for k, press := range keyState {
		if press && k.IsMod() {
			m[k] = true
		}
	}
	return m
}
