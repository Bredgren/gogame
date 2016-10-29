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
