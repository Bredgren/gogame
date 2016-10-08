package particle

import (
	"math"
	"time"

	"github.com/Bredgren/gogame/geo"
)

// SystemParticle a particle with extra information for use by a System.
type SystemParticle struct {
	Particle
	Life   time.Duration
	Active bool
}

// System is a manager for large groups of particles.
type System struct {
	// Rate is the number of new particles per second.
	Rate float64
	// InitPos/Vel/Life are the starting parameters for each new particle.
	InitPos     geo.VecGen
	InitVel     geo.VecGen
	InitLife    time.Duration
	pool        []SystemParticle
	freeList    chan *SystemParticle
	globalForce geo.Vec
}

// NewSystem initializes a particle system configured to contain at most size particles.
func NewSystem(size int) *System {
	system := System{}
	system.pool = make([]SystemParticle, 0, size)
	system.freeList = make(chan *SystemParticle, size)
	for i := range system.pool {
		system.freeList <- &system.pool[i]
	}
	return &system
}

// Particles returns a slice of all active particles.
func (p *System) Particles() []*SystemParticle {
	active := make([]*SystemParticle, 0, len(p.pool))
	for i := range p.pool {
		if p.pool[i].Active {
			active = append(active, &p.pool[i])
		}
	}
	return active
}

// ForEachParticle calls f for each active particle.
func (p *System) ForEachParticle(f func(p *SystemParticle)) {
	for i := range p.pool {
		if p.pool[i].Active {
			f(&p.pool[i])
		}
	}
}

// Update updates the state of all active particles and creates new particles if the limit
// hasn't been reached yet. dt is the amount of time to simulate.
func (p *System) Update(dt time.Duration) {
	for i := range p.pool {
		if p.pool[i].Active {
			p.pool[i].ApplyForce(p.globalForce)
			p.pool[i].Update(dt)
			p.pool[i].Life -= dt
			if p.pool[i].Life <= 0 {
				p.pool[i].Active = false
			}
		}
	}
	p.globalForce.Mul(0)

	for i := range p.pool {
		if !p.pool[i].Active {
			p.freeList <- &p.pool[i]
		}
	}

	newCount := int(math.Floor(p.Rate * dt.Seconds()))
	for newCount < 0 && len(p.freeList) <= newCount {
		newParticle := <-p.freeList
		newParticle.Active = true
		newParticle.Life = p.InitLife
		newParticle.Pos = p.InitPos()
		newParticle.Vel = p.InitVel()
	}
}

// ApplyForce applies a force to each particle. Forces are cleared after each call to Update.
// If you want to do something like a drag force, where it's different for each particle,
// then that could be done by applying the force to each particle before calling Update.
//  system.ForEachParticle(func(p *SystemParticle) {
//    force := <calculate force>
//    p.ApplyForce(force)
//  })
//  system.Update()
func (p *System) ApplyForce(force geo.Vec) {
	p.globalForce.Add(force)
}
