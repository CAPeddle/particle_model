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

	newParticles := []particle.Particle{}

	for x, y := 0, 0; x < w && y < h; x, y = x+1, y+1 {
		newParticles = append(newParticles, particle.Particle{
			Alive: true,
			Location: particle.Position{
				X: x,
				Y: y,
			},
		})
	}

	return &Environment{
		width:     w,
		height:    h,
		particles: newParticles,
	}
}

// Draw paints current game state.
func (environment *Environment) Draw(pix []byte) {

	for i, v := range environment.particles {
		// fmt.Println(i, ": Range [", len(pix),
		// 	"] pixelLocation [", pixelLocation,
		// 	"] maxColLocation [", 4*pixelLocation+3,
		// 	"] LocY[", v.Location.Y*environment.width,
		// 	"] LocX [", v.Location.X, "]")

		//object/particle should respond with a drawable impact on the pixel array
		//should it be an addition on the array?
		//so that this method is dumb?

		fmt.Println("len Col ", len(v.Render().ImageOf),
			"len Rows ", len(v.Render().ImageOf[0]),
			" A ", v.Render().ImageOf[0][0].A,
			" Modulos ImageOf", len(v.Render().ImageOf)%2,
			" Modulos ImageOf[0]", len(v.Render().ImageOf[0])%2)

		if len(v.Render().ImageOf[0])%2 == 0 {
			fmt.Println("Width not uneven")
			break
		}

		if len(v.Render().ImageOf)%2 == 0 {
			fmt.Println("Height not uneven")
			break
		}

		pixelLocation := v.Location.Y*environment.width + v.Location.X

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
