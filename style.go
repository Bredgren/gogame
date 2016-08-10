package gogame

// FillStyle specifies a style used for filling shapes
type FillStyle struct {
	Color
}

// LineCap is a style of line cap
type LineCap string

const (
	// LineCapButt draws a line with no ends
	LineCapButt LineCap = "butt"
	// LineCapRound draws a line with rounded ends with radius equal to half its width
	LineCapRound LineCap = "round"
	// LineCapSquare draws a line with the ends capped with a box that extends by an amount
	// equal to half the lines width
	LineCapSquare LineCap = "square"
)

// LineJoin is the style for the point where two lines are connected
type LineJoin string

const (
	// LineJoinRound joins lines with rounded corners
	LineJoinRound LineJoin = "round"
	// LineJoinBevel joins lines by filling in the triangular gap between them
	LineJoinBevel LineJoin = "bevel"
	// LineJoinMiter joins lines by extending the edges until they meet
	LineJoinMiter LineJoin = "miter"
)

// StrokeStyle specifies a style used for lines
type StrokeStyle struct {
	Color
	Width      float64
	Cap        LineCap
	Join       LineJoin
	MiterLimit float64
	Dash       []float64
	DashOffest float64
}
