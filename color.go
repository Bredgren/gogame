package gogame

import (
	"fmt"
	"math"
)

// Color is an rgba color
type Color string

const (
	// Black is the color black
	Black Color = "rgba(0, 0, 0, 1.0)"
)

// GetColor creates a Color with the specified values. All values should be between 0.0
// and 1.0 (inclusive). An alpha of 0.0 is transparent.
func GetColor(r, g, b, a float64) Color {
	return Color(fmt.Sprintf("rgba(%d, %d, %d, %f)",
		clampToInt(255*r, 0, 255),
		clampToInt(255*g, 0, 255),
		clampToInt(255*b, 0, 255),
		clampToFloat(a, 0.0, 1.0)))
}

func clampToInt(v, min, max float64) int {
	return int(math.Min(math.Max(v, min), max))
}

func clampToFloat(v, min, max float64) float64 {
	return math.Min(math.Max(v, min), max)
}
