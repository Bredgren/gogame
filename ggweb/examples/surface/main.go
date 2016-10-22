package main

import (
	"image/color"
	"math"

	"github.com/Bredgren/gogame/geo"
	"github.com/Bredgren/gogame/ggweb"
)

func main() {
	ggweb.Init(testSurface)
}

func testSurface() {
	width, height := 900, 600
	display := ggweb.NewSurfaceFromID("main")
	display.SetSize(width, height)

	Rect(display, 10, 10)
	Styles(display, 60, 10)
	CircleEllipseArc(display, 270, 10)
	Path(display, 400, 10)
	Transform(display, 600, 10)
	Text(display, 720, 10)
	Image(display, 10, 70)
	Pixel(display, 10, 320)
	Surface(display, 520, 70)
}

func Rect(display *ggweb.Surface, x, y float64) {
	display.StyleColor(ggweb.Fill, color.Black)
	display.DrawRect(ggweb.Fill, display.Rect())

	display.StyleColor(ggweb.Fill, color.RGBA{100, 100, 255, 255})
	display.DrawRect(ggweb.Fill, geo.Rect{X: x, Y: y, W: 40, H: 40})

	display.ClearRect(geo.Rect{X: x + 20, Y: y + 20, W: 25, H: 25})
}

func Styles(display *ggweb.Surface, x, y float64) {
	display.Save()

	display.StyleColor(ggweb.Fill, color.White)
	display.DrawRect(ggweb.Fill, geo.Rect{X: x, Y: y, W: 20, H: 20})

	display.StyleColor(ggweb.Fill, color.RGBA{255, 0, 0, 150})
	display.DrawRect(ggweb.Fill, geo.Rect{X: x + 5, Y: y + 5, W: 20, H: 20})

	display.StyleLinearGradient(ggweb.Fill, ggweb.LinearGradient{
		X1: x + 40, Y1: y, X2: x + 80, Y2: y + 50,
		ColorStops: []ggweb.ColorStop{
			{Position: 0, Color: color.RGBA{0, 255, 0, 255}},
			{Position: 1, Color: color.RGBA{0, 0, 255, 255}},
		},
	})
	display.DrawRect(ggweb.Fill, geo.Rect{X: x + 40, Y: y, W: 40, H: 40})
	display.StyleColor(ggweb.Stroke, color.White)
	display.DrawRect(ggweb.Stroke, geo.Rect{X: x + 40, Y: y, W: 40, H: 40})

	display.StyleRadialGradient(ggweb.Fill, ggweb.RadialGradient{
		X1: x + 110, Y1: y + 20, R1: 25, X2: x + 110, Y2: y + 20, R2: 1,
		ColorStops: []ggweb.ColorStop{
			{Position: 0, Color: color.RGBA{255, 0, 0, 255}},
			{Position: 0.5, Color: color.RGBA{0, 255, 0, 255}},
			{Position: 1, Color: color.RGBA{0, 0, 255, 255}},
		},
	})
	display.DrawRect(ggweb.Fill, geo.Rect{X: x + 90, Y: y, W: 40, H: 40})

	patSurf := ggweb.NewSurface(10, 10)
	patSurf.StyleLinearGradient(ggweb.Fill, ggweb.LinearGradient{
		X1: 0, Y1: 0, X2: 10, Y2: 10,
		ColorStops: []ggweb.ColorStop{
			{Position: 0, Color: color.RGBA{255, 0, 0, 255}},
			{Position: 1, Color: color.RGBA{0, 0, 255, 255}},
		},
	})
	patSurf.DrawRect(ggweb.Fill, patSurf.Rect())
	display.StylePattern(ggweb.Fill, ggweb.Pattern{
		Source: patSurf,
		Type:   ggweb.RepeatXY,
	})
	display.DrawRect(ggweb.Fill, geo.Rect{X: x + 140, Y: y, W: 50, H: 50})

	display.Restore()
}

func CircleEllipseArc(display *ggweb.Surface, x, y float64) {
	display.Save()

	display.StyleColor(ggweb.Fill, color.RGBA{100, 255, 100, 255})
	display.StyleColor(ggweb.Stroke, color.RGBA{100, 255, 100, 255})

	display.DrawCircle(ggweb.Fill, x+20, y+20, 20)
	display.DrawEllipse(ggweb.Fill, geo.Rect{X: x + 50, Y: y, W: 15, H: 40})
	display.DrawArc(ggweb.Fill, geo.Rect{X: x + 70, Y: y, W: 40, H: 20}, math.Pi/4, 5*math.Pi/4, true)
	display.DrawArc(ggweb.Stroke, geo.Rect{X: x + 70, Y: y + 20, W: 40, H: 20}, math.Pi/4, 5*math.Pi/4, true)

	display.Restore()
}

func Path(display *ggweb.Surface, x, y float64) {
	display.Save()

	path := ggweb.NewPath()
	path.MoveTo(x, y)
	path.LineTo(x+40, y)
	path.LineTo(x+40, y+40)
	path.Close()
	path.Rect(geo.Rect{X: x + 50, Y: y, W: 20, H: 20})
	path.MoveTo(x+60, y+30)
	path.Arc(x+60, y+30, 10, 0, 3*math.Pi/2, true)
	path.MoveTo(x+80, y)
	path.ArcTo(x+80, y+40, x+100, y+40, 15)
	r := geo.Rect{X: x + 100, Y: y, W: 20, H: 40}
	path.MoveTo(r.Center())
	path.Ellipse(r, math.Pi/4, math.Pi/4, math.Pi, false)
	path.MoveTo(x+130, y)
	path.QuadraticCurveTo(x+130, y+30, x+150, y+40)
	path.BezierCurveTo(x+160, y+20, x+140, y+20, x+150, y)

	display.StyleColor(ggweb.Fill, color.RGBA{255, 255, 100, 255})
	display.StyleColor(ggweb.Stroke, color.RGBA{255, 100, 100, 255})
	display.SetLineWidth(5)
	display.SetLineJoin(ggweb.LineJoinRound)
	display.SetLineCap(ggweb.LineCapRound)

	display.DrawPath(ggweb.Fill, path)
	display.DrawPath(ggweb.Stroke, path)

	clipX, clipY := x, y+80
	clip := ggweb.NewPath()
	clip.Arc(clipX, clipY, 30, 0, 2*math.Pi, true)
	display.ClipPath(clip)

	patSurf := ggweb.NewSurface(15, 15)
	patSurf.StyleLinearGradient(ggweb.Fill, ggweb.LinearGradient{
		X1: 0, Y1: 0, X2: 15, Y2: 15,
		ColorStops: []ggweb.ColorStop{
			{Position: 0, Color: color.RGBA{255, 0, 0, 255}},
			{Position: 1, Color: color.RGBA{0, 0, 255, 255}},
		},
	})
	patSurf.DrawRect(ggweb.Fill, patSurf.Rect())
	display.StylePattern(ggweb.Fill, ggweb.Pattern{
		Source: patSurf,
		Type:   ggweb.RepeatXY,
	})
	display.DrawRect(ggweb.Fill, geo.Rect{X: clipX - 40, Y: clipY - 40, W: 80, H: 80})

	display.Restore()
}

func Transform(display *ggweb.Surface, x, y float64) {
	display.Save()

	display.Translate(x, y)
	display.StyleColor(ggweb.Stroke, color.RGBA{100, 100, 255, 255})
	display.SetLineWidth(3)
	r := geo.Rect{X: 0, Y: 0, W: 30, H: 30}
	display.DrawRect(ggweb.Stroke, r)

	display.Rotate(-math.Pi / 4)
	display.StyleColor(ggweb.Stroke, color.RGBA{100, 255, 100, 128})
	display.DrawRect(ggweb.Stroke, r)

	display.Scale(1.5, 1.5)
	display.StyleColor(ggweb.Stroke, color.RGBA{255, 100, 100, 128})
	display.DrawRect(ggweb.Stroke, r)

	display.Restore()

	display.Save()
	display.StyleColor(ggweb.Stroke, color.RGBA{255, 100, 255, 255})
	display.SetLineWidth(3)
	display.Transform(1, 0, 1, 1, x+40, y)
	display.DrawRect(ggweb.Stroke, geo.Rect{X: 0, Y: 0, W: 30, H: 30})
	display.Restore()
}

func Text(display *ggweb.Surface, x, y float64) {
	display.Save()

	display.StyleColor(ggweb.Fill, color.White)
	display.StyleColor(ggweb.Stroke, color.White)

	display.DrawText(ggweb.Fill, "default text", x, y+10)

	f := ggweb.Font{
		Size:   20,
		Family: ggweb.FontFamilyMonospace,
		Weight: ggweb.FontWeightBold,
	}
	display.SetFont(&f)
	display.DrawText(ggweb.Stroke, "different font", x, y+30)

	display.SetTextAlign(ggweb.TextAlignCenter)
	display.SetTextBaseline(ggweb.TextBaselineMiddle)
	display.DrawText(ggweb.Fill, "centered", x+display.TextWidth("different font")/2, y+40)

	display.Restore()
}

func Image(display *ggweb.Surface, x, y float64) {
	display.Save()

	img := ggweb.LoadImage("gopher.png")
	display.Blit(img, x, y)
	display.BlitArea(img, geo.Rect{X: 50, Y: 10, W: 100, H: 100}, x+img.Rect().W+5, y)

	display.Restore()
}

func Pixel(display *ggweb.Surface, x, y float64) {
	display.Save()

	area := geo.Rect{X: 0, Y: 0, W: display.Rect().W - 20, H: 260}
	pixels := display.PixelData(area)
	for i := 0; i < len(pixels); i++ {
		pixels[i].R = 255 - pixels[i].R
		pixels[i].G = 255 - pixels[i].G
		pixels[i].B = 255 - pixels[i].B
	}
	area.X = x
	area.Y = y
	display.SetPixelData(pixels, area)

	display.Restore()
}

func Surface(display *ggweb.Surface, x, y float64) {
	display.Save()

	s := ggweb.NewSurface(int(display.Rect().W), int(display.Rect().H))
	s.Blit(display, 0, 0)

	display.Translate(x, y)
	display.Scale(0.4, 0.4)
	display.Blit(s, 0, 0)
	display.StyleColor(ggweb.Stroke, color.White)
	display.SetLineWidth(3)
	display.DrawRect(ggweb.Stroke, s.Rect())

	display.Restore()
}

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
// }

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
