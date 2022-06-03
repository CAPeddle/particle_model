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

	for x, y := particle.ParticleWidth, particle.ParticleHeight; x < w-particle.ParticleWidth && y < h-particle.ParticleHeight; x, y = x+particle.ParticleWidth+1, y+particle.ParticleHeight+1 {
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

	for _, v := range environment.particles {

		if len(v.Render().ImageOf[0])%2 == 0 {
			fmt.Println("Width not uneven")
			break
		}

		if len(v.Render().ImageOf)%2 == 0 {
			fmt.Println("Height not uneven")
			break
		}

		noRows := len(v.Render().ImageOf)
		noCols := len(v.Render().ImageOf[0])

		for rowIndex, row := range v.Render().ImageOf {
			//TODO improve the implicit conversion
			rowPixel := rowIndex - noRows/2

			for colIndex, pixel := range row {
				colPixel := colIndex - noCols/2

				pixelLocation := (v.Location.Y-rowPixel)*environment.width + v.Location.X + colPixel

				if 4*pixelLocation+3 > len(pix) || 4*pixelLocation < 0 {
					fmt.Println("Pixel out of bounds")
					break
				}

				if v.Alive {
					pix[4*pixelLocation] = pixel.R
					pix[4*pixelLocation+1] = pixel.G
					pix[4*pixelLocation+2] = pixel.B
					pix[4*pixelLocation+3] = pixel.A
				} else {
					pix[4*pixelLocation] = 0
					pix[4*pixelLocation+1] = 0
					pix[4*pixelLocation+2] = 0
					pix[4*pixelLocation+3] = 0
				}
			}
		}
	}
}

func (environment *Environment) Update() {
	for particleIndex := range environment.particles {
		environment.particles[particleIndex].Location.X = environment.particles[particleIndex].Location.X + 1

		if environment.particles[particleIndex].Location.X > environment.width-particle.ParticleWidth {
			environment.particles[particleIndex].Location.X = particle.ParticleWidth
		}

		environment.particles[particleIndex].Location.Y = environment.particles[particleIndex].Location.Y + 1

		if environment.particles[particleIndex].Location.Y > environment.height-particle.ParticleHeight {
			environment.particles[particleIndex].Location.Y = particle.ParticleHeight
		}
	}
}
