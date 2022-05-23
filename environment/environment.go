package environment

import (
	"CAPeddle/particle_model/particle"
)

//particles have a location
type Environment struct {
	width, height int
	particles     []particle.Particle
}

// NewLife returns a new Life game state with a random initial state.
func NewEnvironment(w, h int) *Environment {

	testParticle := particle.Particle{}
	print(testParticle.String())

	newParticles := []particle.Particle{}

	newParticles = append(newParticles, particle.Particle{
		Alive: true,
		Location: particle.Position{
			X: 10,
			Y: 10,
		},
	})

	newParticles = append(newParticles, particle.Particle{
		Alive: false,
		Location: particle.Position{
			X: 20,
			Y: 20,
		},
	})

	newParticles = append(newParticles, particle.Particle{
		Alive: true,
		Location: particle.Position{
			X: 30,
			Y: 30,
		},
	})

	newParticles = append(newParticles, particle.Particle{
		Alive: false,
		Location: particle.Position{
			X: 40,
			Y: 40,
		},
	})

	return &Environment{
		width:     w,
		height:    h,
		particles: newParticles,
	}
}

// Draw paints current game state.
func (environment *Environment) Draw(pix []byte) {

	i := 0
	for y := 0; y < environment.height; y++ {
		for x, v := range environment.particles {
			if v.Alive && x == v.Location.X {
				pix[4*i] = 0xff
				pix[4*i+1] = 0xff
				pix[4*i+2] = 0xff
				pix[4*i+3] = 0xff
			} else {
				pix[4*i] = 0
				pix[4*i+1] = 0
				pix[4*i+2] = 0
				pix[4*i+3] = 0
			}
			i++
		}
	}
}
