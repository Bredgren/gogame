package main

import (
	"math"
	"time"

	"github.com/Bredgren/gogame"
	"github.com/Bredgren/gogame/event"
	"github.com/Bredgren/gogame/geo"
	"github.com/Bredgren/gogame/key"
	"github.com/Bredgren/gogame/particle"
)

func main() {
	gogame.Ready(onReady)
}

func onReady() {
	display := gogame.MainDisplay()
	display.SetMode(900, 600)
	display.Fill(gogame.FillBlack)
	setup()
	gogame.SetMainLoop(loop)
}

var ps1 *particle.System
var ps2 *particle.System

var explosionPos geo.Vec

func setup() {
	ps1 = particle.NewSystem(200)
	ps1.InitLife = time.Duration(3 * time.Second)
	ps1.Rate = 50
	ps1.InitMass = 1
	ps1.InitPos = geo.StaticVec(geo.Vec{X: 150, Y: 200})
	ps1.InitVel = geo.RandVecArc(0, -200, math.Pi/4, 3*math.Pi/4)

	ps2 = particle.NewSystem(100)
	ps2.InitLife = time.Duration(500 * time.Millisecond)
	ps2.Rate = 0
	ps2.InitMass = 1
	ps2.InitPos = geo.DynamicVec(&explosionPos)
	ps2.InitVel = geo.RandVecCircle(50, 500)
}

var lastT time.Duration

func loop(t time.Duration) {
	dt := t - lastT
	lastT = t

	for evt := event.Poll(); evt.Type != event.NoEvent; evt = event.Poll() {
		switch evt.Type {
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

	if gogame.PressedKeys()[key.W] {
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

	display := gogame.MainDisplay()
	display.Fill(gogame.FillBlack)
	ps1.ForEachParticle(func(p *particle.SystemParticle) {

		display.DrawCircle(p.Pos.X, p.Pos.Y, 5, &gogame.FillStyle{
			Colorer: gogame.ColorCSS("#FFF"),
		})
	})
	ps2.ForEachParticle(func(p *particle.SystemParticle) {
		// Note that using Color is a major performance bottleneck because it converts numbers
		// to a strings, using ColorCSS, like above, can save a lot of time.
		display.DrawCircle(p.Pos.X, p.Pos.Y, 5, &gogame.FillStyle{
			Colorer: gogame.Color{R: 1, G: 0.2, B: 0.2, A: p.Life.Seconds()},
		})
	})
	textStyle := gogame.TextStyle{
		Colorer:  gogame.White,
		Align:    gogame.TextAlignLeft,
		Baseline: gogame.TextBaselineTop,
	}
	display.DrawText(gogame.Stats.LoopDuration.String(), 2, 0, &gogame.Font{Size: 20}, &textStyle)
	display.DrawText("click to explode", 2, 30, &gogame.Font{Size: 15}, &textStyle)
	display.DrawText("w for wind", 2, 45, &gogame.Font{Size: 15}, &textStyle)
	display.Flip()
}
