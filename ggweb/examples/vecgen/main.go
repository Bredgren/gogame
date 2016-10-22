package main

import (
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/Bredgren/gogame/geo"
	"github.com/Bredgren/gogame/ggweb"
)

func main() {
	ggweb.Init(onReady)
}

const n = 2000

func onReady() {
	rand.Seed(time.Now().Unix())

	width, height := 900, 600
	display := ggweb.NewSurfaceFromID("main")
	display.SetSize(width, height)
	display.StyleColor(ggweb.Fill, color.Black)
	display.DrawRect(ggweb.Fill, display.Rect())

	gens := []geo.VecGen{}

	display.StyleColor(ggweb.Stroke, color.RGBA{100, 100, 255, 255})
	display.SetLineWidth(2)

	gap := 20.0

	circleR := 100.0
	circlePos := geo.Vec{X: circleR + gap, Y: circleR + gap}
	display.DrawCircle(ggweb.Stroke, circlePos.X, circlePos.Y, circleR)
	gens = append(gens, geo.OffsetVec(geo.RandVecCircle(0, circleR), geo.StaticVec(circlePos)))

	circleHoleR := 50.0
	circleHolePos := geo.Vec{X: circlePos.X + 2*circleR + gap, Y: circlePos.Y}
	display.DrawCircle(ggweb.Stroke, circleHolePos.X, circleHolePos.Y, circleHoleR)
	display.DrawCircle(ggweb.Stroke, circleHolePos.X, circleHolePos.Y, circleR)
	gens = append(gens, geo.OffsetVec(geo.RandVecCircle(circleHoleR, circleR), geo.StaticVec(circleHolePos)))

	arcMin, arcMax := 3*math.Pi/4, 3*math.Pi/2
	arcRect := geo.Rect{X: circleHolePos.X + circleR + gap, Y: gap, W: 2 * circleR, H: 2 * circleR}
	display.DrawArc(ggweb.Stroke, arcRect, arcMin, arcMax, true)
	arcHoleRect := geo.Rect{W: 2 * circleHoleR, H: 2 * circleHoleR}
	arcHoleRect.SetCenter(arcRect.Center())
	display.DrawArc(ggweb.Stroke, arcHoleRect, arcMin, arcMax, true)
	display.DrawLine(
		math.Cos(arcMin)*circleHoleR+arcRect.CenterX(),
		-math.Sin(arcMin)*circleHoleR+arcRect.CenterY(),
		math.Cos(arcMin)*circleR+arcRect.CenterX(),
		-math.Sin(arcMin)*circleR+arcRect.CenterY())
	display.DrawLine(
		math.Cos(arcMax)*circleHoleR+arcRect.CenterX(),
		-math.Sin(arcMax)*circleHoleR+arcRect.CenterY(),
		math.Cos(arcMax)*circleR+arcRect.CenterX(),
		-math.Sin(arcMax)*circleR+arcRect.CenterY())
	gens = append(gens, geo.OffsetVec(geo.RandVecArc(circleHoleR, circleR, arcMin, arcMax),
		geo.StaticVec(geo.Vec{X: arcRect.CenterX(), Y: arcRect.CenterY()})))

	r := geo.Rect{X: gap, Y: circlePos.Y + circleR + gap, W: 300, H: 200}
	display.DrawRect(ggweb.Stroke, r)
	gens = append(gens, geo.RandVecRect(r))

	x, y := gap+r.W+gap, r.Y
	w, h, t := 300.0, 150.0, 40.0
	rs := []geo.Rect{
		{X: x, Y: y, W: w - t, H: t},
		{X: x + w - t, Y: y, W: t, H: h - t},
		{X: x + t, Y: y + h - t, W: w - t, H: t},
		{X: x, Y: y + t, W: t, H: h - t},
	}
	for _, r := range rs {
		display.DrawRect(ggweb.Stroke, r)
	}
	gens = append(gens, geo.RandVecRects(rs))

	pointR := 1.0
	display.StyleColor(ggweb.Fill, color.White)
	for i := 0; i < n; i++ {
		for _, gen := range gens {
			p := gen()
			display.DrawCircle(ggweb.Fill, p.X, p.Y, pointR)
		}
	}
}
