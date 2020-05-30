package geometry

import "image/color"

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func (cp *ColoredPoint) Distance(cp2 *ColoredPoint) float64 {
	return cp.Point.Distance(cp2.Point)
}

func (cp *ColoredPoint) Distance2(p *Point) float64 {
	return cp.Point.Distance(*p)
}
