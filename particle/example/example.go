package main

import (
	"math"
	"time"

	"github.com/Bredgren/gogame"
	"github.com/Bredgren/gogame/event"
	"github.com/Bredgren/gogame/geo"
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
	ps1.InitPos = geo.StaticVec(geo.Vec{X: 150, Y: 300})
	ps1.InitVel = geo.RandVecArc(0, -150, math.Pi/4, 3*math.Pi/4)

	ps2 = particle.NewSystem(100)
	ps2.InitLife = time.Duration(500 * time.Millisecond)
	ps2.Rate = 0
	ps2.InitMass = 1
	ps2.InitPos = geo.DynamicVec(&explosionPos)
	ps2.InitVel = geo.RandVecCircle(50, 300)
}

var lastT time.Duration

func loop(t time.Duration) {
	dt := t - lastT
	lastT = t

	for evt := event.Poll(); evt.Type != event.NoEvent; evt = event.Poll() {
		switch evt.Type {
		case event.KeyDown:
		case event.KeyUp:
		case event.MouseButtonDown:
			data := evt.Data.(event.MouseData)
			if data.Button == 0 {
				explosionPos.X = data.Pos.X
				explosionPos.Y = data.Pos.Y
				ps2.Rate = 100 / dt.Seconds()
			}
		case event.MouseButtonUp:
		}
	}

	gravity := geo.Vec{Y: 100}

	ps1.ApplyForce(gravity)
	ps1.Update(dt)
	ps2.ForEachParticle(func(p *particle.SystemParticle) {
		p.ApplyForce(p.Vel.Times(-0.2))
	})
	ps2.Update(dt)
	ps2.Rate = 0

	display := gogame.MainDisplay()
	display.Fill(gogame.FillBlack)
	ps1.ForEachParticle(func(p *particle.SystemParticle) {
		display.DrawCircle(p.Pos.X, p.Pos.Y, 5, &gogame.FillStyle{
			Colorer: gogame.Color{R: 1, G: 1, B: 1, A: p.Life.Seconds() / 3},
		})
		// display.SetAt(int(p.Pos.X), int(p.Pos.Y), gogame.White)
	})
	ps2.ForEachParticle(func(p *particle.SystemParticle) {
		display.DrawCircle(p.Pos.X, p.Pos.Y, 5, &gogame.FillStyle{
			Colorer: gogame.Color{R: 1, G: 0.2, B: 0.2, A: p.Life.Seconds()},
		})
		// display.SetAt(int(p.Pos.X), int(p.Pos.Y), gogame.White)
	})
	display.DrawText(gogame.Stats.LoopDuration.String(), 2, 0,
		&gogame.Font{
			Size: 20,
		},
		&gogame.TextStyle{
			Colorer:  gogame.White,
			Align:    gogame.TextAlignLeft,
			Baseline: gogame.TextBaselineTop,
		})
	display.Flip()
}
