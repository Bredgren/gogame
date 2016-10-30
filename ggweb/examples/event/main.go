package main

import (
	"fmt"
	"image/color"
	"math"
	"time"

	"github.com/Bredgren/gogame/event"
	"github.com/Bredgren/gogame/ggweb"
	"github.com/Bredgren/gogame/key"
)

func main() {
	ggweb.Init(onReady)
}

var display *ggweb.Surface
var charSurf *ggweb.Surface
var evtSurf *ggweb.Surface

func onReady() {
	r := ggweb.WindowRect()
	width, height := int(r.W), int(r.H)
	display = ggweb.NewSurfaceFromID("main")
	display.SetSize(width, height)
	display.StyleColor(ggweb.Fill, color.Black)
	display.DrawRect(ggweb.Fill, display.Rect())

	ggweb.RegisterEvents(display)

	charSurf = ggweb.NewSurface(width, height)
	evtSurf = ggweb.NewSurface(100, 150)

	// Prevent scrolling
	ggweb.PreventKeyDefault[key.Space] = true
	ggweb.PreventKeyDefault[key.Tab] = true

	// ggweb.PromptBeforeQuit = "Are you sure?"
	ggweb.DisableContextMenu = true

	ggweb.SetMainLoop(mainLoop)
}

const (
	padding     = 0
	charPadding = 0
)

var (
	charX float64 = padding
	charY float64 = padding
)

var font1 = ggweb.Font{
	Size:   15,
	Family: ggweb.FontFamilyMonospace,
}

var font2 = ggweb.Font{
	Size:   15,
	Family: ggweb.FontFamilyMonospace,
}

func mainLoop(t time.Duration) {
	newChars := []rune{}
	evtLogs := []string{}
	for evt := event.Poll(); evt.Type != event.NoEvent; evt = event.Poll() {
		switch evt.Type {
		case event.WindowResize:
			data := evt.Data.(event.ResizeData)
			evtLogs = append(evtLogs, fmt.Sprintf("resize: (%d, %d)", data.W, data.H))
			display.SetSize(data.W, data.H)
			display.StyleColor(ggweb.Fill, color.Black)
			display.DrawRect(ggweb.Fill, display.Rect())
			c := charSurf.Copy()
			charSurf.SetSize(data.W, data.H)
			charSurf.Blit(c, 0, 0)
		case event.KeyDown:
			data := evt.Data.(event.KeyData)
			k := data.Key
			evtLogs = append(evtLogs, fmt.Sprintf("keydown: %s", k))
			if k.Rune != '\u0000' {
				if k.ShiftRune != '\u0000' && (data.Mod[key.LShift] || data.Mod[key.RShift]) {
					newChars = append(newChars, k.ShiftRune)
				} else {
					newChars = append(newChars, k.Rune)
				}
			}
			switch k {
			// case key.R:
			// 	ggweb.Log("Register events")
			// 	ggweb.RegisterEvents(display)
			// case key.U:
			// 	ggweb.Log("Unregister events")
			// 	ggweb.UnregisterEvents(display)
			}
		case event.KeyUp:
			data := evt.Data.(event.KeyData)
			k := data.Key
			evtLogs = append(evtLogs, fmt.Sprintf("keyup: %s", k))
		case event.MouseMotion:
			data := evt.Data.(event.MouseMotionData)
			evtLogs = append(evtLogs, fmt.Sprintf("mouse move: (%.0f, %.0f) (%.0f, %.0f)",
				data.Pos.X, data.Pos.Y, data.Rel.X, data.Rel.Y))
		case event.MouseButtonDown:
			data := evt.Data.(event.MouseData)
			evtLogs = append(evtLogs, fmt.Sprintf("mouse down: %d (%.0f, %.0f)",
				data.Button, data.Pos.X, data.Pos.Y))
		case event.MouseButtonUp:
			data := evt.Data.(event.MouseData)
			evtLogs = append(evtLogs, fmt.Sprintf("mouse up: %d (%.0f, %.0f)",
				data.Button, data.Pos.X, data.Pos.Y))
		case event.MouseWheel:
			data := evt.Data.(event.MouseWheelData)
			evtLogs = append(evtLogs, fmt.Sprintf("mouse wheel: (%.1f, %.1f %.1f)",
				data.Dx, data.Dy, data.Dz))
		}
	}

	if len(evtLogs) > 0 {
		// We need to resize evtSurf if any of the new logs would be too wide
		newEvtSurf := ggweb.NewSurface(0, 0)
		newEvtSurf.SetFont(&font2)
		maxWidth := 0.0
		for _, e := range evtLogs {
			w := newEvtSurf.TextWidth(e)
			if w > maxWidth {
				maxWidth = w
			}
		}

		newEvtSurf.SetSize(int(math.Max(maxWidth, evtSurf.Rect().W)), len(evtLogs)*int(font2.Size))
		// Reapply font since it is lost with SetSize
		newEvtSurf.SetFont(&font2)
		newEvtSurf.StyleColor(ggweb.Fill, color.White)
		newEvtSurf.SetTextAlign(ggweb.TextAlignLeft)
		newEvtSurf.SetTextBaseline(ggweb.TextBaselineTop)
		for i, e := range evtLogs {
			newEvtSurf.DrawText(ggweb.Fill, e, 0, float64(i)*font2.Size)
		}

		evtSurfCopy := evtSurf.Copy()
		if newEvtSurf.Rect().W > evtSurf.Rect().W {
			evtSurf.SetSize(int(newEvtSurf.Rect().W), int(evtSurf.Rect().H))
		}
		evtSurf.ClearRect(evtSurf.Rect())
		evtSurf.Blit(newEvtSurf, 0, 0)
		evtSurf.Blit(evtSurfCopy, 0, newEvtSurf.Rect().H)
	}

	charSurf.StyleColor(ggweb.Fill, color.RGBA{0, 0, 0, 5})
	charSurf.DrawRect(ggweb.Fill, charSurf.Rect())

	charSurf.StyleColor(ggweb.Fill, color.White)
	charSurf.SetFont(&font1)
	charSurf.SetTextAlign(ggweb.TextAlignLeft)
	charSurf.SetTextBaseline(ggweb.TextBaselineTop)
	w := charSurf.TextWidth(" ")
	for _, r := range newChars {
		c := string(r)
		switch r {
		case key.Enter.Rune, key.NpEnter.Rune:
			charY += font1.Size + charPadding
			charX = padding
		case key.Backspace.Rune:
			charX = charX - w - charPadding
			if charX < 0 {
				charX = 0
				charY = math.Max(padding, charY-font1.Size-charPadding)
			}
		default:
			charSurf.DrawText(ggweb.Fill, c, charX, charY)
			charX += w + charPadding
			if charX > charSurf.Rect().W {
				charY += font1.Size + charPadding
				charX = padding
			}
		}
	}

	display.Blit(charSurf, 0, 0)

	r := evtSurf.Rect()
	r.Move(display.Rect().W-r.W, display.Rect().H-r.H)
	display.Blit(evtSurf, r.X, r.Y)

	display.StyleColor(ggweb.Stroke, color.White)
	display.SetLineWidth(2)
	display.DrawLine(charX, charY, charX, charY+font1.Size)
}
