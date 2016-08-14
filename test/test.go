// To run the tests run "gopherjs serve $GOPATH/github.com/Bredgren/gogame/test" and navigate
// to http://localhost:8080/github.com/Bredgren/gogame/test/ in your browser.
package main

import (
	"math"

	"github.com/Bredgren/gogame"
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
	testRect()

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
	display := gogame.GetDisplay()
	display.SetMode(width, height)
	display.Fill(&gogame.FillStyle{Colorer: gogame.Black})

	display.DrawRect(&gogame.Rect{X: 11, Y: 11, W: 48, H: 48}, &gogame.StrokeStyle{Colorer: gogame.White, Width: 4})
	display.DrawRect(&gogame.Rect{X: 70, Y: 10, W: 50, H: 50},
		&gogame.FillStyle{Colorer: &gogame.LinearGradient{
			X1: 0, Y1: 0, X2: 50, Y2: 50,
			ColorStops: []gogame.ColorStop{
				{0.0, gogame.Red},
				{1.0, gogame.Blue},
			},
		}})
	display.DrawRect(&gogame.Rect{X: 130, Y: 10, W: 50, H: 50},
		&gogame.FillStyle{Colorer: &gogame.RadialGradient{
			X1: 50 / 2, Y1: 50 / 2, R1: 40, X2: 50 / 4, Y2: 50 / 4, R2: 1,
			ColorStops: []gogame.ColorStop{
				{0.0, gogame.Blue},
				{1.0, gogame.Green},
			},
		}})

	s := gogame.NewSurface(10, 10)
	s.DrawRect(&gogame.Rect{X: 0, Y: 0, W: 10, H: 10},
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

	display.DrawRect(&gogame.Rect{X: 190, Y: 10, W: 50, H: 50},
		&gogame.FillStyle{Colorer: &pattern})

	display.Blit(s, 30, 30)

	display.DrawCircle(35, 95, 20, &gogame.StrokeStyle{
		Colorer: &pattern,
		Width:   10,
	})

	display.DrawEllipse(&gogame.Rect{X: 70, Y: 70, W: 110, H: 50}, &gogame.StrokeStyle{
		Colorer: gogame.White,
		Width:   2,
	})

	display.DrawArc(&gogame.Rect{X: 190, Y: 70, W: 110, H: 50}, 0, 0.75*math.Pi, &gogame.StrokeStyle{
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

	display.DrawLines([][2]float64{{390, 130}, {440, 130}, {390, 180}}, &gogame.FillStyle{
		Colorer: gogame.White,
	})
}
