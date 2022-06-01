package particle

import (
	"fmt"
	"image/color"
)

type Picture struct {
	ImageOf [5][5]color.RGBA
}

type Position struct {
	X int
	Y int
}

func newPicture() *Picture {

	p := Picture{
		ImageOf: [5][5]color.RGBA{},
	}

	for i := range p.ImageOf {
		for t := range p.ImageOf[i] {
			p.ImageOf[i][t].R = 0x0f
			p.ImageOf[i][t].G = 0xff
			p.ImageOf[i][t].B = 0xff
			p.ImageOf[i][t].A = 0x0f
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
