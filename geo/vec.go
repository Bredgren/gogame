package geo

import (
	"math"
	"math/rand"
)

// Vec is a 2D vector.
type Vec struct {
	X, Y float64
}

// Len returns the length of the vector.
func (v Vec) Len() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Len2 is the length of the vector squared.
func (v Vec) Len2() float64 {
	return v.X*v.X + v.Y*v.Y
}

// SetLen sets the length of the vector.
func (v *Vec) SetLen(l float64) {
	len := math.Sqrt(v.X*v.X + v.Y*v.Y)
	v.X = v.X / len * l
	v.Y = v.Y / len * l
}

// Add modifies v to be the sum of v2 and itself.
func (v *Vec) Add(v2 Vec) {
	v.X += v2.X
	v.Y += v2.Y
}

// Plus returns a new vector that is the sum of the two vectors.
func (v Vec) Plus(v2 Vec) Vec {
	return Vec{X: v.X + v2.X, Y: v.Y + v2.Y}
}

// Sub modifies v to be the difference between itself and v2.
func (v *Vec) Sub(v2 Vec) {
	v.X -= v2.X
	v.Y -= v2.Y
}

// Minus returns a new vector that is the difference of the two vectors.
func (v Vec) Minus(v2 Vec) Vec {
	return Vec{X: v.X - v2.X, Y: v.Y - v2.Y}
}

// Mul modifies v to be itself times n.
func (v *Vec) Mul(n float64) {
	v.X *= n
	v.Y *= n
}

// Times returns a new vector that is v times n.
func (v Vec) Times(n float64) Vec {
	return Vec{X: v.X * n, Y: v.Y * n}
}

// Normalize modifies v to be of length one in the same direction.
func (v *Vec) Normalize() {
	len := math.Sqrt(v.X*v.X + v.Y*v.Y)
	v.X = v.X / len
	v.Y = v.Y / len
}

// Normalized returns a new vector of length one in the same direction as v.
func (v Vec) Normalized() Vec {
	len := math.Sqrt(v.X*v.X + v.Y*v.Y)
	return Vec{X: v.X / len, Y: v.Y / len}
}

// Dot returns the dot product between the two vectors.
func (v Vec) Dot(v2 Vec) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

// Limit constrains the length of the vector be no greater than len. If the vector is already
// less than len then no change is made.
func (v *Vec) Limit(len float64) {
	l := v.Len()
	if l > len {
		v.SetLen(len)
	}
}

// Angle returns the radians relative to the positive x-axis (counterclockwise in screen
// coordinates). The returned values is in the range (-π, π].
func (v Vec) Angle() float64 {
	return math.Atan2(v.Y, v.X)
}

// AngleFrom returns the radians from v2 to v (counterclockwise in screen coordinates).
// The returned values is in the range (-π, π].
func (v Vec) AngleFrom(v2 Vec) float64 {
	r := math.Atan2(v.Y, v.X) - math.Atan2(v2.Y, v2.X)
	if r > math.Pi {
		r -= 2 * math.Pi
	}
	return r
}

// Rotate rotates the vector (counterclockwise in screen coordinates) by the given radians.
func (v *Vec) Rotate(rad float64) {
	v.X, v.Y = v.X*math.Cos(rad)-v.Y*math.Sin(rad), v.X*math.Sin(rad)+v.Y*math.Cos(rad)
}

// Rotated returns a new vector that is equal to this one rotated (counterclockwise in
// screen coordinates) by the given radians.
func (v Vec) Rotated(rad float64) Vec {
	return Vec{X: v.X*math.Cos(rad) - v.Y*math.Sin(rad), Y: v.X*math.Sin(rad) + v.Y*math.Cos(rad)}
}

// Equals returns true if both components of both vectors are within the error e.
func (v Vec) Equals(v2 Vec, e float64) bool {
	return math.Abs(v.X-v2.X) < e && math.Abs(v.Y-v2.Y) < e
}

// RandVec returns a unit vector in a random direction.
func RandVec() Vec {
	rad := rand.Float64() * 2 * math.Pi
	return Vec{X: math.Cos(rad), Y: math.Sin(rad)}
}
