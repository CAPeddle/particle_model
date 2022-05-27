package environment

import (
	"CAPeddle/particle_model/particle"
	"fmt"
)

//particles have a location
type Environment struct {
	width, height int
	particles     []particle.Particle
}

// NewLife returns a new Life game state with a random initial state.
func NewEnvironment(w, h int) *Environment {

	testParticle := particle.Particle{}
	fmt.Print("Width [", w, "] Height [", h, "] Test Particle", testParticle.String())

	newParticles := []particle.Particle{}

	// for i, j := 0, 0; i < 10; i, j = i+1, j+1 {
	// 	fmt.Println("Hello, playground", i, " ", j)
	// }

	for x, y := 0, 0; x < w && y < h; x, y = x+1, y+1 {
		newParticles = append(newParticles, particle.Particle{
			Alive: true,
			Location: particle.Position{
				X: x,
				Y: y,
			},
		})
	}

	newParticles = append(newParticles, particle.Particle{
		Alive: true,
		Location: particle.Position{
			X: w - 1,
			Y: h - 1,
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

	for i, v := range environment.particles {
		pixelLocation := v.Location.Y*environment.width + v.Location.X

		// fmt.Println(i, ": Range [", len(pix),
		// 	"] pixelLocation [", pixelLocation,
		// 	"] maxColLocation [", 4*pixelLocation+3,
		// 	"] LocY[", v.Location.Y*environment.width,
		// 	"] LocX [", v.Location.X, "]")

		if 4*pixelLocation > len(pix) {
			fmt.Println(i, ": Exceeds range")
			break
		}

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
