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
			X: 0,
			Y: 0,
		},
	})

	newParticles = append(newParticles, particle.Particle{
		Alive: true,
		Location: particle.Position{
			X: 20,
			Y: 10,
		},
	})

	newParticles = append(newParticles, particle.Particle{
		Alive: true,
		Location: particle.Position{
			X: 30,
			Y: 10,
		},
	})

	newParticles = append(newParticles, particle.Particle{
		Alive: true,
		Location: particle.Position{
			X: w,
			Y: h,
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

	for _, v := range environment.particles {
		pixelLocation := v.Location.Y*environment.width + v.Location.X
		if v.Alive {
			pix[4*pixelLocation] = 0xff
			pix[4*pixelLocation+1] = 0xff
			pix[4*pixelLocation+2] = 0xff
			pix[4*pixelLocation+3] = 0xff
		} else {
			pix[4*pixelLocation] = 0
			pix[4*pixelLocation+1] = 0
			pix[4*pixelLocation+2] = 0
			pix[4*pixelLocation+3] = 0
		}
	}
}
