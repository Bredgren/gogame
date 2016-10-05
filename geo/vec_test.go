package geo

import (
	"math"
	"testing"
)

func TestVecLen(t *testing.T) {
	cases := []struct {
		v    Vec
		want float64
	}{
		{Vec{X: 3, Y: 4}, 5},
	}

	for i, c := range cases {
		got := c.v.Len()
		if got != c.want {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestVecLen2(t *testing.T) {
	cases := []struct {
		v    Vec
		want float64
	}{
		{Vec{X: 3, Y: 4}, 25},
	}

	for i, c := range cases {
		got := c.v.Len2()
		if got != c.want {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestVecSetLen(t *testing.T) {
	cases := []struct {
		v    Vec
		want float64
	}{
		{Vec{X: 3, Y: 4}, 10},
	}

	for i, c := range cases {
		c.v.SetLen(c.want)
		got := c.v.Len()
		if got != c.want {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestVecAdd(t *testing.T) {
	cases := []struct {
		v1   Vec
		v2   Vec
		want Vec
	}{
		{Vec{X: 3, Y: -4}, Vec{X: -3, Y: 4}, Vec{X: 0, Y: 0}},
		{Vec{X: 3, Y: 4}, Vec{X: 3, Y: 4}, Vec{X: 6, Y: 8}},
	}

	for i, c := range cases {
		got := c.v1
		got.Add(c.v2)
		if got != c.want {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestVecPlus(t *testing.T) {
	cases := []struct {
		v1   Vec
		v2   Vec
		want Vec
	}{
		{Vec{X: 3, Y: -4}, Vec{X: -3, Y: 4}, Vec{X: 0, Y: 0}},
		{Vec{X: 3, Y: 4}, Vec{X: 3, Y: 4}, Vec{X: 6, Y: 8}},
	}

	for i, c := range cases {
		got := c.v1.Plus(c.v2)
		if got != c.want {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestVecSub(t *testing.T) {
	cases := []struct {
		v1   Vec
		v2   Vec
		want Vec
	}{
		{Vec{X: 3, Y: -4}, Vec{X: -3, Y: 4}, Vec{X: 6, Y: -8}},
		{Vec{X: 3, Y: 4}, Vec{X: 3, Y: 4}, Vec{X: 0, Y: 0}},
	}

	for i, c := range cases {
		got := c.v1
		got.Sub(c.v2)
		if got != c.want {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestVecMinus(t *testing.T) {
	cases := []struct {
		v1   Vec
		v2   Vec
		want Vec
	}{
		{Vec{X: 3, Y: -4}, Vec{X: -3, Y: 4}, Vec{X: 6, Y: -8}},
		{Vec{X: 3, Y: 4}, Vec{X: 3, Y: 4}, Vec{X: 0, Y: 0}},
	}

	for i, c := range cases {
		got := c.v1.Minus(c.v2)
		if got != c.want {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestVecMul(t *testing.T) {
	cases := []struct {
		v    Vec
		n    float64
		want Vec
	}{
		{Vec{X: 3, Y: -4}, 2, Vec{X: 6, Y: -8}},
		{Vec{X: 3, Y: 4}, 0, Vec{X: 0, Y: 0}},
	}

	for i, c := range cases {
		got := c.v
		got.Mul(c.n)
		if got != c.want {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestVecTimes(t *testing.T) {
	cases := []struct {
		v    Vec
		n    float64
		want Vec
	}{
		{Vec{X: 3, Y: -4}, 2, Vec{X: 6, Y: -8}},
		{Vec{X: 3, Y: 4}, 0, Vec{X: 0, Y: 0}},
	}

	for i, c := range cases {
		got := c.v.Times(c.n)
		if got != c.want {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestVecNormalize(t *testing.T) {
	cases := []struct {
		v    Vec
		want Vec
	}{
		{Vec{X: 5, Y: 0}, Vec{X: 1, Y: 0}},
		{Vec{X: 0, Y: -4}, Vec{X: 0, Y: -1}},
	}

	for i, c := range cases {
		got := c.v
		got.Normalize()
		if got != c.want {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestVecNormalized(t *testing.T) {
	cases := []struct {
		v    Vec
		want Vec
	}{
		{Vec{X: 5, Y: 0}, Vec{X: 1, Y: 0}},
		{Vec{X: 0, Y: -4}, Vec{X: 0, Y: -1}},
	}

	for i, c := range cases {
		got := c.v.Normalized()
		if got != c.want {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestVecDot(t *testing.T) {
	cases := []struct {
		v1, v2 Vec
		want   float64
	}{
		{Vec{X: 5, Y: 0}, Vec{X: 1, Y: 0}, 5},
		{Vec{X: 0, Y: -4}, Vec{X: 0, Y: -1}, 4},
		{Vec{X: 1, Y: 2}, Vec{X: 2, Y: -1}, 0},
	}

	for i, c := range cases {
		got := c.v1.Dot(c.v2)
		if got != c.want {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestVecRand(t *testing.T) {
	cases := 100000
	maxErr := 1e-10
	for i := 0; i < cases; i++ {
		got := RandVec()
		if math.Abs(got.Len()-1) > maxErr {
			t.Errorf("case %d: %#v is length %f", i, got, got.Len())
		}
	}
}

func TestVecLimit(t *testing.T) {
	cases := []struct {
		v    Vec
		len  float64
		want Vec
	}{
		{Vec{X: 5, Y: 0}, 2, Vec{X: 2, Y: 0}},
		{Vec{X: 0, Y: -4}, 2, Vec{X: 0, Y: -2}},
		{Vec{X: 3, Y: 4}, 6, Vec{X: 3, Y: 4}},
	}

	for i, c := range cases {
		got := c.v
		got.Limit(c.len)
		if got != c.want {
			t.Errorf("case %d: len: %#v, got %#v, want %#v", i, c.len, got, c.want)
		}
	}
}

func TestVecAngle(t *testing.T) {
	cases := []struct {
		v    Vec
		want float64
	}{
		{Vec{X: 5, Y: 0}, 0},
		{Vec{X: -3, Y: 0}, math.Pi},
		{Vec{X: 0, Y: -4}, -math.Pi / 2},
		{Vec{X: 0, Y: 1}, math.Pi / 2},
	}

	for i, c := range cases {
		got := c.v.Angle()
		if math.Abs(got-c.want) > 1e-10 {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestVecAngleFrom(t *testing.T) {
	cases := []struct {
		v1, v2 Vec
		want   float64
	}{
		{Vec{X: 5, Y: 0}, Vec{X: 2, Y: 0}, 0},
		{Vec{X: -5, Y: 0}, Vec{X: -2, Y: 0}, 0},
		{Vec{X: 0, Y: -4}, Vec{X: -1, Y: -1}, math.Pi / 4},
		{Vec{X: -1, Y: 0}, Vec{X: 1, Y: 0}, math.Pi},
		{Vec{X: -1, Y: 0}, Vec{X: 0, Y: -1}, -math.Pi / 2},
	}

	for i, c := range cases {
		got := c.v1.AngleFrom(c.v2)
		if math.Abs(got-c.want) > 1e-10 {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestVecRotate(t *testing.T) {
	cases := []struct {
		v    Vec
		rad  float64
		want Vec
	}{
		{Vec{X: 5, Y: 0}, 0, Vec{X: 5, Y: 0}},
		{Vec{X: 3, Y: 0}, math.Pi / 2, Vec{X: 0, Y: 3}},
		{Vec{X: 3, Y: 0}, -math.Pi / 2, Vec{X: 0, Y: -3}},
		{Vec{X: 3, Y: 0}, math.Pi, Vec{X: -3, Y: 0}},
		{Vec{X: 3, Y: 0}, -math.Pi, Vec{X: -3, Y: 0}},
		{Vec{X: 0, Y: -1}, math.Pi, Vec{X: 0, Y: 1}},
		{Vec{X: 0, Y: -1}, math.Pi / 4, Vec{X: 1, Y: -1}.Normalized()},
		{Vec{X: -1, Y: 0}, -math.Pi, Vec{X: 1, Y: 0}},
		{Vec{X: -1, Y: 0}, 3 * math.Pi / 2, Vec{X: 0, Y: 1}},
	}

	for i, c := range cases {
		got := c.v
		got.Rotate(c.rad)
		if !vecsEqual(got, c.want) {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestVecRotated(t *testing.T) {
	cases := []struct {
		v    Vec
		rad  float64
		want Vec
	}{
		{Vec{X: 5, Y: 0}, 0, Vec{X: 5, Y: 0}},
		{Vec{X: 3, Y: 0}, math.Pi / 2, Vec{X: 0, Y: 3}},
		{Vec{X: 3, Y: 0}, -math.Pi / 2, Vec{X: 0, Y: -3}},
		{Vec{X: 3, Y: 0}, math.Pi, Vec{X: -3, Y: 0}},
		{Vec{X: 3, Y: 0}, -math.Pi, Vec{X: -3, Y: 0}},
		{Vec{X: 0, Y: -1}, math.Pi, Vec{X: 0, Y: 1}},
		{Vec{X: 0, Y: -1}, math.Pi / 4, Vec{X: 1, Y: -1}.Normalized()},
		{Vec{X: -1, Y: 0}, -math.Pi, Vec{X: 1, Y: 0}},
		{Vec{X: -1, Y: 0}, 3 * math.Pi / 2, Vec{X: 0, Y: 1}},
	}

	for i, c := range cases {
		got := c.v.Rotated(c.rad)
		if !vecsEqual(got, c.want) {
			t.Errorf("case %d: got %#v, want %#v", i, got, c.want)
		}
	}
}

func TestVecRotateStress(t *testing.T) {
	cases := 100000
	for i := 0; i < cases; i++ {
		v1 := RandVec()
		v2 := RandVec()
		between := v1.AngleFrom(v2)
		rotated := v2.Rotated(between)
		v2rotated := v2
		v2rotated.Rotate(between)
		if !vecsEqual(v1, rotated) || !vecsEqual(v1, v2rotated) {
			t.Errorf("case %d: v1: %#v v2: %#v between: %#v rotated: %#v v2rotated: %#v",
				i, v1, v2, between, rotated, v2rotated)
		}
	}
}

func vecsEqual(v1, v2 Vec) bool {
	e := 1e-10
	return math.Abs(v1.X-v2.X) < e && math.Abs(v1.Y-v2.Y) < e
}
