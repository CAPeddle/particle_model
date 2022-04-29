package particle

import (
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

type drawable interface {
	render() picture
}

func (r Particle) render() picture {
	rendering := newPicture()
	return *rendering
}
