package particle

import (
	"time"

	"github.com/Bredgren/gogame/geo"
)

// Particle is a generatic particle that has a position and velocity. It also allows
// forces to be applied to it.
type Particle struct {
	// Pos is the position of the particle. No particular units are assumed.
	Pos geo.Vec
	// Vel is the velocity of the particle. It is the number units per second Pos changes
	// on a call to Update.
	Vel geo.Vec
	// A Mass of 0 will ignore applied forces.
	Mass  float64
	accel geo.Vec
}

// Update causes all applied forces since the last call to Update to take effect on the
// particle and updates the position based on the, possibly new, velocity. dt is the
// amount of time to advance.
func (p *Particle) Update(dt time.Duration) {
	p.Vel.Add(p.accel.Times(dt.Seconds()))
	p.Pos.Add(p.Vel.Times(dt.Seconds()))
	p.accel.Mul(0)
}

// ApplyForce applies the given force to the particle taking its Mass into account. It will
// not affect the particle until Update is called. If multiple forces are applied then they
// will accumulate. Forces will not persist between calls to Update.
func (p *Particle) ApplyForce(force geo.Vec) {
	if p.Mass > 0 {
		p.accel.Add(force.DividedBy(p.Mass))
	}
}
