package ggweb

import "fmt"

// Font describes the style of text.
type Font struct {
	// Size is the Font's height in pixels
	Size    int
	Family  FontFamily
	Style   FontStyle
	Variant FontVariant
	Weight  FontWeight
}

// String implements the Stringer interface. The format of the returned string is the same
// as one would use to set the font attribute in CSS.
func (f *Font) String() string {
	s := ""
	if f == nil {
		return s
	}

	if f.Style != "" {
		s = string(f.Style)
	}

	const sep = " "
	if f.Variant != "" {
		if s != "" {
			s += sep
		}
		s += string(f.Variant)
	}
	if f.Weight != "" {
		if s != "" {
			s += sep
		}
		s += string(f.Weight)
	}

	if s != "" {
		s += sep
	}
	s += fmt.Sprintf("%dpx", f.Size)

	family := f.Family
	if f.Family == "" {
		family = FontFamilySerif
	}
	if s != "" {
		s += sep
	}
	s += string(family)

	return s
}
