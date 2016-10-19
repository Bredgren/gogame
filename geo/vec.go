package geo

import (
	"math"
	"math/rand"

	"github.com/Bredgren/wrand"
)

// Vec is a 2D vector. Many of the functions for Vec have two versions, one that modifies
// the Vec and one that returns a new Vec. Their names follow a convention that is hopefully
// inuitive. For example, when working with Vec as a value you use v1.Plus(v2) which returns
// a new value and reads like when working with other value types such as "1 + 2".
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

// SetLen sets the length of the vector. Negative lengths will flip the vectors direction.
// If v's length is zero then it will remain unchanged.
func (v *Vec) SetLen(l float64) {
	len := math.Sqrt(v.X*v.X + v.Y*v.Y)
	if len == 0 {
		return
	}
	v.X = v.X / len * l
	v.Y = v.Y / len * l
}

// WithLen returns a new vector in the same direction as v but with the given length.
// It v's length is 0 then the zero vector is returned.
func (v Vec) WithLen(l float64) Vec {
	len := math.Sqrt(v.X*v.X + v.Y*v.Y)
	if len == 0 {
		return Vec{}
	}
	v.X = v.X / len * l
	v.Y = v.Y / len * l
	return v
}

// Dist returns the distance between the two vectors.
func (v Vec) Dist(v2 Vec) float64 {
	return math.Sqrt((v.X-v2.X)*(v.X-v2.X) + (v.Y-v2.Y)*(v.Y-v2.Y))
}

// Dist2 returns the distance squared between the two vectors.
func (v Vec) Dist2(v2 Vec) float64 {
	return (v.X-v2.X)*(v.X-v2.X) + (v.Y-v2.Y)*(v.Y-v2.Y)
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

// Div modifies v to be itself divided by n.
func (v *Vec) Div(n float64) {
	v.X /= n
	v.Y /= n
}

// DividedBy returns a new vector that is v divided by n.
func (v Vec) DividedBy(n float64) Vec {
	return Vec{X: v.X / n, Y: v.Y / n}
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

// Project modifies v to be the vector that is the projection of v onto v2.
func (v *Vec) Project(v2 Vec) {
	v2.Normalize()
	v2.Mul(v.X*v2.X + v.Y*v2.Y)
	v.X, v.Y = v2.X, v2.Y
}

// Projected return the vetor that is v projected onto v2.
func (v Vec) Projected(v2 Vec) Vec {
	v2.Normalize()
	v2.Mul(v.X*v2.X + v.Y*v2.Y)
	return v2
}

// Limit constrains the length of the vector be no greater than len. If the vector is already
// less than len then no change is made.
func (v *Vec) Limit(len float64) {
	l := v.Len()
	if l > len {
		v.SetLen(len)
	}
}

// Limited returns a new vector in the same direction as v with length no greater than
// len. The vector returned will be equivalent to v if v.Len() <= len.
func (v Vec) Limited(len float64) Vec {
	l := v.Len()
	if l > len {
		v.SetLen(len)
	}
	return v
}

// Angle returns the radians relative to the positive x-axis (counterclockwise in screen
// coordinates). The returned value is in the range [-π, π).
func (v Vec) Angle() float64 {
	return -math.Atan2(v.Y, v.X)
}

// AngleFrom returns the radians from v2 to v (counterclockwise in screen coordinates).
// The returned value is in the range [-π, π).
func (v Vec) AngleFrom(v2 Vec) float64 {
	r := math.Atan2(v2.Y, v2.X) - math.Atan2(v.Y, v.X)
	if r < -math.Pi {
		r += 2 * math.Pi
	}
	return r
}

// Rotate rotates the vector (counterclockwise in screen coordinates) by the given radians.
func (v *Vec) Rotate(rad float64) {
	v.X, v.Y = v.X*math.Cos(rad)+v.Y*math.Sin(rad), -v.X*math.Sin(rad)+v.Y*math.Cos(rad)
}

// Rotated returns a new vector that is equal to this one rotated (counterclockwise in
// screen coordinates) by the given radians.
func (v Vec) Rotated(rad float64) Vec {
	return Vec{X: v.X*math.Cos(rad) + v.Y*math.Sin(rad), Y: -v.X*math.Sin(rad) + v.Y*math.Cos(rad)}
}

// Equals returns true if the corresponding components of the vectors are within the error e.
func (v Vec) Equals(v2 Vec, e float64) bool {
	return math.Abs(v.X-v2.X) < e && math.Abs(v.Y-v2.Y) < e
}

// VecGen (Vector Generator) is a function that returns a vector.
type VecGen func() Vec

// RandVec returns a unit vector in a random direction.
func RandVec() Vec {
	rad := rand.Float64() * 2 * math.Pi
	return Vec{X: math.Cos(rad), Y: math.Sin(rad)}
}

// StaticVec returns a VecGen that always returns the constant vector v.
func StaticVec(v Vec) VecGen {
	return func() Vec {
		return v
	}
}

// DynamicVec returns a VecGen that always returns a copy of the vector pointed to by v.
func DynamicVec(v *Vec) VecGen {
	return func() Vec {
		return *v
	}
}

// OffsetVec returns a VecGen that adds offset to gen. For example, if you wanted to use
// RandCircle as an initial position you might use
//  OffsetVec(RandVecCircle(5, 10), StaticVec(Vec{X: 100, Y: 100}))
// to center the circle at position 100, 100.
func OffsetVec(gen VecGen, offset VecGen) VecGen {
	return func() Vec {
		return gen().Plus(offset())
	}
}

// RandVecCircle returns a VecGen that will generate a random vector within the given radii.
// Negative radii are undefined.
func RandVecCircle(minRadius, maxRadius float64) VecGen {
	return func() Vec {
		return RandVec().Times(circleRadius(minRadius, maxRadius))
	}
}

// RandVecArc returns a VecGen that will generate a random vector within the slice of a
// circle defined by the parameters. The radians are relative to the +x axis.
// Negative radii are undefined.
func RandVecArc(minRadius, maxRadius, minRadians, maxRadians float64) VecGen {
	if maxRadians < minRadians {
		minRadians, maxRadians = maxRadians, minRadians
	}
	return func() Vec {
		r := circleRadius(minRadius, maxRadius)
		rad := rand.Float64()*(maxRadians-minRadians) + minRadians
		return Vec{X: r}.Rotated(rad)
	}
}

// RandVecRect returns a VecGen that will generate a random vector within the given Rect.
func RandVecRect(rect Rect) VecGen {
	return func() Vec {
		return Vec{
			X: rand.Float64()*rect.W + rect.X,
			Y: rand.Float64()*rect.H + rect.Y,
		}
	}
}

// RandVecRects returns a VecGen that will generate a random vector that is uniformly
// distributed between all the given rects. If the slice given is empty then the zero
// vector is returned.
func RandVecRects(rects []Rect) VecGen {
	if len(rects) == 0 {
		return func() Vec { return Vec{} }
	}
	areas := make([]float64, len(rects))
	for i := range rects {
		areas[i] = rects[i].Area()
	}
	return func() Vec {
		rect := rects[wrand.SelectIndex(areas)]
		return Vec{
			X: rand.Float64()*rect.W + rect.X,
			Y: rand.Float64()*rect.H + rect.Y,
		}
	}
}

// Returns a uniformaly distributed radius between minR and maxR.
func circleRadius(minR, maxR float64) float64 {
	if maxR == 0 || maxR == minR {
		return maxR
	}
	unitMin := minR / maxR
	unitMin *= unitMin
	return math.Sqrt(rand.Float64()*(1-unitMin)+unitMin) * maxR
}
