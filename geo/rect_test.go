package geo

import "testing"

func TestRectInflate(t *testing.T) {
	cases := []struct {
		r, want Rect
		dw, dh  float64
	}{
		{
			r:  Rect{X: 1, Y: 1, W: 5, H: 5},
			dw: 2, dh: 2,
			want: Rect{X: 0, Y: 0, W: 7, H: 7},
		},
		{
			r:  Rect{X: 1, Y: 1, W: 5, H: 5},
			dw: -1, dh: -1,
			want: Rect{X: 1.5, Y: 1.5, W: 4, H: 4},
		},
	}

	for i, c := range cases {
		got := c.r.Inflate(c.dw, c.dh)
		if got != c.want {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}

		c.r.InflateIP(c.dw, c.dh)
		if c.r != c.want {
			t.Errorf("IP case %d: got %#v, want %#v", i, c.r, c.want)
		}
	}
}

func TestRectClamp(t *testing.T) {
	bounds := Rect{X: 1, Y: 1, W: 5, H: 5}
	cases := []struct {
		bounds, r, want Rect
	}{
		{
			bounds: bounds,
			r:      Rect{X: 0, Y: 0, W: 1, H: 1},
			want:   Rect{X: 1, Y: 1, W: 1, H: 1},
		},
		{
			bounds: bounds,
			r:      Rect{X: 7, Y: 6, W: 1, H: 1},
			want:   Rect{X: 5, Y: 5, W: 1, H: 1},
		},
		{
			bounds: bounds,
			r:      Rect{X: 7, Y: 6, W: 7, H: 7},
			want:   Rect{X: 0, Y: 0, W: 7, H: 7},
		},
	}

	for i, c := range cases {
		got := c.r.Clamp(&c.bounds)
		if got != c.want {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}

		c.r.ClampIP(&c.bounds)
		if c.r != c.want {
			t.Errorf("IP case %d: got %#v, want %#v", i, c.r, c.want)
		}
	}
}

func TestRectIntersect(t *testing.T) {
	cases := []struct {
		r1, r2, want Rect
	}{
		{
			r1:   Rect{X: 1, Y: 1, W: 5, H: 5},
			r2:   Rect{X: 0, Y: 0, W: 2, H: 3},
			want: Rect{X: 1, Y: 1, W: 1, H: 2},
		},
		{
			r1:   Rect{X: 1, Y: 1, W: 5, H: 5},
			r2:   Rect{X: 2, Y: 3, W: 4, H: 5},
			want: Rect{X: 2, Y: 3, W: 4, H: 3},
		},
		{
			r1:   Rect{X: 1, Y: 1, W: 5, H: 5},
			r2:   Rect{X: 2, Y: 2, W: 2, H: 2},
			want: Rect{X: 2, Y: 2, W: 2, H: 2},
		},
		{
			r1:   Rect{X: 1, Y: 1, W: 5, H: 5},
			r2:   Rect{X: 6, Y: 6, W: 2, H: 2},
			want: Rect{X: 6, Y: 6, W: 0, H: 0},
		},
	}

	for i, c := range cases {
		got := c.r1.Intersect(&c.r2)
		if got != c.want {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}

		got = c.r2.Intersect(&c.r1)
		if got != c.want {
			t.Errorf("reverse case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

// func TestRectUnion(t *testing.T) {
// }

// func TestRectUnionIP(t *testing.T) {
// }

// func TestRectUnionAll(t *testing.T) {
// }

// func TestRectFit(t *testing.T) {
// }

// func TestRectNormalize(t *testing.T) {
// }

// func TestRectContains(t *testing.T) {
// }

// func TestRectCollidePoint(t *testing.T) {
// }

// func TestRectCollideRect(t *testing.T) {
// }
