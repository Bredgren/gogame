package gogame

import (
	"fmt"
	"math"
)

// Color is an rgba color
type Color struct {
	R, G, B, A float64
}

var (
	// Black is the color black
	Black = Color{0.0, 0.0, 0.0, 1.0}
	// White is the color white
	White = Color{1.0, 1.0, 1.0, 1.0}
)

// GetColor creates a Color with the specified values. All values should be between 0.0
// and 1.0 (inclusive). An alpha of 0.0 is transparent.
func (c *Color) String() string {
	return fmt.Sprintf("rgba(%d, %d, %d, %f)",
		clampToInt(255*c.R, 0, 255),
		clampToInt(255*c.G, 0, 255),
		clampToInt(255*c.B, 0, 255),
		clampToFloat(c.A, 0.0, 1.0))
}

func clampToInt(v, min, max float64) int {
	return int(math.Min(math.Max(v, min), max))
}

func clampToFloat(v, min, max float64) float64 {
	return math.Min(math.Max(v, min), max)
}
