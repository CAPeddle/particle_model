package environment

import (
	"CAPeddle/particle_model/particle"
	"fmt"
	"math/rand"
	"time"
)

//particles have a location
type Environment struct {
	width, height int
	particles     []particle.Particle
}

// NewLife returns a new Life game state with a random initial state.
func NewEnvironment(w, h int) *Environment {

	newParticles := []particle.Particle{}

	src := rand.NewSource(time.Now().UnixNano())

	remaining := 0
	cache := int64(0)

	for x := 0; x < w-particle.ParticleWidth; x = x + particle.ParticleWidth {
		for y := 0; y < h-particle.ParticleHeight; y = y + particle.ParticleHeight {

			if remaining == 0 {
				cache, remaining = src.Int63(), 63
			}

			result := cache&0x01 == 1
			cache >>= 1
			remaining--

			newParticles = append(newParticles, particle.Particle{
				Alive: result,
				LocationRender: particle.Position{
					X: x,
					Y: y,
				},
			})
		}
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

				pixelLocation := (v.LocationRender.Y-rowPixel)*environment.width + v.LocationRender.X + colPixel

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
		environment.particles[particleIndex].LocationRender.X = environment.particles[particleIndex].LocationRender.X + 1

		if environment.particles[particleIndex].LocationRender.X > environment.width-particle.ParticleWidth {
			environment.particles[particleIndex].LocationRender.X = particle.ParticleWidth
		}

		environment.particles[particleIndex].LocationRender.Y = environment.particles[particleIndex].LocationRender.Y + 1

		if environment.particles[particleIndex].LocationRender.Y > environment.height-particle.ParticleHeight {
			environment.particles[particleIndex].LocationRender.Y = particle.ParticleHeight
		}
	}
}
