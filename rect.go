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

// Copy returns a new Rect that is idential to this one
func (r *Rect) Copy() Rect {
	return Rect{X: r.X, Y: r.Y, W: r.W, H: r.H}
}

// Move returns a new Rect moved by the given offset relative to this one
func (r *Rect) Move(dx, dy int) Rect {
	return Rect{X: r.X + dx, Y: r.Y + dy, W: r.W, H: r.H}
}

// MoveIP moves the Rect by the given offset, in place
func (r *Rect) MoveIP(dx, dy int) {
	r.X += dx
	r.Y += dy
}

// Inflate returns a new Rect with the same center whose size is chaged by the given amount
func (r *Rect) Inflate(dw, dh int) Rect {
	return Rect{X: r.X - dw/2, Y: r.Y - dh/2, W: r.W + dw, H: r.H + dh}
}

// InflateIP keeps the same center but changes the size by the given amount, in place
func (r *Rect) InflateIP(dw, dh int) {
	r.X -= dw / 2
	r.Y -= dh / 2
	r.W += dw
	r.H += dh
}

// Clamp returns a new Rect that is moved to be within bounds. If it is too large than it
// is centered within bounds.
func (r *Rect) Clamp(bounds *Rect) Rect {
	var newX, newY int
	if r.W > bounds.W {
		newX = bounds.CenterX() - r.W/2
	} else {
		newX = clampInt(r.X, bounds.X, bounds.Right()-r.W)
	}
	if r.H > bounds.H {
		newY = bounds.CenterY() - r.H/2
	} else {
		newY = clampInt(r.Y, bounds.Y, bounds.Bottom()-r.H)
	}
	return Rect{X: newX, Y: newY, W: r.W, H: r.H}
}

// ClampIP moves this Rect so that it is within bounds. If it is too large than it is
// centered within bounds.
func (r *Rect) ClampIP(bounds *Rect) {
	if r.W > bounds.W {
		r.SetCenterX(bounds.CenterX())
	} else {
		r.X = clampInt(r.X, bounds.X, bounds.Right()-r.W)
	}
	if r.H > bounds.H {
		r.SetCenterY(bounds.CenterY())
	} else {
		r.Y = clampInt(r.Y, bounds.Y, bounds.Bottom()-r.H)
	}
}

// Intersect returns a new Rect that is within both Rects. If there is no intersection
// the returned Rect will have 0 size.
func (r *Rect) Intersect(other Rect) Rect {
	newX := maxInt(r.X, other.X)
	newY := maxInt(r.Y, other.Y)
	return Rect{
		X: newX,
		Y: newY,
		W: minInt(r.Right(), other.Right()) - newX,
		H: minInt(r.Bottom(), other.Bottom()) - newY,
	}
}

// Union returns a Rect that contains both Rects
func (r *Rect) Union(other *Rect) Rect {
	newX := minInt(r.X, other.X)
	newY := minInt(r.Y, other.Y)
	return Rect{
		X: newX,
		Y: newY,
		W: maxInt(r.Right(), other.Right()) - newX,
		H: maxInt(r.Bottom(), other.Bottom()) - newY,
	}
}

// UnionIP same as Union but in place.
func (r *Rect) UnionIP(other *Rect) {
	newX := minInt(r.X, other.X)
	newY := minInt(r.Y, other.Y)
	r.W = maxInt(r.Right(), other.Right()) - newX
	r.H = maxInt(r.Bottom(), other.Bottom()) - newY
	r.X = newX
	r.Y = newY
}

// UnionAll returns a Rect that contains all Rects
func (r *Rect) UnionAll(others []*Rect) Rect {
	newX, newY := r.X, r.Y
	for _, other := range others {
		newX = minInt(newX, other.X)
		newY = minInt(newY, other.Y)
	}

	farRight, farBottom := r.Right(), r.Bottom()
	for _, other := range others {
		farRight = maxInt(farRight, other.Right())
		farBottom = minInt(farBottom, other.Bottom())
	}

	return Rect{X: newX, Y: newY, W: farRight - newX, H: farBottom - newY}
}

// Fit returns a new Rect that is moved and resized to fit within bounds while maintaining
// its original aspect ratio
func (r *Rect) Fit(bounds *Rect) Rect {
	newW := bounds.H * (r.W / r.H)
	if newW <= bounds.W {
		return Rect{X: clampInt(r.X, bounds.X, bounds.Right()-newW), Y: bounds.Y, W: newW, H: bounds.H}
	}
	newH := bounds.W * (r.H / r.W)
	return Rect{X: bounds.X, Y: clampInt(r.Y, bounds.Y, bounds.Bottom()-newH), W: bounds.W, H: newH}
}

// Normalize fips the Rect in place if its size is negative
func (r *Rect) Normalize() {
	if r.W < 0 {
		r.X += r.W
		r.W = -r.W
	}
	if r.H < 0 {
		r.Y += r.H
		r.H = -r.H
	}
}

// Contains returns true if other is completely inside this one
func (r *Rect) Contains(other *Rect) bool {
	return other.X >= r.X && other.Y >= r.W && other.Right() <= r.Right() && other.Bottom() <= r.Bottom()
}

// CollidePoint returns true if the point is within the Rect. A point along the right or
// bottom edge is not considered inside.
func (r *Rect) CollidePoint(x, y int) bool {
	return x >= r.X && x < r.Right() && y >= r.Y && y < r.Bottom()
}

// CollideRect returns true if the Rects overlap
func (r *Rect) CollideRect(other *Rect) bool {
	return r.X < other.Right() && r.Right() > other.X && r.Y < other.Bottom() && r.Bottom() > other.Y
}

// CollideList returns the index of the first Rect this one collides with, or -1 if it
// collides with none
func (r *Rect) CollideList(others []*Rect) int {
	for i, other := range others {
		if r.CollideRect(other) {
			return i
		}
	}
	return -1
}

// CollideListAll returns a list of indices of the Rects that collide with this one, or an
// empty list if none
func (r *Rect) CollideListAll(others []*Rect) []int {
	list := make([]int, 0, len(others))
	for i, other := range others {
		if r.CollideRect(other) {
			list = append(list, i)
		}
	}
	return list
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func clampInt(i, min, max int) int {
	return maxInt(minInt(i, max), min)
}
