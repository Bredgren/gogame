// To run the tests run "gopherjs serve $GOPATH/github.com/Bredgren/gogame/test" and navigate
// to http://localhost:8080/github.com/Bredgren/gogame/test/ in your browser.
package main

import (
	"image/color"

	"github.com/Bredgren/gogame/geo"
	"github.com/Bredgren/gogame/ggweb"
)

func main() {
	ggweb.Init(testCanvas)
}

func testCanvas() {
	width, height := 900, 600
	display := ggweb.NewSurfaceFromID("main")
	display.SetSize(width, height)

	display.StyleColor(ggweb.Fill, color.Gray{})
	display.DrawRect(ggweb.Fill, display.Rect())

	display.StyleColor(ggweb.Fill, color.Gray{255})
	display.DrawRect(ggweb.Fill, geo.Rect{X: 10, Y: 10, W: 20, H: 20})

	display.StyleColor(ggweb.Fill, color.RGBA{255, 0, 0, 255})
	display.DrawRect(ggweb.Fill, geo.Rect{X: 20, Y: 20, W: 20, H: 20})

	display.StyleColor(ggweb.Fill, color.RGBA{0, 0, 255, 150})
	display.DrawRect(ggweb.Fill, geo.Rect{X: 25, Y: 25, W: 20, H: 20})

	// 	// Test localstorage
	// 	v1, ok1 := gogame.LocalStorageGet("key")
	// 	gogame.Log("Get invalid key from storage. v1 is", v1, "ok1 is", ok1)
	// 	gogame.LocalStorageSet("key", "value")
	// 	v2, ok2 := gogame.LocalStorageGet("key")
	// 	gogame.Log("Get valid key from storage. v2 is", v2, "ok2 is", ok2)
	// 	gogame.LocalStorageRemove("key")
	// 	v3, ok3 := gogame.LocalStorageGet("key")
	// 	gogame.Log("Get removed key from storage. v3 is", v3, "ok3 is", ok3)

	// 	// Test shapes and styles
	// 	display.DrawRect(geo.Rect{X: 11, Y: 11, W: 48, H: 48}, &gogame.StrokeStyle{Colorer: gogame.White, Width: 4})
	// 	display.DrawRect(geo.Rect{X: 70, Y: 10, W: 50, H: 50},
	// 		&gogame.FillStyle{Colorer: &gogame.LinearGradient{
	// 			X1: 0, Y1: 0, X2: 50, Y2: 50,
	// 			ColorStops: []gogame.ColorStop{
	// 				{Position: 0.0, Colorer: gogame.Red},
	// 				{Position: 1.0, Colorer: gogame.Blue},
	// 			},
	// 		}})
	// 	display.DrawRect(geo.Rect{X: 130, Y: 10, W: 50, H: 50},
	// 		&gogame.FillStyle{Colorer: &gogame.RadialGradient{
	// 			X1: 50 / 2, Y1: 50 / 2, R1: 40, X2: 50 / 4, Y2: 50 / 4, R2: 1,
	// 			ColorStops: []gogame.ColorStop{
	// 				{Position: 0.0, Colorer: gogame.Blue},
	// 				{Position: 1.0, Colorer: gogame.Green},
	// 			},
	// 		}})

	// 	s := gogame.NewSurface(10, 10)
	// 	s.DrawRect(geo.Rect{X: 0, Y: 0, W: 10, H: 10},
	// 		&gogame.FillStyle{Colorer: &gogame.LinearGradient{
	// 			X1: 0, Y1: 10, X2: 10, Y2: 0,
	// 			ColorStops: []gogame.ColorStop{
	// 				{Position: 0.0, Colorer: gogame.Red},
	// 				{Position: 1.0, Colorer: gogame.Green},
	// 			},
	// 		}})

	// 	pattern := gogame.Pattern{
	// 		Source: s,
	// 		Type:   gogame.Repeat,
	// 	}

	// 	display.DrawRect(geo.Rect{X: 190, Y: 10, W: 50, H: 50},
	// 		&gogame.FillStyle{Colorer: &pattern})

	// 	display.Blit(s, 30, 30)

	// 	display.DrawCircle(35, 95, 20, &gogame.StrokeStyle{
	// 		Colorer: &pattern,
	// 		Width:   10,
	// 	})

	// 	display.DrawEllipse(geo.Rect{X: 70, Y: 70, W: 110, H: 50}, &gogame.StrokeStyle{
	// 		Colorer: gogame.White,
	// 		Width:   2,
	// 	})

	// 	display.DrawArc(geo.Rect{X: 190, Y: 70, W: 110, H: 50}, 0, 0.75*math.Pi, &gogame.StrokeStyle{
	// 		Colorer: &gogame.Color{R: 1.0, G: 1.0, A: 1.0},
	// 		Width:   2,
	// 	})

	// 	display.DrawLine(10, 130, 60, 180, &gogame.StrokeStyle{
	// 		Colorer: gogame.White,
	// 		Width:   2,
	// 	})

	// 	display.DrawLine(70, 130, 70, 180, &gogame.StrokeStyle{
	// 		Colorer: gogame.White,
	// 		Cap:     gogame.LineCapRound,
	// 		Width:   1,
	// 	})

	// 	display.DrawLine(75.5, 130, 75.5, 180, &gogame.StrokeStyle{
	// 		Colorer: gogame.White,
	// 		Cap:     gogame.LineCapSquare,
	// 		Width:   1,
	// 	})

	// 	display.DrawLine(80, 130, 80, 180, &gogame.StrokeStyle{
	// 		Colorer: gogame.White,
	// 		Cap:     gogame.LineCapRound,
	// 		Width:   2,
	// 	})

	// 	display.DrawLine(85.5, 130, 85.5, 180, &gogame.StrokeStyle{
	// 		Colorer: gogame.White,
	// 		Cap:     gogame.LineCapSquare,
	// 		Width:   2,
	// 	})

	// 	display.DrawLine(100, 130, 200, 130, &gogame.StrokeStyle{
	// 		Colorer: gogame.White,
	// 		Width:   2,
	// 		Dash:    []float64{5, 1, 4, 2, 3, 3, 2, 4, 1, 5},
	// 	})

	// 	display.DrawLines([][2]float64{{210, 130}, {260, 130}, {210, 180}}, &gogame.StrokeStyle{
	// 		Colorer: gogame.White,
	// 		Join:    gogame.LineJoinRound,
	// 		Cap:     gogame.LineCapRound,
	// 		Width:   6,
	// 	})

	// 	display.DrawLines([][2]float64{{270, 130}, {320, 130}, {270, 180}}, &gogame.StrokeStyle{
	// 		Colorer: gogame.White,
	// 		Join:    gogame.LineJoinMiter,
	// 		Cap:     gogame.LineCapSquare,
	// 		Width:   6,
	// 	})

	// 	display.DrawLines([][2]float64{{330, 130}, {380, 130}, {330, 180}}, &gogame.StrokeStyle{
	// 		Colorer: gogame.White,
	// 		Join:    gogame.LineJoinBevel,
	// 		Cap:     gogame.LineCapButt,
	// 		Width:   6,
	// 	})

	// 	display.DrawLines([][2]float64{{390, 130}, {440, 130}, {390, 180}}, gogame.FillWhite)

	// 	// Test text
	// 	font := gogame.Font{
	// 		Size: 50,
	// 	}
	// 	display.DrawText("Hello", 10, 190, &font, &gogame.TextStyle{
	// 		Colorer:   gogame.White,
	// 		Type:      gogame.Stroke,
	// 		Baseline:  gogame.TextBaselineHanging,
	// 		LineWidth: 2,
	// 	})
	// 	w := font.Width("Hello", &gogame.TextStyle{})
	// 	display.DrawText("World!", float64(10+w), 200, &font, &gogame.TextStyle{
	// 		Colorer: &gogame.LinearGradient{
	// 			X1: 0, Y1: 0, X2: float64(w), Y2: 0,
	// 			ColorStops: []gogame.ColorStop{
	// 				{Position: 0.0, Colorer: gogame.Blue},
	// 				{Position: 1.0, Colorer: gogame.Green},
	// 			},
	// 		},
	// 		Type:     gogame.Fill,
	// 		Baseline: gogame.TextBaselineMiddle,
	// 	})

	// 	grid := gogame.NewSurface(50, 50)
	// 	grid.DrawRect(geo.Rect{X: 0, Y: 0, W: 50, H: 50},
	// 		&gogame.FillStyle{Colorer: &pattern})

	// 	display.Blit(grid, 10, 250)

	// 	scaledGrid := grid.Scaled(2, 2)
	// 	display.Blit(scaledGrid, 70, 250)

	// 	rotatedGrid := grid.Rotated(0.12 * math.Pi)
	// 	display.Blit(rotatedGrid, 180, 250)

	// 	rotatedScaledGrid := scaledGrid.Rotated(0.25 * math.Pi)
	// 	display.Blit(rotatedScaledGrid, 250, 250)

	// 	font = gogame.Font{
	// 		Size:   25,
	// 		Family: gogame.FontFamily("courier new, monospace"),
	// 	}
	// 	textStyle := gogame.TextStyle{
	// 		Colorer:  gogame.Black,
	// 		Type:     gogame.Fill,
	// 		Align:    gogame.TextAlignRight,
	// 		Baseline: gogame.TextBaselineHanging,
	// 	}
	// 	text := font.Render("Text surface", &textStyle, &gogame.FillStyle{
	// 		Colorer: gogame.Color{R: 0.8, G: 0.8, B: 1.0, A: 0.7},
	// 	})
	// 	display.Blit(text, 10, 360)

	// 	rotatedText := text.Rotated(0.75 * math.Pi)
	// 	display.Blit(rotatedText, 160, 330)

	// 	// Test nil styles
	// 	gogame.DefaultColor = gogame.Red
	// 	var style *gogame.FillStyle
	// 	display.DrawRect(geo.Rect{X: 250, Y: 10, W: 10, H: 10}, style)
	// 	display.DrawCircle(270, 15, 5, nil)
	// 	display.DrawLine(250, 25, 260, 35, nil)
	// 	display.DrawText("nil", 265, 30, nil, nil)
	// 	text = font.Render("nil", nil, nil)
	// 	display.Blit(text, 250, 40)
	// 	var nilFont *gogame.Font
	// 	text = nilFont.Render("nil2", nil, nil)
	// 	display.Blit(text, 280, 40)

	// 	// Test SetAt
	// 	for x := 10; x < 110; x++ {
	// 		p := float64(x-10) / 100
	// 		c := gogame.Color{R: 1.0 - p, B: p, A: 1.0}
	// 		display.SetAt(x, 400, c)
	// 	}

	// 	// Test clipping
	// 	display.SetClip(geo.Rect{X: 400, Y: 10, W: 100, H: 100})
	// 	display.Fill(gogame.FillWhite)
	// 	display.DrawCircle(500, 110, 20, &gogame.FillStyle{})
	// 	display.ClearClip()
	// 	display.DrawCircle(500, 60, 20, &gogame.FillStyle{})

	// 	// Test copy
	// 	copy := display.Copy()
	// 	copy = copy.Scaled(0.4, 0.2)
	// 	copy.DrawRect(copy.Rect(), &gogame.StrokeStyle{Colorer: gogame.White, Width: 2})
	// 	display.Blit(copy, float64(display.Width()-copy.Width()), float64(display.Height()-copy.Height()))

	// 	// Test curves
	// 	display.SetAt(590, 20, gogame.Color{G: 1, A: 1})
	// 	display.DrawQuadraticCurve(550, 10, 550, 60, 590, 20, &gogame.StrokeStyle{
	// 		Colorer: gogame.White,
	// 		Width:   5,
	// 	})

	// 	display.DrawQuadraticCurves([][2]float64{
	// 		{620, 30}, {630, 0},
	// 		{640, 30}, {670, 40},
	// 		{640, 50}, {630, 80},
	// 		{620, 50}, {590, 40},
	// 		{620, 30},
	// 	}, gogame.FillWhite)

	// 	display.DrawQuadraticCurves([][2]float64{
	// 		{620, 130}, {630, 100},
	// 		{640, 130}, {670, 140},
	// 		{640, 150}, {630, 180},
	// 		{620, 150}, {590, 140},
	// 		{620, 130}, {0, 0}, // Last point ignored
	// 	}, gogame.FillWhite)

	// 	display.SetAt(695, 40, gogame.Color{G: 1, A: 1})
	// 	display.SetAt(695, 0, gogame.Color{G: 1, A: 1})
	// 	display.DrawBezierCurve(670, 20, 720, 20, 695, 40, 695, 0, &gogame.StrokeStyle{
	// 		Colorer: gogame.White,
	// 		Width:   5,
	// 	})

	// 	display.DrawBezierCurves([][2]float64{
	// 		{750, 20}, {775, 40}, {775, 0},
	// 		{800, 20}, {780, 45}, {820, 45},
	// 		{800, 70}, {775, 50}, {775, 90},
	// 		{750, 70}, {770, 45}, {730, 45},
	// 		{750, 20},
	// 	}, gogame.FillWhite)

	// 	display.DrawBezierCurves([][2]float64{
	// 		{750, 120}, {775, 140}, {775, 100},
	// 		{800, 120}, {780, 145}, {820, 145},
	// 		{800, 170}, {775, 150}, {775, 190},
	// 		{750, 170}, {770, 145}, {730, 145},
	// 		{750, 120}, {0, 0}, {0, 0}, // Last two ignored
	// 	}, gogame.FillWhite)

	// 	// Test Image
	// 	img := gogame.LoadImage("img.png")
	// 	display.Blit(img, 420, 220)

	// 	f := gogame.Font{
	// 		Size: 20,
	// 	}
	// 	t := gogame.TextStyle{
	// 		Colorer:   gogame.White,
	// 		LineWidth: 2,
	// 	}
	// 	display.DrawText("Press 'c' to change cursor", 10, 440, &f, &t)
	// 	display.DrawText("Press 'p' to toggle music", 10, 460, &f, &t)
	// 	display.DrawText("Press 's' to play sound", 10, 480, &f, &t)

	// 	// Test composite opertaions
	// 	w, h := 100, 50
	// 	mask := gogame.NewSurface(w, h)
	// 	mask.DrawRect(geo.Rect{X: float64(h) / 2, Y: 0, W: float64(w) - float64(h), H: float64(h)}, gogame.FillWhite)
	// 	mask.DrawArc(geo.Rect{X: 0, Y: 0, W: float64(h), H: float64(h)}, math.Pi/2, 3*math.Pi/2, gogame.FillWhite)
	// 	mask.DrawArc(geo.Rect{X: float64(w - h), Y: 0, W: float64(h), H: float64(h)}, -math.Pi/2, math.Pi/2, gogame.FillWhite)

	// 	fill := gogame.NewSurface(w, h)
	// 	for x := 2; x < w; x += 5 {
	// 		fill.DrawLine(float64(x), 0, float64(x), float64(h), &gogame.StrokeStyle{
	// 			Colorer: gogame.Red,
	// 			Width:   2,
	// 		})
	// 	}
	// 	fill.BlitComp(mask, 0, 0, composite.DestinationIn)

	// 	fill.DrawLine(float64(h)/2, 1, float64(w)-float64(h)/2, 1, &gogame.StrokeStyle{
	// 		Colorer: gogame.Red,
	// 		Width:   2,
	// 	})
	// 	fill.DrawLine(float64(h)/2, float64(h)-1, float64(w)-float64(h)/2, float64(h)-1, &gogame.StrokeStyle{
	// 		Colorer: gogame.Red,
	// 		Width:   2,
	// 	})
	// 	fill.DrawArc(geo.Rect{X: 1, Y: 1, W: float64(h), H: float64(h) - 2}, math.Pi/2, 3*math.Pi/2, &gogame.StrokeStyle{
	// 		Colorer: gogame.Red,
	// 		Width:   2,
	// 	})
	// 	fill.DrawArc(geo.Rect{X: float64(w-h) - 1, Y: 1, W: float64(h), H: float64(h) - 2}, -math.Pi/2, math.Pi/2, &gogame.StrokeStyle{
	// 		Colorer: gogame.Red,
	// 		Width:   2,
	// 	})
	// 	display.Blit(fill, 450, 320)

	// 	testSubSurf()

	// 	display.Flip()

	// 	eventSurf := gogame.NewSurface(300, 100)
	// 	eventSurf.Fill(gogame.FillBlack)
	// 	eventFont := gogame.Font{
	// 		Size:   eventSurf.Height() / 7,
	// 		Family: gogame.FontFamilySansSerif,
	// 	}
	// 	eventStyle := gogame.TextStyle{
	// 		Colorer: gogame.White,
	// 		Type:    gogame.Fill,
	// 	}

	// 	fpsStyle := gogame.TextStyle{
	// 		Colorer: gogame.White,
	// 		Type:    gogame.Fill,
	// 	}
	// 	maxFpsWidth := 0

	// 	done := make(chan struct{}, 1)
	// 	go func() {
	// 		<-done
	// 		display.Fill(gogame.FillWhite)
	// 		display.Update([]geo.Rect{
	// 			{X: 500, Y: 500, W: 10, H: 10},
	// 			{X: 520, Y: 505, W: 15, H: 10},
	// 		})
	// 	}()

	// 	music := sound.New("Voice Over Under.mp3")
	// 	music.SetLoop(true)
	// 	music.SetVolume(0.2)
	// 	music.Pause()
	// 	sfx := sound.New("pop1.wav")
	// 	sfx.SetVolume(0.3)

	// 	gogame.Log("start main loop")
	// 	gogame.SetMainLoop(func(t time.Duration) {
	// 		for evt := event.Poll(); evt.Type != event.NoEvent; evt = event.Poll() {
	// 			msg := ""
	// 			switch evt.Type {
	// 			case event.Quit:
	// 				msg = "quit"
	// 				gogame.UnsetMainLoop()
	// 				gogame.Log("quit")
	// 				done <- struct{}{}
	// 			case event.VideoResize:
	// 				data := evt.Data.(event.ResizeData)
	// 				msg = fmt.Sprintf("resize: (%d, %d)", data.W, data.H)
	// 			case event.KeyDown:
	// 				data := evt.Data.(event.KeyData)
	// 				k := data.Key
	// 				char := k.Rune()
	// 				if data.Mod[key.LShift] || data.Mod[key.RShift] {
	// 					char = unicode.ToUpper(char)
	// 				}
	// 				msg = fmt.Sprintf("keydown: %s %s", k, string(char))
	// 			case event.KeyUp:
	// 				data := evt.Data.(event.KeyData)
	// 				k := data.Key
	// 				char := k.Rune()
	// 				if data.Mod[key.LShift] || data.Mod[key.RShift] {
	// 					char = unicode.ToUpper(char)
	// 				}
	// 				msg = fmt.Sprintf("keyup: %s %s", k, string(char))
	// 				switch k {
	// 				case key.Escape:
	// 					msg = "quit (by escape)"
	// 					gogame.UnsetMainLoop()
	// 					done <- struct{}{}
	// 				case key.C:
	// 					if display.Cursor() == gogame.CursorDefault {
	// 						display.SetCursor(gogame.CursorProgress)
	// 					} else {
	// 						display.SetCursor(gogame.CursorDefault)

	// 					}
	// 				case key.P:
	// 					if music.Paused() {
	// 						music.Play()
	// 					} else {
	// 						music.Pause()
	// 					}
	// 				case key.S:
	// 					sfx.PlayFromStart()
	// 					// case key.F:
	// 					// 	gogame.SetFullscreen(!gogame.GetFullscreen())
	// 				}
	// 			case event.MouseButtonDown:
	// 				data := evt.Data.(event.MouseData)
	// 				msg = fmt.Sprintf("mousedown: (%.0f, %.0f) %d", data.Pos.X, data.Pos.Y, data.Button)
	// 			case event.MouseButtonUp:
	// 				data := evt.Data.(event.MouseData)
	// 				msg = fmt.Sprintf("mouseup: (%.0f, %.0f) %d", data.Pos.X, data.Pos.Y, data.Button)
	// 			case event.MouseMotion:
	// 				data := evt.Data.(event.MouseMotionData)
	// 				msg = fmt.Sprintf("mousemove: (%.0f, %.0f) (%.0f, %.0f) %v", data.Pos.X, data.Pos.Y, data.Rel.Dx, data.Rel.Dy, data.Buttons)
	// 			case event.MouseWheel:
	// 				data := evt.Data.(event.MouseWheelData)
	// 				msg = fmt.Sprintf("wheel: %f, %f, %f", data.Dx, data.Dy, data.Dz)
	// 			}
	// 			text := eventFont.Render(msg, &eventStyle, gogame.FillBlack)
	// 			eventSurf.Blit(eventSurf, 0, float64(text.Height()))
	// 			eventSurf.DrawRect(geo.Rect{X: 0, Y: 0, W: float64(eventSurf.Width()), H: float64(text.Height())}, gogame.FillBlack)
	// 			eventSurf.Blit(text, 0, 0)
	// 		}

	// 		display.Blit(eventSurf, 0, float64(display.Height()-eventSurf.Height()))

	// 		text := eventFont.Render(gogame.Stats.LoopDuration.String(), &fpsStyle, gogame.FillBlack)
	// 		if text.Width() > maxFpsWidth {
	// 			maxFpsWidth = text.Width()
	// 		}
	// 		r := text.Rect()
	// 		r.W = float64(maxFpsWidth)
	// 		r.Y = float64(display.Height() - eventSurf.Height() - text.Height())
	// 		display.DrawRect(r, gogame.FillBlack)
	// 		display.Blit(text, 0, r.Y)

	// 		// display.Flip()
	// 		display.Update([]geo.Rect{{X: r.X, Y: r.Y, W: math.Max(r.W, eventSurf.Rect().W), H: r.H + eventSurf.Rect().H}})
	// 	})
}

// func testSubSurf() {
// 	display := gogame.MainDisplay()
// 	sub := display.SubSurface(geo.Rect{X: 600, Y: 200, W: 20, H: 20})
// 	sub.DrawCircle(sub.Rect().CenterX(), sub.Rect().CenterY(), sub.Rect().W/2+2, gogame.FillWhite)

// 	display.Blit(sub, 650, 200)

// 	subCopy := sub.Copy()
// 	display.Blit(subCopy, 700, 200)

// 	// Clipping parent and subsurface doesn't work properly
// 	// r1 := geo.Rect{X: 600, Y: 250, W: 100, H: 75}
// 	// r2 := r1.Move(75, 50)
// 	// stroke := &gogame.StrokeStyle{Colorer: gogame.White, Width: 2}
// 	// display.DrawRect(r1, stroke)
// 	// display.DrawRect(r2, stroke)
// 	// display.SetClip(r1)
// 	// display.DrawCircle(r1.Left(), r1.Bottom(), 10, gogame.FillWhite)
// 	// sub2 := display.SubSurface(r2)
// 	// sub2.DrawCircle(0, 0, 10, gogame.FillWhite)
// 	// sub2.DrawCircle(r2.W, r2.H, 10, gogame.FillWhite)
// 	// display.DrawCircle(r1.Right(), r1.Top(), 10, gogame.FillWhite)
// 	// display.ClearClip()
// }
