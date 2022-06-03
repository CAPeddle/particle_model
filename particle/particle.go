package particle

import (
	"fmt"
	"image/color"
)

const (
	ParticleWidth  = 5
	ParticleHeight = 5
)

type Picture struct {
	ImageOf [ParticleHeight][ParticleWidth]color.RGBA
}

type Position struct {
	X int
	Y int
}

func newPicture() *Picture {

	p := Picture{
		//[row][column]
		ImageOf: [ParticleHeight][ParticleWidth]color.RGBA{},
	}

	for i := range p.ImageOf {
		for t := range p.ImageOf[i] {
			p.ImageOf[i][t].R = 245
			p.ImageOf[i][t].G = 40
			p.ImageOf[i][t].B = 145
			p.ImageOf[i][t].A = 0
		}
	}

	return &p
}

type Particle struct {
	Alive    bool
	Location Position
}

// type Drawable interface {
// 	Render() picture
// }

func (a Particle) String() string {
	return fmt.Sprintf("%d", a.Render().ImageOf[1][1])
}

func (r Particle) Render() Picture {
	rendering := newPicture()
	return *rendering
}
