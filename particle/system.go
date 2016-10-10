package particle

import (
	"math"
	"time"

	"github.com/Bredgren/gogame/geo"
)

// SystemParticle a particle with extra information for use by a System.
type SystemParticle struct {
	Particle
	Life       time.Duration
	Active     bool
	inFreeList bool
}

// System is a manager for large groups of particles.
type System struct {
	// Rate is the number of new particles per second.
	Rate float64
	// InitPos/Vel/Mass/Life are the starting parameters for each new particle.
	InitPos      geo.VecGen
	InitVel      geo.VecGen
	InitMass     geo.NumGen
	InitLife     time.Duration
	pool         []SystemParticle
	freeList     chan *SystemParticle
	globalForce  geo.Vec
	lastParticle time.Duration
}

// NewSystem initializes a particle system configured to contain at most size particles.
func NewSystem(size int) *System {
	s := System{}
	s.pool = make([]SystemParticle, size)
	s.freeList = make(chan *SystemParticle, size)
	for i := range s.pool {
		s.freeList <- &s.pool[i]
		s.pool[i].inFreeList = true
	}
	return &s
}

// Particles returns a slice of all active particles.
func (s *System) Particles() []*SystemParticle {
	active := make([]*SystemParticle, 0, len(s.pool))
	for i := range s.pool {
		if s.pool[i].Active {
			active = append(active, &s.pool[i])
		}
	}
	return active
}

// ForEachParticle calls f for each active particle.
func (s *System) ForEachParticle(f func(p *SystemParticle)) {
	for i := range s.pool {
		if s.pool[i].Active {
			f(&s.pool[i])
		}
	}
}

// Update updates the state of all active particles and creates new particles if the limit
// hasn't been reached yet. dt is the amount of time to simulate.
func (s *System) Update(dt time.Duration) {
	for i := range s.pool {
		if s.pool[i].Active {
			s.pool[i].ApplyForce(s.globalForce)
			s.pool[i].Update(dt)
			s.pool[i].Life -= dt
			if s.pool[i].Life <= 0 {
				s.pool[i].Active = false
			}
		}
	}
	s.globalForce.Mul(0)

	for i := range s.pool {
		if !s.pool[i].Active && !s.pool[i].inFreeList {
			s.freeList <- &s.pool[i]
			s.pool[i].inFreeList = true
		}
	}

	s.lastParticle += dt
	newCount := int(math.Floor(s.Rate * s.lastParticle.Seconds()))
	for newCount > 0 && len(s.freeList) > 0 {
		newParticle := <-s.freeList
		newParticle.inFreeList = false
		newParticle.Active = true
		newParticle.Life = s.InitLife
		newParticle.Pos = s.InitPos()
		newParticle.Vel = s.InitVel()
		newParticle.Mass = s.InitMass()
		s.lastParticle = 0
		newCount--
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
func (s *System) ApplyForce(force geo.Vec) {
	s.globalForce.Add(force)
}
