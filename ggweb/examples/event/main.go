package main

import (
	"image/color"
	"time"

	"github.com/Bredgren/gogame/event"
	"github.com/Bredgren/gogame/ggweb"
	"github.com/Bredgren/gogame/key"
)

func main() {
	ggweb.Init(onReady)
}

var display *ggweb.Surface

func onReady() {
	r := ggweb.WindowRect()
	width, height := int(r.W), int(r.H)
	display = ggweb.NewSurfaceFromID("main")
	display.SetSize(width, height)
	display.StyleColor(ggweb.Fill, color.Black)
	display.DrawRect(ggweb.Fill, display.Rect())

	// Prevent scrolling
	ggweb.PreventKeyDefault[key.Space] = true

	ggweb.SetMainLoop(mainLoop)
}

func mainLoop(t time.Duration) {
	for evt := event.Poll(); evt.Type != event.NoEvent; evt = event.Poll() {
		ggweb.Log(evt)
		switch evt.Type {
		case event.KeyDown:
			data := evt.Data.(event.KeyData)
			k := data.Key
			if k.Rune != '\u0000' {
				if k.ShiftRune != '\u0000' && (data.Mod[key.LShift] || data.Mod[key.RShift]) {
					ggweb.Log(string(k.ShiftRune))
				} else {
					ggweb.Log(string(k.Rune))
				}
			} else {
				ggweb.Log(k.Name)
			}
		}
	}

	display.StyleColor(ggweb.Fill, color.Black)
	display.DrawRect(ggweb.Fill, display.Rect())
}
