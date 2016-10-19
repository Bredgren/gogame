package ggweb

// DrawType specifies the method used for drawing.
type DrawType string

const (
	// Fill draws within the shape's boundaries.
	Fill DrawType = "fill"
	// Stroke draws only the shape's boundaries.
	Stroke DrawType = "stroke"
)
