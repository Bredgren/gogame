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
// package. Don't forget to call RegisterEvents for the surface you would like to receive
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

// PromptBeforeQuit will ask the user for confirmation before leaving the page if not an
// empty string. The browser may or may not use this string in its message to the user.
var PromptBeforeQuit string

var DisableContextMenu bool

func addGlobalEvents() {
	js.Global.Call("addEventListener", "resize", func(e *js.Object) {
		if err := event.Post(event.Event{
			Type: event.WindowResize,
			Data: event.ResizeData{
				W: js.Global.Get("innerWidth").Int(),
				H: js.Global.Get("innerHeight").Int(),
			},
		}); err != nil {
			Warn("Event skipped because queue is full", e)
		}
	})

	js.Global.Call("addEventListener", "beforeunload", func(e *js.Object) {
		// The Quit event is a bit tricky for the browser, and I don't think it's worth the
		// trouble right now. We'll just stick with providing an easy way to prompt before
		// leaving.
		// if err := event.Post(event.Event{Type: event.Quit}); err != nil {
		// 	Warn("Event skipped because queue is full", e)
		// }
		if PromptBeforeQuit != "" {
			e.Set("returnValue", PromptBeforeQuit)
		}
	})
	js.Global.Call("addEventListener", "contextmenu", func(e *js.Object) {
		if DisableContextMenu {
			e.Call("preventDefault")
		}
	})

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
		if PreventKeyDefault[k] {
			e.Call("preventDefault")
		}

		keyState[k] = false
		if err := event.Post(event.Event{
			Type: event.KeyUp,
			Data: event.KeyData{Key: k, Mod: ModKeys()},
		}); err != nil {
			Warn("Event skipped because queue is full", e)
		}
	})
}

func handleMouseMove(e *js.Object) {
	x, y := e.Get("offsetX").Float(), e.Get("offsetY").Float()
	dx, dy := e.Get("movementX").Float(), e.Get("movementY").Float()
	mouseState.Pos.X = x
	mouseState.Pos.Y = y
	mouseState.Rel.X = dx
	mouseState.Rel.Y = dy
	if err := event.Post(event.Event{
		Type: event.MouseMotion,
		Data: event.MouseMotionData{
			Pos:     geo.Vec{X: x, Y: y},
			Rel:     geo.Vec{X: dx, Y: dy},
			Buttons: MousePressed(),
		},
	}); err != nil {
		Warn("Event skipped because queue is full", e)
	}
}

func handleMouseDown(e *js.Object) {
	button := e.Get("button").Int()
	mouseState.Buttons[button] = true
	if err := event.Post(event.Event{
		Type: event.MouseButtonDown,
		Data: event.MouseData{
			Pos: geo.Vec{
				X: e.Get("offsetX").Float(),
				Y: e.Get("offsetY").Float(),
			},
			Button: button,
		},
	}); err != nil {
		Warn("Event skipped because queue is full", e)
	}
}

func handleMouseUp(e *js.Object) {
	button := e.Get("button").Int()
	mouseState.Buttons[button] = false
	if err := event.Post(event.Event{
		Type: event.MouseButtonUp,
		Data: event.MouseData{
			Pos: geo.Vec{
				X: e.Get("offsetX").Float(),
				Y: e.Get("offsetY").Float(),
			},
			Button: button,
		},
	}); err != nil {
		Warn("Event skipped because queue is full", e)
	}
}

var eventsRegisteredTo *Surface

// RegisterEvents sets up the surface to receive mouse events. Only one surface can accept
// events at a time, calling RegisterEvents on a multiple surfaces will unregister them on
// previous ones.
func RegisterEvents(s *Surface) {
	if eventsRegisteredTo != nil {
		UnregisterEvents(eventsRegisteredTo)
	}

	s.Canvas.Call("addEventListener", "mousemove", handleMouseMove)
	s.Canvas.Call("addEventListener", "mousedown", handleMouseDown)
	s.Canvas.Call("addEventListener", "mouseup", handleMouseUp)

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

// UnregisterEvents causes the surface to stop receiving events.
func UnregisterEvents(s *Surface) {
	s.Canvas.Call("removeEventListener", "mousemove", handleMouseMove)
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
		mouseState.Rel.X, mouseState.Rel.Y = 0, 0
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

var mouseState = struct {
	Buttons map[int]bool
	Pos     geo.Vec
	Rel     geo.Vec
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
func MousePos() geo.Vec {
	return mouseState.Pos
}

// MouseRel returns the last relative change in mouse position.
func MouseRel() geo.Vec {
	return mouseState.Rel
}

// WindowRect returns a rectangle that covers the entire inner window of the browser.
func WindowRect() geo.Rect {
	return geo.Rect{
		W: js.Global.Get("innerWidth").Float(),
		H: js.Global.Get("innerHeight").Float(),
	}
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
