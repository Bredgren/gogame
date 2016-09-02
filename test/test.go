// To run the tests run "gopherjs serve $GOPATH/github.com/Bredgren/gogame/test" and navigate
// to http://localhost:8080/github.com/Bredgren/gogame/test/ in your browser.
package main

import (
	"fmt"
	"math"
	"time"
	"unicode"

	"github.com/Bredgren/gogame"
	"github.com/Bredgren/gogame/event"
	"github.com/Bredgren/gogame/geo"
	"github.com/Bredgren/gogame/key"
	"github.com/gopherjs/jquery"
)

var jq = jquery.NewJQuery

var resultList jquery.JQuery

type result struct {
	TestName string
	Errors   []string
}

func main() {
	resultList = jq("#test-results")

	ready := gogame.Ready()
	go func() {
		<-ready
		testCanvas()
	}()
}

func appendResultSection(sectionName string, results []*result) {
	section := jq("<li>").AddClass("result-section")
	section.Append(jq("<h1>").SetText(sectionName))
	sectionRes := jq("<ul>").AddClass("result-section-list")
	section.Append(sectionRes)
	for _, result := range results {
		res := jq("<li>").AddClass("result")
		res.Append(jq("<h2>").SetText(result.TestName))
		if len(result.Errors) == 0 {
			res.AddClass("result-pass")
		} else {
			res.AddClass("result-fail")
			errors := jq("<ul>").AddClass("result-error-list")
			for _, e := range result.Errors {
				errors.Append(jq("<li>").AddClass("result-error").SetText(e))
			}
			res.Append(errors)
		}
		sectionRes.Append(res)
	}
	resultList.Append(section)
}

func testCanvas() {
	width, height := 900, 600
	display := gogame.MainDisplay()
	display.SetMode(width, height)
	display.Fill(gogame.FillBlack)

	// Test shapes and styles
	display.DrawRect(geo.Rect{X: 11, Y: 11, W: 48, H: 48}, &gogame.StrokeStyle{Colorer: gogame.White, Width: 4})
	display.DrawRect(geo.Rect{X: 70, Y: 10, W: 50, H: 50},
		&gogame.FillStyle{Colorer: &gogame.LinearGradient{
			X1: 0, Y1: 0, X2: 50, Y2: 50,
			ColorStops: []gogame.ColorStop{
				{0.0, gogame.Red},
				{1.0, gogame.Blue},
			},
		}})
	display.DrawRect(geo.Rect{X: 130, Y: 10, W: 50, H: 50},
		&gogame.FillStyle{Colorer: &gogame.RadialGradient{
			X1: 50 / 2, Y1: 50 / 2, R1: 40, X2: 50 / 4, Y2: 50 / 4, R2: 1,
			ColorStops: []gogame.ColorStop{
				{0.0, gogame.Blue},
				{1.0, gogame.Green},
			},
		}})

	s := gogame.NewSurface(10, 10)
	s.DrawRect(geo.Rect{X: 0, Y: 0, W: 10, H: 10},
		&gogame.FillStyle{Colorer: &gogame.LinearGradient{
			X1: 0, Y1: 10, X2: 10, Y2: 0,
			ColorStops: []gogame.ColorStop{
				{0.0, gogame.Red},
				{1.0, gogame.Green},
			},
		}})

	pattern := gogame.Pattern{
		Source: s,
		Type:   gogame.Repeat,
	}

	display.DrawRect(geo.Rect{X: 190, Y: 10, W: 50, H: 50},
		&gogame.FillStyle{Colorer: &pattern})

	display.Blit(s, 30, 30)

	display.DrawCircle(35, 95, 20, &gogame.StrokeStyle{
		Colorer: &pattern,
		Width:   10,
	})

	display.DrawEllipse(geo.Rect{X: 70, Y: 70, W: 110, H: 50}, &gogame.StrokeStyle{
		Colorer: gogame.White,
		Width:   2,
	})

	display.DrawArc(geo.Rect{X: 190, Y: 70, W: 110, H: 50}, 0, 0.75*math.Pi, &gogame.StrokeStyle{
		Colorer: &gogame.Color{R: 1.0, G: 1.0, A: 1.0},
		Width:   2,
	})

	display.DrawLine(10, 130, 60, 180, &gogame.StrokeStyle{
		Colorer: gogame.White,
		Width:   2,
	})

	display.DrawLine(70, 130, 70, 180, &gogame.StrokeStyle{
		Colorer: gogame.White,
		Cap:     gogame.LineCapRound,
		Width:   1,
	})

	display.DrawLine(75.5, 130, 75.5, 180, &gogame.StrokeStyle{
		Colorer: gogame.White,
		Cap:     gogame.LineCapSquare,
		Width:   1,
	})

	display.DrawLine(80, 130, 80, 180, &gogame.StrokeStyle{
		Colorer: gogame.White,
		Cap:     gogame.LineCapRound,
		Width:   2,
	})

	display.DrawLine(85.5, 130, 85.5, 180, &gogame.StrokeStyle{
		Colorer: gogame.White,
		Cap:     gogame.LineCapSquare,
		Width:   2,
	})

	display.DrawLine(100, 130, 200, 130, &gogame.StrokeStyle{
		Colorer: gogame.White,
		Width:   2,
		Dash:    []float64{5, 1, 4, 2, 3, 3, 2, 4, 1, 5},
	})

	display.DrawLines([][2]float64{{210, 130}, {260, 130}, {210, 180}}, &gogame.StrokeStyle{
		Colorer: gogame.White,
		Join:    gogame.LineJoinRound,
		Cap:     gogame.LineCapRound,
		Width:   6,
	})

	display.DrawLines([][2]float64{{270, 130}, {320, 130}, {270, 180}}, &gogame.StrokeStyle{
		Colorer: gogame.White,
		Join:    gogame.LineJoinMiter,
		Cap:     gogame.LineCapSquare,
		Width:   6,
	})

	display.DrawLines([][2]float64{{330, 130}, {380, 130}, {330, 180}}, &gogame.StrokeStyle{
		Colorer: gogame.White,
		Join:    gogame.LineJoinBevel,
		Cap:     gogame.LineCapButt,
		Width:   6,
	})

	display.DrawLines([][2]float64{{390, 130}, {440, 130}, {390, 180}}, gogame.FillWhite)

	// Test text
	font := gogame.Font{
		Size: 50,
	}
	display.DrawText("Hello", 10, 190, &font, &gogame.TextStyle{
		Colorer:   gogame.White,
		Type:      gogame.Stroke,
		Baseline:  gogame.TextBaselineHanging,
		LineWidth: 2,
	})
	w := font.Width("Hello", &gogame.TextStyle{})
	display.DrawText("World!", float64(10+w), 200, &font, &gogame.TextStyle{
		Colorer: &gogame.LinearGradient{
			X1: 0, Y1: 0, X2: float64(w), Y2: 0,
			ColorStops: []gogame.ColorStop{
				{0.0, gogame.Blue},
				{1.0, gogame.Green},
			},
		},
		Type:     gogame.Fill,
		Baseline: gogame.TextBaselineMiddle,
	})

	grid := gogame.NewSurface(50, 50)
	grid.DrawRect(geo.Rect{X: 0, Y: 0, W: 50, H: 50},
		&gogame.FillStyle{Colorer: &pattern})

	display.Blit(grid, 10, 250)

	scaledGrid := grid.Scaled(2, 2)
	display.Blit(scaledGrid, 70, 250)

	rotatedGrid := grid.Rotated(0.12 * math.Pi)
	display.Blit(rotatedGrid, 180, 250)

	rotatedScaledGrid := scaledGrid.Rotated(0.25 * math.Pi)
	display.Blit(rotatedScaledGrid, 250, 250)

	font = gogame.Font{
		Size:   25,
		Family: gogame.FontFamily("courier new, monospace"),
	}
	textStyle := gogame.TextStyle{
		Colorer:  gogame.Black,
		Type:     gogame.Fill,
		Align:    gogame.TextAlignRight,
		Baseline: gogame.TextBaselineHanging,
	}
	text := font.Render("Text surface", &textStyle, &gogame.FillStyle{
		Colorer: gogame.Color{R: 0.8, G: 0.8, B: 1.0, A: 0.7},
	})
	display.Blit(text, 10, 360)

	rotatedText := text.Rotated(0.75 * math.Pi)
	display.Blit(rotatedText, 160, 330)

	// Test nil styles
	gogame.DefaultColor = gogame.Red
	var style *gogame.FillStyle
	display.DrawRect(geo.Rect{X: 250, Y: 10, W: 10, H: 10}, style)
	display.DrawCircle(270, 15, 5, nil)
	display.DrawLine(250, 25, 260, 35, nil)
	display.DrawText("nil", 265, 30, nil, nil)
	text = font.Render("nil", nil, nil)
	display.Blit(text, 250, 40)
	var nilFont *gogame.Font
	text = nilFont.Render("nil2", nil, nil)
	display.Blit(text, 280, 40)

	// Test SetAt
	for x := 10; x < 110; x++ {
		p := float64(x-10) / 100
		c := gogame.Color{R: 1.0 - p, B: p, A: 1.0}
		display.SetAt(x, 400, c)
	}

	// Test clipping
	display.SetClip(geo.Rect{X: 400, Y: 10, W: 100, H: 100})
	display.Fill(gogame.FillWhite)
	display.DrawCircle(500, 110, 20, &gogame.FillStyle{})
	display.ClearClip()
	display.DrawCircle(500, 60, 20, &gogame.FillStyle{})

	// Test copy
	copy := display.Copy()
	copy = display.Scaled(0.2, 0.2)
	copy.DrawRect(copy.Rect(), &gogame.StrokeStyle{Colorer: gogame.White, Width: 2})
	display.Blit(copy, float64(display.Width()-copy.Width()), float64(display.Height()-copy.Height()))

	// Test curves
	display.SetAt(590, 20, gogame.Green)
	display.DrawQuadraticCurve(550, 10, 550, 60, 590, 20, &gogame.StrokeStyle{
		Colorer: gogame.White,
		Width:   5,
	})

	display.DrawQuadraticCurves([][2]float64{
		{620, 30}, {630, 0},
		{640, 30}, {670, 40},
		{640, 50}, {630, 80},
		{620, 50}, {590, 40},
		{620, 30},
	}, gogame.FillWhite)

	display.DrawQuadraticCurves([][2]float64{
		{620, 130}, {630, 100},
		{640, 130}, {670, 140},
		{640, 150}, {630, 180},
		{620, 150}, {590, 140},
		{620, 130}, {0, 0}, // Last point ignored
	}, gogame.FillWhite)

	display.SetAt(695, 40, gogame.Green)
	display.SetAt(695, 0, gogame.Green)
	display.DrawBezierCurve(670, 20, 720, 20, 695, 40, 695, 0, &gogame.StrokeStyle{
		Colorer: gogame.White,
		Width:   5,
	})

	display.DrawBezierCurves([][2]float64{
		{750, 20}, {775, 40}, {775, 0},
		{800, 20}, {780, 45}, {820, 45},
		{800, 70}, {775, 50}, {775, 90},
		{750, 70}, {770, 45}, {730, 45},
		{750, 20},
	}, gogame.FillWhite)

	display.DrawBezierCurves([][2]float64{
		{750, 120}, {775, 140}, {775, 100},
		{800, 120}, {780, 145}, {820, 145},
		{800, 170}, {775, 150}, {775, 190},
		{750, 170}, {770, 145}, {730, 145},
		{750, 120}, {0, 0}, {0, 0}, // Last two ignored
	}, gogame.FillWhite)

	display.Flip()

	eventSurf := gogame.NewSurface(300, 100)
	eventSurf.Fill(gogame.FillBlack)
	eventFont := gogame.Font{
		Size:   eventSurf.Height() / 7,
		Family: gogame.FontFamilySansSerif,
	}
	eventStyle := gogame.TextStyle{
		Colorer: gogame.White,
		Type:    gogame.Fill,
	}

	fpsStyle := gogame.TextStyle{
		Colorer: gogame.White,
		Type:    gogame.Fill,
	}
	maxFpsWidth := 0

	done := make(chan struct{}, 1)
	go func() {
		<-done
		display.Fill(gogame.FillWhite)
		display.Update([]geo.Rect{
			{X: 500, Y: 500, W: 10, H: 10},
			{X: 520, Y: 505, W: 15, H: 10},
		})
	}()

	gogame.Log("start main loop")
	gogame.SetMainLoop(func(t time.Duration) {
		for evt := event.Poll(); evt.Type != event.NoEvent; evt = event.Poll() {
			msg := ""
			switch evt.Type {
			case event.Quit:
				msg = "quit"
				gogame.UnsetMainLoop()
				gogame.Log("quit")
				done <- struct{}{}
			case event.VideoResize:
				data := evt.Data.(event.ResizeData)
				msg = fmt.Sprintf("resize: (%d, %d)", data.W, data.H)
			case event.KeyDown:
				data := evt.Data.(event.KeyData)
				k := data.Key
				char := k.Rune()
				if data.Mod[key.LShift] || data.Mod[key.RShift] {
					char = unicode.ToUpper(char)
				}
				msg = fmt.Sprintf("keydown: %s %s", k, string(char))
			case event.KeyUp:
				data := evt.Data.(event.KeyData)
				k := data.Key
				char := k.Rune()
				if data.Mod[key.LShift] || data.Mod[key.RShift] {
					char = unicode.ToUpper(char)
				}
				msg = fmt.Sprintf("keyup: %s %s", k, string(char))
				switch k {
				case key.Escape:
					msg = "quit (by escape)"
					gogame.UnsetMainLoop()
					done <- struct{}{}
					// case key.F:
					// 	gogame.SetFullscreen(!gogame.GetFullscreen())
				}
			case event.MouseButtonDown:
				data := evt.Data.(event.MouseData)
				msg = fmt.Sprintf("mousedown: (%.0f, %.0f) %d", data.Pos.X, data.Pos.Y, data.Button)
			case event.MouseButtonUp:
				data := evt.Data.(event.MouseData)
				msg = fmt.Sprintf("mouseup: (%.0f, %.0f) %d", data.Pos.X, data.Pos.Y, data.Button)
			case event.MouseMotion:
				data := evt.Data.(event.MouseMotionData)
				msg = fmt.Sprintf("mousemove: (%.0f, %.0f) (%.0f, %.0f) %v", data.Pos.X, data.Pos.Y, data.Rel.Dx, data.Rel.Dy, data.Buttons)
			}
			text := eventFont.Render(msg, &eventStyle, gogame.FillBlack)
			eventSurf.Blit(eventSurf, 0, float64(text.Height()))
			eventSurf.DrawRect(geo.Rect{X: 0, Y: 0, W: float64(eventSurf.Width()), H: float64(text.Height())}, gogame.FillBlack)
			eventSurf.Blit(text, 0, 0)
		}

		display.Blit(eventSurf, 0, float64(display.Height()-eventSurf.Height()))

		text := eventFont.Render(gogame.Stats.LoopDuration.String(), &fpsStyle, gogame.FillBlack)
		if text.Width() > maxFpsWidth {
			maxFpsWidth = text.Width()
		}
		r := text.Rect()
		r.W = float64(maxFpsWidth)
		r.Y = float64(display.Height() - eventSurf.Height() - text.Height())
		display.DrawRect(r, gogame.FillBlack)
		display.Blit(text, 0, r.Y)

		// display.Flip()
		display.Update([]geo.Rect{{X: r.X, Y: r.Y, W: math.Max(r.W, eventSurf.Rect().W), H: r.H + eventSurf.Rect().H}})
	})
}
