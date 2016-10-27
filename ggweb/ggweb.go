package ggweb

import (
	"time"

	"github.com/Bredgren/gogame/event"
	"github.com/Bredgren/gogame/geo"
	"github.com/Bredgren/gogame/key"
	"github.com/gopherjs/gopherjs/js"
)

var ready bool
var console *js.Object

// Init sets up ggweb and waits for the page to load then calls the onReady function.
// Calling Init more than once will have no effect. Init takes care of setting up the events
// for window resizing, quiting, and key presses. These can be retrieved via the gogame/events
// package. Don't forget to call RegisterEvents for the surface you would like to recieve
// the others event types.
func Init(onReady func()) {
	onload := func() {
		if ready {
			return
		}
		addGlobalEvents()
		go onReady()
		ready = true
	}
	if js.Global.Get("document").Get("readyState").String() == "complete" {
		onload()
		return
	}
	js.Global.Get("document").Call("addEventListener", "DOMContentLoaded", onload, false)
	js.Global.Call("addEventListener", "load", onload, false)

	// In order for tests (like color_test.go) to work they can't use any gopherjs things.
	// initializing console in the global scope causes tests to try and use js so they fail,
	// that is why we initialize console here.
	console = js.Global.Get("console")
}

func addGlobalEvents() {
	// js.Global.Call("addEventListener", "resize", func(e *js.Object) {
	// 	if err := event.Post(event.Event{
	// 		Type: event.VideoResize,
	// 		Data: event.ResizeData{
	// 			W: js.Global.Get("innerWidth").Int(),
	// 			H: js.Global.Get("innerHeight").Int(),
	// 		},
	// 	}); err != nil {
	// 		Warn("Event skipped because queue is full", e)
	// 	}
	// })

	// js.Global.Set("onbeforeunload", func(e *js.Object) {
	// 	if err := event.Post(event.Event{Type: event.Quit}); err != nil {
	// 		Warn("Event skipped because queue is full", e)
	// 	}
	// })

	js.Global.Call("addEventListener", "keydown", func(e *js.Object) {
		k := EventToKey(e)

		if PreventKeyDefault[k] {
			e.Call("preventDefault")
		}

		// Ignore key repeats
		if keyState[k] {
			return
		}

		keyState[k] = true
		if err := event.Post(event.Event{
			Type: event.KeyDown,
			Data: event.KeyData{Key: k, Mod: ModKeys()},
		}); err != nil {
			Warn("Event skipped because queue is full", e)
		}
	})

	js.Global.Call("addEventListener", "keyup", func(e *js.Object) {
		k := EventToKey(e)
		keyState[k] = false
		if err := event.Post(event.Event{
			Type: event.KeyUp,
			Data: event.KeyData{Key: k, Mod: ModKeys()},
		}); err != nil {
			Warn("Event skipped because queue is full", e)
		}
	})
}

func RegisterEvents(s *Surface) {
	// 	canvas := display.frontSurface.Canvas()

	// 	canvas.Call("addEventListener", "mousemove", func(e *js.Object) {
	// 		x, y := e.Get("offsetX").Float(), e.Get("offsetY").Float()
	// 		dx, dy := e.Get("movementX").Float(), e.Get("movementY").Float()
	// 		mouseState.PosX = x
	// 		mouseState.PosY = y
	// 		mouseState.RelX = dx
	// 		mouseState.RelY = dy
	// 		if err := event.Post(event.Event{
	// 			Type: event.MouseMotion,
	// 			Data: event.MouseMotionData{
	// 				Pos:     struct{ X, Y float64 }{X: x, Y: y},
	// 				Rel:     struct{ X, Y float64 }{X: dx, Y: dy},
	// 				Buttons: MousePressed(),
	// 			},
	// 		}); err != nil {
	// 			Warn("Event skipped because queue is full", e)
	// 		}
	// 	})

	// 	canvas.Call("addEventListener", "mousedown", func(e *js.Object) {
	// 		button := e.Get("button").Int()
	// 		mouseState.Buttons[button] = true
	// 		if err := event.Post(event.Event{
	// 			Type: event.MouseButtonDown,
	// 			Data: event.MouseData{
	// 				Pos: struct{ X, Y float64 }{
	// 					X: e.Get("offsetX").Float(),
	// 					Y: e.Get("offsetY").Float(),
	// 				},
	// 				Button: button,
	// 			},
	// 		}); err != nil {
	// 			Warn("Event skipped because queue is full", e)
	// 		}
	// 	})

	// 	canvas.Call("addEventListener", "mouseup", func(e *js.Object) {
	// 		button := e.Get("button").Int()
	// 		mouseState.Buttons[button] = false
	// 		if err := event.Post(event.Event{
	// 			Type: event.MouseButtonUp,
	// 			Data: event.MouseData{
	// 				Pos: struct{ X, Y float64 }{
	// 					X: e.Get("offsetX").Float(),
	// 					Y: e.Get("offsetY").Float(),
	// 				},
	// 				Button: button,
	// 			},
	// 		}); err != nil {
	// 			Warn("Event skipped because queue is full", e)
	// 		}
	// 	})

	// 	canvas.Call("addEventListener", "wheel", func(e *js.Object) {
	// 		dx, dy, dz := e.Get("deltaX").Float(), e.Get("deltaY").Float(), e.Get("deltaZ").Float()
	// 		if err := event.Post(event.Event{
	// 			Type: event.MouseWheel,
	// 			Data: event.MouseWheelData{
	// 				Dx: dx,
	// 				Dy: dy,
	// 				Dz: dz,
	// 			},
	// 		}); err != nil {
	// 			Warn("Event skipped because queue is full", e)
	// 		}
	// 	})
}

func UnRegisterEvents(s *Surface) {
}

// PreventKeyDefault is a set of key that should have their default behavior prevented.
var PreventKeyDefault = map[key.Key]bool{}

// var PreventDefaultMouse = map[int]bool{}

// // Stats holds various bits of information that one may find useful.
// var Stats = struct {
// 	// LoopDuration is the amount of time that the last execution of the main loop took.
// 	LoopDuration time.Duration
// }{}

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
		// start := time.Now()
		loop(time.Duration(timestamp.Float()) * time.Millisecond)
		// Stats.LoopDuration = time.Now().Sub(start)
		// mouseState.RelX, mouseState.RelY = 0, 0
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

// // var isFullscreen bool

// // SetFullscreen sets or unsetd fullscreen mode.
// // func SetFullscreen(fullscreen bool) {
// // 	// display.canvas.Call("requestFullScreen")
// // 	display.frontSurface.Canvas().Call("webkitRequestFullScreen")
// // 	// display.canvas.Call("mozRequestFullScreen")
// // 	isFullscreen = fullscreen
// // }

// // Fullscreen returns true if fullscreen is currently active.
// // func Fullscreen() bool {
// // 	return isFullscreen
// // }

// Log prints to the console. This won't work until ggweb is initialized.
func Log(args ...interface{}) {
	console.Call("log", args...)
}

// Warn prints a warning to the console. This won't work until ggweb is initialized.
func Warn(args ...interface{}) {
	console.Call("warn", args...)
}

// Info prints an info log to the console. This won't work until ggweb is initialized.
func Info(args ...interface{}) {
	console.Call("info", args...)
}

// Error prints an error to the console. This won't work until ggweb is initialized.
func Error(args ...interface{}) {
	console.Call("error", args...)
}

var keyState = map[key.Key]bool{}

// var mouseState = struct {
// 	Buttons    map[int]bool
// 	PosX, PosY float64
// 	RelX, RelY float64
// }{
// 	Buttons: make(map[int]bool),
// }

// // PressedKeys returns a map that contoins all pressed keys mapping to true.
// func PressedKeys() map[key.Key]bool {
// 	m := make(map[key.Key]bool)
// 	for k, press := range keyState {
// 		if press {
// 			m[k] = true
// 		}
// 	}
// 	return m
// }

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

// // MousePressed returns a map that contains all pressed mouse buttons mapping to true.
// func MousePressed() map[int]bool {
// 	m := make(map[int]bool)
// 	for b, press := range mouseState.Buttons {
// 		if press {
// 			m[b] = true
// 		}
// 	}
// 	return m
// }

// // MousePos returns the mouses current x and y positions.
// func MousePos() (x, y float64) {
// 	return mouseState.PosX, mouseState.PosY
// }

// // MouseRel returns the last relative change in mouse position.
// func MouseRel() (dx, dy float64) {
// 	return mouseState.RelX, mouseState.RelY
// }

// WindowRect returns a rectangle that covers the entire inner window of the browser.
func WindowRect() geo.Rect {
	return geo.Rect{
		W: js.Global.Get("innerWidth").Float(),
		H: js.Global.Get("innerHeight").Float(),
	}
}

// // LocalStorageGet retrieves the value associated with the given key. If there is no value
// // then ok will be false.
// func LocalStorageGet(key string) (val string, ok bool) {
// 	v := js.Global.Get("localStorage").Call("getItem", key)
// 	if v == nil {
// 		return "", false
// 	}
// 	return v.String(), true
// }

// // LocalStorageSet sets the given key's value to val.
// func LocalStorageSet(key, val string) {
// 	js.Global.Get("localStorage").Call("setItem", key, val)
// }

// // LocalStorageRemove removes the given key (and it's value) from local storage.
// func LocalStorageRemove(key string) {
// 	js.Global.Get("localStorage").Call("removeItem", key)
// }
