package particle

import (
	"fmt"
	"image/color"
)

type picture struct {
	imageOf [5][5]color.RGBA
}

func newPicture() *picture {

	p := picture{}

	for i := range p.imageOf {
		for t := range p.imageOf[i] {
			p.imageOf[i][t].R = 0xff
			p.imageOf[i][t].G = 0xff
			p.imageOf[i][t].B = 0xff
			p.imageOf[i][t].A = 0xff
		}
	}

	return &p
}

type Particle struct {
}

type Drawable interface {
	render() picture
}

func (a Particle) String() string {
	return fmt.Sprintf("%d", a.Render().imageOf[1][1])
}

func (r Particle) Render() picture {
	rendering := newPicture()
	return *rendering
}
