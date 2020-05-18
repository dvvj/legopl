package geometry

import "testing"

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
