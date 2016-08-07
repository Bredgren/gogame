package gogame

// Rect defines rectangular coordinates.
type Rect struct {
	X, Y, W, H int
}

// Top returns the top boundary
func (r *Rect) Top() (top int) {
	return r.Y
}

// SetTop sets the top boundary
func (r *Rect) SetTop(top int) {
	r.Y = top
}

// Bottom returns the bottom boundary
func (r *Rect) Bottom() int {
	return r.Y + r.H
}

// SetBottom sets the bottom boundary
func (r *Rect) SetBottom(bottom int) {
	r.Y = bottom - r.H
}

// Left returns the left boundary
func (r *Rect) Left() int {
	return r.X
}

// SetLeft sets the left boundary
func (r *Rect) SetLeft(left int) {
	r.X = left
}

// Right returns the right boundary
func (r *Rect) Right() int {
	return r.X + r.W
}

// SetRight sets the right boundary
func (r *Rect) SetRight(right int) {
	r.X = right - r.W
}

// Width returns the width
func (r *Rect) Width() int {
	return r.W
}

// SetWidth sets the width
func (r *Rect) SetWidth(w int) {
	r.W = w
}

// Height returns the height
func (r *Rect) Height() int {
	return r.H
}

// SetHeight sets the height
func (r *Rect) SetHeight(h int) {
	r.H = h
}

// Size returns the width and height
func (r *Rect) Size() (w, h int) {
	return r.W, r.H
}

// SetSize sets the width and height
func (r *Rect) SetSize(w, h int) {
	r.SetWidth(w)
	r.SetHeight(h)
}

// TopLeft returns the coordinates of the top left corner
func (r *Rect) TopLeft() (x, y int) {
	return r.Left(), r.Top()
}

// SetTopLeft sets the coordinates of the top left corner
func (r *Rect) SetTopLeft(x, y int) {
	r.SetLeft(x)
	r.SetRight(y)
}

// BottomLeft returns the coordinates of the bottom left corner
func (r *Rect) BottomLeft() (x, y int) {
	return r.Left(), r.Bottom()
}

// SetBottomLeft set the coordinates of the bottom left corner
func (r *Rect) SetBottomLeft(x, y int) {
	r.SetLeft(x)
	r.SetBottom(y)
}

// TopRight returns the coordinates of the top right corner
func (r *Rect) TopRight() (x, y int) {
	return r.Right(), r.Top()
}

// SetTopRight sets the coordinates of the top right corner
func (r *Rect) SetTopRight(x, y int) {
	r.SetRight(x)
	r.SetTop(y)
}

// BottomRight returns the coordinates of the bottom right corner
func (r *Rect) BottomRight() (x, y int) {
	return r.Right(), r.Bottom()
}

// SetBottomRight sets the coordinates of the bottom right corner
func (r *Rect) SetBottomRight(x, y int) {
	r.SetRight(x)
	r.SetBottom(y)
}

// MidTop returns the coordinates at the top of the rectangle above the center
func (r *Rect) MidTop() (x, y int) {
	return r.CenterX(), r.Top()
}

// SetMidTop sets the coordinates at the top of the rectangle above the center
func (r *Rect) SetMidTop(x, y int) {
	r.SetCenterX(x)
	r.SetTop(y)
}

// MidBottom returns the coordinates at the bottom of the rectangle bellow the center
func (r *Rect) MidBottom() (x, y int) {
	return r.CenterX(), r.Bottom()
}

// SetMidBottom sets the coordinates at the bottom of the rectangle bellow the center
func (r *Rect) SetMidBottom(x, y int) {
	r.SetCenterX(x)
	r.SetBottom(y)
}

// MidLeft returns the coordinates at the left of the rectangle in line with the center
func (r *Rect) MidLeft() (x, y int) {
	return r.Left(), r.CenterY()
}

// SetMidLeft sets the coordinates at the left of the rectangle in line with the center
func (r *Rect) SetMidLeft(x, y int) {
	r.SetLeft(x)
	r.SetCenterY(y)
}

// MidRight returns the coordinates at the right of the rectangle in line with the center
func (r *Rect) MidRight() (x, y int) {
	return r.Right(), r.CenterY()
}

// SetMidRight sets the coordinates at the right of the rectangle in line with the center
func (r *Rect) SetMidRight(x, y int) {
	r.SetRight(x)
	r.SetCenterY(y)
}

// Center returns the center coordinates
func (r *Rect) Center() (x, y int) {
	return r.CenterX(), r.CenterY()
}

// SetCenter sets the center coordinates
func (r *Rect) SetCenter(x, y int) {
	r.SetCenterX(x)
	r.SetCenterY(y)
}

// CenterX returns the center x coordinates
func (r *Rect) CenterX() int {
	return r.X + r.W/2
}

// SetCenterX sets the center x coordinates
func (r *Rect) SetCenterX(x int) {
	r.X = x - r.W/2
}

// CenterY returns the center y coordinates
func (r *Rect) CenterY() int {
	return r.Y + r.H/2
}

// SetCenterY set the center y coordinates
func (r *Rect) SetCenterY(y int) {
	r.Y = y - r.H/2
}
