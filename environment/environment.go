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
	})

	newParticles = append(newParticles, particle.Particle{
		Alive: false,
	})

	newParticles = append(newParticles, particle.Particle{
		Alive: true,
	})

	newParticles = append(newParticles, particle.Particle{
		Alive: false,
	})

	return &Environment{
		width:     w,
		height:    h,
		particles: newParticles,
	}
}

// Draw paints current game state.
func (environment *Environment) Draw(pix []byte) {

	print(" Draw:")
	for i, v := range environment.particles {
		print(i, v.Alive)
	}

	// i := 0
	// for y := 0; y < environment.height; y++ {
	// 	for _, v := range environment.particles[y] {
	// 		if v {
	// 			pix[4*i] = 0xff
	// 			pix[4*i+1] = 0xff
	// 			pix[4*i+2] = 0xff
	// 			pix[4*i+3] = 0xff
	// 		} else {
	// 			pix[4*i] = 0
	// 			pix[4*i+1] = 0
	// 			pix[4*i+2] = 0
	// 			pix[4*i+3] = 0
	// 		}
	// 		i++
	// 	}
	// }
}
