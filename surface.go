package gogame

// Surface represents an image or drawable sureface
type Surface interface {
	Blit(source Surface, x, y int)
	Fill(Color)
	Width() int
	Height() int
	//Copy() Surface
	//Scroll(dx, dy int)
	//GetAt(x, y int) Color
	//SetAt(x, y int, Color)
	//SetClip(Rect)
	//GetClip() Rect
	//GetSubsurface(Rect) Surface
	//GetParent() Surface
	//GetRect() Rect
}

var _ Surface = &surface{}

type surface struct {
	width  int
	height int
}

// NewSurface creates a new Surface
func NewSurface(width, height int) Surface {
	return &surface{}
}

// Blit draws the given surface to this one at the given position
func (s *surface) Blit(source Surface, x, y int) {
	return
}

// Fill fills the whole surface with one color
func (s *surface) Fill(color Color) {

}

// Width returns the width of the surface in pixels
func (s *surface) Width() int {
	return s.width
}

// Height returns the height of the surface in pixels
func (s *surface) Height() int {
	return s.height
}

//type subsurface struct
