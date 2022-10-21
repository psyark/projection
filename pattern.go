package projection

import (
	"image"
	"image/color"

	"github.com/fogleman/gg"
	"github.com/psyark/projection/transform"
)

var _ gg.Pattern = &Pattern{}

type Pattern struct {
	texture image.Image
	w, h    float64
	mesh    [4]transform.Point
	inv     transform.Matrix
}

func (p *Pattern) ColorAt(x int, y int) color.Color {
	tx, ty := p.inv.Transform(float64(x), float64(y))
	return p.texture.At(int(p.w*tx), int(p.h*ty))
}

func NewPattern(texture image.Image, mesh [4]transform.Point) *Pattern {
	mat := transform.NewDeformation(
		mesh[0],
		mesh[1],
		mesh[3],
		mesh[2],
	)

	return &Pattern{
		texture: texture,
		w:       float64(texture.Bounds().Dx()),
		h:       float64(texture.Bounds().Dy()),
		mesh:    mesh,
		inv:     mat.Inverse(),
	}
}
