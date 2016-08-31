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

// SetFullscreen sets or unsetd fullscreen mode.
func SetFullscreen(fullscreen bool) {
	//canvas.canvas.Call("requestFullScreen")
	display.canvas.Call("webkitRequestFullScreen")
	//canvas.canvas.Call("mozRequestFullScreen")
}

// Log prints to the console.
func Log(args ...interface{}) {
	console.Call("log", args...)
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
	// Setup event listeners for the display.
	js.Global.Call("addEventListener", jquery.RESIZE, func(e *js.Object) {
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

	js.Global.Call("addEventListener", jquery.KEYDOWN, func(e *js.Object) {
		k := key.FromJsEvent(e)
		// Ignore key repeats
		if keyState[k] {
			return
		}
		keyState[k] = true
		if err := event.Post(event.Event{
			Type: event.KeyDown,
			Data: event.KeyData{
				Key: k,
				Mod: GetModKeys(),
			},
		}); err != nil {
			Log("Warning: event skipped because queue is full", e)
		}
	})

	js.Global.Call("addEventListener", jquery.KEYUP, func(e *js.Object) {
		k := key.FromJsEvent(e)
		keyState[k] = false
		if err := event.Post(event.Event{
			Type: event.KeyUp,
			Data: event.KeyData{
				Key: k,
				Mod: GetModKeys(),
			},
		}); err != nil {
			Log("Warning: event skipped because queue is full", e)
		}
	})

	canvas := display.frontSurface.GetCanvas()

	canvas.Call("addEventListener", jquery.MOUSEMOVE, func(e *js.Object) {
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
				Buttons: GetMousePressed(),
			},
		}); err != nil {
			Log("Warning: event skipped because queue is full", e)
		}
	})

	canvas.Call("addEventListener", jquery.MOUSEDOWN, func(e *js.Object) {
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

	canvas.Call("addEventListener", jquery.MOUSEUP, func(e *js.Object) {
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

// GetMousePressed returns a map which indicates which mouse buttons are pressed.
func GetMousePressed() map[int]bool {
	m := make(map[int]bool)
	for b, press := range mouseState.Buttons {
		if press {
			m[b] = true
		}
	}
	return m
}

// GetMousePos returns the mouses current x and y positions.
func GetMousePos() (x, y float64) {
	return mouseState.PosX, mouseState.PosY
}

// GetMouseRel returns the last relative change in mouse position.
func GetMouseRel() (dx, dy float64) {
	return mouseState.RelX, mouseState.RelY
}
