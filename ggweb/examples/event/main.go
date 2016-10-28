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

	charSurf = ggweb.NewSurface(width, height)
	evtSurf = ggweb.NewSurface(100, 150)

	// Prevent scrolling
	ggweb.PreventKeyDefault[key.Space] = true
	ggweb.PreventKeyDefault[key.Tab] = true

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
		case event.Quit:
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
			evtLogs = append(evtLogs, fmt.Sprintf("keydown: %s", k.Name))
			if k.Rune != '\u0000' {
				if k.ShiftRune != '\u0000' && (data.Mod[key.LShift] || data.Mod[key.RShift]) {
					newChars = append(newChars, k.ShiftRune)
				} else {
					newChars = append(newChars, k.Rune)
				}
			}
		case event.KeyUp:
			data := evt.Data.(event.KeyData)
			k := data.Key
			evtLogs = append(evtLogs, fmt.Sprintf("keyup: %s", k.Name))
		case event.MouseMotion:
		case event.MouseButtonDown:
		case event.MouseButtonUp:
		case event.MouseWheel:
		}
	}

	logSurf := ggweb.NewSurface(0, 0)
	logSurf.StyleColor(ggweb.Fill, color.White)
	logSurf.SetFont(&font2)
	logSurf.SetTextAlign(ggweb.TextAlignLeft)
	logSurf.SetTextBaseline(ggweb.TextBaselineTop)
	maxWidth := 0.0
	for _, e := range evtLogs {
		w := logSurf.TextWidth(e)
		ggweb.Log(e, w)
		if w > maxWidth {
			maxWidth = w
		}
	}
	logSurf.SetSize(int(math.Max(maxWidth, evtSurf.Rect().W)), len(evtLogs)*int(font2.Size))
	for i, e := range evtLogs {
		logSurf.DrawText(ggweb.Fill, e, 0, float64(i)*font2.Size)
		display.StyleColor(ggweb.Fill, color.White)
		display.SetFont(&font2)
		display.SetTextAlign(ggweb.TextAlignLeft)
		display.SetTextBaseline(ggweb.TextBaselineTop)
		display.DrawText(ggweb.Fill, e, 20, 20)
	}

	if len(evtLogs) > 0 {
		evtSurfCopy := evtSurf.Copy()
		if logSurf.Rect().W > evtSurf.Rect().W {
			evtSurf.SetSize(int(logSurf.Rect().W), int(evtSurf.Rect().H))
		}
		evtSurf.Blit(logSurf, 0, 0)
		evtSurf.Blit(evtSurfCopy, 0, logSurf.Rect().H)
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
		}
	}

	display.Blit(charSurf, 0, 0)
	display.Blit(logSurf, display.Rect().W-evtSurf.Rect().W, display.Rect().H-evtSurf.Rect().H)

	display.StyleColor(ggweb.Stroke, color.White)
	display.SetLineWidth(2)
	display.DrawLine(charX, charY, charX, charY+font1.Size)
}
