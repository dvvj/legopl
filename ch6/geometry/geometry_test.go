package geometry

import (
	"image/color"
	"testing"
)

func TestDistance(t *testing.T) {
	p := Point{X: 0, Y: 0}
	q := Point{X: 3, Y: 4}

	d := Distance(p, q)
	if d != 5.0 {
		t.Error("Distance( 0,0 - 3,4 ) != 5")
	}
}

func TestPointDistance(t *testing.T) {
	p := Point{X: 0, Y: 0}

	var testdata = []struct {
		input Point
		exp   float64
	}{
		{Point{3, 4}, 5.0},
		{Point{6, 8}, 10.0},
	}

	for _, q := range testdata {
		res := p.Distance(q.input)
		if res != q.exp {
			t.Errorf("Distance %v != %v\n", q.input, q.exp)
		}
	}
}

func TestPathDistance(t *testing.T) {
	var testdata = []struct {
		path Path
		exp  float64
	}{
		{
			Path{Point{0, 0}, Point{3, 4}, Point{6, 8}},
			10.0,
		}, {
			Path{Point{0, 0}, Point{3, 4}, Point{6, 8}},
			11.0,
		},
	}

	for _, p := range testdata {
		res := p.path.Distance()
		if res != p.exp {
			t.Errorf("Distance %v != %v\n", p.path, p.exp)
		}
	}
}

func TestColoredPoint(t *testing.T) {
	var cp ColoredPoint

	cp.X = 3
	cp.Point.Y = 4

	if cp.Point.X != 3 || cp.Y != 4 {
		t.Error("err")
	}

	cp2 := ColoredPoint{Point{6, 8}, color.RGBA{255, 0, 0, 255}}
	dist := cp.Distance(&cp2)
	if dist != 5 {
		t.Error("err")
	}

	dist = cp.Distance2(&cp2.Point)
	if dist != 5 {
		t.Error("err")
	}

}
