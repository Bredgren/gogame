// Package gogame is wrapper around gopherjs that makes it more convenient to work with for making games.
// It also provides several utilities commonly used in games.
//
// TODO:
// - Support for browsers other than Chrome
// - Sprite?
// - Spritesheet?
// - Primitive shape collision?
//   - Circles, rectangles, lines, rays
// - Network?
package gogame

import (
	"log"
	"time"

	"github.com/Bredgren/gogame/event"
	"github.com/Bredgren/gogame/key"
	"github.com/gopherjs/gopherjs/js"
)

var console = js.Global.Get("console")

var display *Display

// Ready calls the callback function when the page has loaded and gogame is ready to be
// used. The first canvas in the DOM will be used as the initial main display. If there
// is no canvas, or you would like to set the main display to a different one then use
// the SetMainDisplay function.
func Ready(callback func()) {
	onload := func() {
		d, err := NewDisplay(js.Global.Get("document").Call("getElementsByTagName", "canvas").Index(0))
		if err != nil {
			// No canvas available, do nothing.
		}
		SetMainDisplay(d)
		log.Println("gogame ready")
		go callback()
	}
	if js.Global.Get("document").Get("readyState").String() == "complete" {
		onload()
		return
	}
	js.Global.Get("document").Call("addEventListener", "DOMContentLoaded", onload, false)
	js.Global.Call("addEventListener", "load", onload, false)
}

// SetMainDisplay changes the main canvas being used. If unset then gogame will default
// to the first canvas in the DOM. This is also the only display that will receive input
// events.
func SetMainDisplay(d *Display) {
	unsetupDisplay()
	display = d
	setupDisplay()
}

// MainDisplay returns the main Display being used.
func MainDisplay() *Display {
	return display
}

// Stats holds various bits of information that one may find useful.
var Stats = struct {
	// LoopDuration is the amount of time that the last execution of the main loop took.
	LoopDuration time.Duration
}{}

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
		start := time.Now()
		loop(time.Duration(timestamp.Float()) * time.Millisecond)
		Stats.LoopDuration = time.Now().Sub(start)
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

// var isFullscreen bool

// SetFullscreen sets or unsetd fullscreen mode.
// func SetFullscreen(fullscreen bool) {
// 	// display.canvas.Call("requestFullScreen")
// 	display.frontSurface.Canvas().Call("webkitRequestFullScreen")
// 	// display.canvas.Call("mozRequestFullScreen")
// 	isFullscreen = fullscreen
// }

// Fullscreen returns true if fullscreen is currently active.
// func Fullscreen() bool {
// 	return isFullscreen
// }

// Log prints to the console.
func Log(args ...interface{}) {
	console.Call("log", args...)
}

func unsetupDisplay() {
	// Clean up display before we stop using it by removing event listeners.
	if display == nil {
		return
	}
	canvas := display.frontSurface.Canvas()
	canvas.Call("removeEventListener")
	js.Global.Call("removeEventListener")
}

func setupDisplay() {
	// Setup event listeners for the display.
	js.Global.Call("addEventListener", "resize", func(e *js.Object) {
		if err := event.Post(event.Event{
			Type: event.VideoResize,
			Data: event.ResizeData{
				W: js.Global.Get("innerWidth").Int(),
				H: js.Global.Get("innerHeight").Int(),
			},
		}); err != nil {
			Log("Warning: event skipped because queue is full", e)
		}
	})

	js.Global.Set("onbeforeunload", func(e *js.Object) {
		if err := event.Post(event.Event{Type: event.Quit}); err != nil {
			Log("Warning: event skipped because queue is full", e)
		}
	})

	js.Global.Call("addEventListener", "keydown", func(e *js.Object) {
		k := key.FromJsEvent(e)
		// Ignore key repeats
		if keyState[k] {
			return
		}
		keyState[k] = true
		if err := event.Post(event.Event{
			Type: event.KeyDown,
			Data: event.KeyData{Key: k, Mod: ModKeys()},
		}); err != nil {
			Log("Warning: event skipped because queue is full", e)
		}
	})

	js.Global.Call("addEventListener", "keyup", func(e *js.Object) {
		k := key.FromJsEvent(e)
		keyState[k] = false
		if err := event.Post(event.Event{
			Type: event.KeyUp,
			Data: event.KeyData{Key: k, Mod: ModKeys()},
		}); err != nil {
			Log("Warning: event skipped because queue is full", e)
		}
	})

	canvas := display.frontSurface.Canvas()

	canvas.Call("addEventListener", "mousemove", func(e *js.Object) {
		x, y := e.Get("offsetX").Float(), e.Get("offsetY").Float()
		dx, dy := e.Get("movementX").Float(), e.Get("movementY").Float()
		mouseState.PosX = x
		mouseState.PosY = y
		mouseState.RelX = dx
		mouseState.RelY = dy
		if err := event.Post(event.Event{
			Type: event.MouseMotion,
			Data: event.MouseMotionData{
				Pos:     struct{ X, Y float64 }{X: x, Y: y},
				Rel:     struct{ Dx, Dy float64 }{Dx: dx, Dy: dy},
				Buttons: MousePressed(),
			},
		}); err != nil {
			Log("Warning: event skipped because queue is full", e)
		}
	})

	canvas.Call("addEventListener", "mousedown", func(e *js.Object) {
		button := e.Get("button").Int()
		mouseState.Buttons[button] = true
		if err := event.Post(event.Event{
			Type: event.MouseButtonDown,
			Data: event.MouseData{
				Pos: struct{ X, Y float64 }{
					X: e.Get("offsetX").Float(),
					Y: e.Get("offsetY").Float(),
				},
				Button: button,
			},
		}); err != nil {
			Log("Warning: event skipped because queue is full", e)
		}
	})

	canvas.Call("addEventListener", "mouseup", func(e *js.Object) {
		button := e.Get("button").Int()
		mouseState.Buttons[button] = false
		if err := event.Post(event.Event{
			Type: event.MouseButtonUp,
			Data: event.MouseData{
				Pos: struct{ X, Y float64 }{
					X: e.Get("offsetX").Float(),
					Y: e.Get("offsetY").Float(),
				},
				Button: button,
			},
		}); err != nil {
			Log("Warning: event skipped because queue is full", e)
		}
	})
}

var keyState = map[key.Key]bool{}
var mouseState = struct {
	Buttons    map[int]bool
	PosX, PosY float64
	RelX, RelY float64
}{
	Buttons: make(map[int]bool),
}

// PressedKeys returns a map that contoins all pressed keys mapping to true.
func PressedKeys() map[key.Key]bool {
	m := make(map[key.Key]bool)
	for k, press := range keyState {
		if press {
			m[k] = true
		}
	}
	return m
}

// ModKeys returns just the state for the modifier keys.
func ModKeys() map[key.Key]bool {
	m := make(map[key.Key]bool)
	for k, press := range keyState {
		if press && k.IsMod() {
			m[k] = true
		}
	}
	return m
}

// MousePressed returns a map that contains all pressed mouse buttons mapping to true.
func MousePressed() map[int]bool {
	m := make(map[int]bool)
	for b, press := range mouseState.Buttons {
		if press {
			m[b] = true
		}
	}
	return m
}

// MousePos returns the mouses current x and y positions.
func MousePos() (x, y float64) {
	return mouseState.PosX, mouseState.PosY
}

// MouseRel returns the last relative change in mouse position.
func MouseRel() (dx, dy float64) {
	return mouseState.RelX, mouseState.RelY
}

// LocalStorageGet retrieves the value associated with the given key. If there is no value
// then ok will be false.
func LocalStorageGet(key string) (val string, ok bool) {
	v := js.Global.Get("localStorage").Call("getItem", key)
	if v == nil {
		return "", false
	}
	return v.String(), true
}

// LocalStorageSet sets the given key's value to val.
func LocalStorageSet(key, val string) {
	js.Global.Get("localStorage").Call("setItem", key, val)
}

// LocalStorageRemove removes the given key (and it's value) from local storage.
func LocalStorageRemove(key string) {
	js.Global.Get("localStorage").Call("removeItem", key)
}
