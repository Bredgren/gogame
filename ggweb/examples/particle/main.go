package main

import (
	"image/color"
	"math"
	"time"

	"github.com/Bredgren/gogame/event"
	"github.com/Bredgren/gogame/geo"
	"github.com/Bredgren/gogame/ggweb"
	"github.com/Bredgren/gogame/key"
	"github.com/Bredgren/gogame/particle"
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

	ggweb.RegisterEvents(display)
	setup()

	ggweb.DisableContextMenu = true

	ggweb.SetMainLoop(mainLoop)
}

var ps1 *particle.System
var ps2 *particle.System

var explosionPos geo.Vec

func setup() {
	ps1 = particle.NewSystem(400)
	ps1.InitLife = time.Duration(3 * time.Second)
	ps1.Rate = 100
	ps1.InitMass = geo.RandNum(0.5, 3)
	ps1.InitPos = geo.StaticVec(geo.Vec{X: 150, Y: 200})
	ps1.InitVel = geo.RandVecArc(0, 200, math.Pi/4, 3*math.Pi/4)

	ps2 = particle.NewSystem(100)
	ps2.InitLife = time.Duration(500 * time.Millisecond)
	ps2.Rate = 0
	ps2.InitMass = geo.RandNum(0.9, 1.1)
	ps2.InitPos = geo.DynamicVec(&explosionPos)
	ps2.InitVel = geo.RandVecCircle(50, 500)
}

var lastT time.Duration

var font1 = ggweb.Font{
	Size: 20,
}

var font2 = ggweb.Font{
	Size: 15,
}

func mainLoop(t time.Duration) {
	dt := t - lastT
	lastT = t

	for evt := event.Poll(); evt.Type != event.NoEvent; evt = event.Poll() {
		switch evt.Type {
		case event.WindowResize:
			data := evt.Data.(event.ResizeData)
			display.SetSize(data.W, data.H)
		case event.MouseButtonDown:
			data := evt.Data.(event.MouseData)
			if data.Button == 0 {
				explosionPos.X = data.Pos.X
				explosionPos.Y = data.Pos.Y
				// Make all 100 particles at once
				ps2.Rate = 100 / dt.Seconds()
			}
		}
	}

	if ggweb.PressedKeys()[key.W] {
		wind := geo.Vec{X: 1000}
		ps1.ApplyForce(wind)
		ps2.ApplyForce(wind)
	}

	gravity := geo.Vec{Y: 500}

	ps1.ApplyForce(gravity)
	ps1.Update(dt)
	ps2.ForEachParticle(func(p *particle.SystemParticle) {
		p.ApplyForce(p.Vel.Normalized().Times(-p.Vel.Len2() * 0.01))
	})
	ps2.Update(dt)
	ps2.Rate = 0 // Stop making particles

	display.StyleColor(ggweb.Fill, color.Black)
	display.DrawRect(ggweb.Fill, display.Rect())
	display.StyleColor(ggweb.Fill, color.White)
	ps1.ForEachParticle(func(p *particle.SystemParticle) {
		display.DrawCircle(ggweb.Fill, p.Pos.X, p.Pos.Y, 2.5*p.Mass)
		// a := geo.Rect{X: p.Pos.X, Y: p.Pos.Y, W: 1, H: 1}
		// pdata := display.PixelData(a)
		// pdata[0].R = 255
		// pdata[0].G = 255
		// pdata[0].B = 255
		// display.SetPixelData(pdata, a)
	})
	ps2.ForEachParticle(func(p *particle.SystemParticle) {
		display.StyleColor(ggweb.Fill, color.RGBA{255, 30, 30, uint8(255 * p.Life / ps2.InitLife)})
		display.DrawCircle(ggweb.Fill, p.Pos.X, p.Pos.Y, 5*p.Mass)
	})

	display.StyleColor(ggweb.Fill, color.White)
	display.SetTextAlign(ggweb.TextAlignLeft)
	display.SetTextBaseline(ggweb.TextBaselineTop)
	display.SetFont(&font1)
	display.DrawText(ggweb.Fill, ggweb.Stats.LoopDuration.String(), 2, 0)
	display.SetFont(&font2)
	display.DrawText(ggweb.Fill, "click to explode", 2, 30)
	display.DrawText(ggweb.Fill, "w for wind", 2, 45)
}
