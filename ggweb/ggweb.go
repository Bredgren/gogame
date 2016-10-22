package ggweb

import (
	"log"

	"github.com/gopherjs/gopherjs/js"
)

var ready bool

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
		log.Println("ggweb ready")
		go onReady()
		ready = true
	}
	if js.Global.Get("document").Get("readyState").String() == "complete" {
		onload()
		return
	}
	js.Global.Get("document").Call("addEventListener", "DOMContentLoaded", onload, false)
	js.Global.Call("addEventListener", "load", onload, false)

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
	// 		Log("Warning: event skipped because queue is full", e)
	// 	}
	// })

	// js.Global.Set("onbeforeunload", func(e *js.Object) {
	// 	if err := event.Post(event.Event{Type: event.Quit}); err != nil {
	// 		Log("Warning: event skipped because queue is full", e)
	// 	}
	// })

	// js.Global.Call("addEventListener", "keydown", func(e *js.Object) {
	// 	k := key.FromJsEvent(e)
	// 	// Ignore key repeats
	// 	if keyState[k] {
	// 		return
	// 	}
	// 	keyState[k] = true
	// 	if err := event.Post(event.Event{
	// 		Type: event.KeyDown,
	// 		Data: event.KeyData{Key: k, Mod: ModKeys()},
	// 	}); err != nil {
	// 		Log("Warning: event skipped because queue is full", e)
	// 	}
	// })

	// js.Global.Call("addEventListener", "keyup", func(e *js.Object) {
	// 	k := key.FromJsEvent(e)
	// 	keyState[k] = false
	// 	if err := event.Post(event.Event{
	// 		Type: event.KeyUp,
	// 		Data: event.KeyData{Key: k, Mod: ModKeys()},
	// 	}); err != nil {
	// 		Log("Warning: event skipped because queue is full", e)
	// 	}
	// })
}

func RegisterEvents(s *Surface) {
}

func UnRegisterEvents(s *Surface) {
}
