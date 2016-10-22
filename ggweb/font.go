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

// FontStyle is the style of the font.
type FontStyle string

const (
	// FontStyleNormal is the default style.
	FontStyleNormal FontStyle = "normal"
	// FontStyleItalic makes the font italics.
	FontStyleItalic FontStyle = "italic"
	// FontStyleOblique makes the font oblique.
	FontStyleOblique FontStyle = "oblique"
)

// FontVariant normal or caps.
type FontVariant string

const (
	// FontVariantNormal is the default variant.
	FontVariantNormal FontVariant = "normal"
	// FontVariantSmallCaps makes the font all capital letters.
	FontVariantSmallCaps FontVariant = "small-caps"
)

// FontWeight normal or bold.
type FontWeight string

const (
	// FontWeightNormal is the default weight.
	FontWeightNormal FontWeight = "normal"
	// FontWeightBold makes the font bold.
	FontWeightBold FontWeight = "bold"
)

// FontFamily is the overall appearence of the font. The predefined families are the generic
// ones, it is possible to define any valid css font-family though, e.g.
// FontFamily("courier new, monospace")
type FontFamily string

const (
	// FontFamilySerif has strokes at the ends of the characters.
	FontFamilySerif FontFamily = "serif"
	// FontFamilySansSerif has plain endings.
	FontFamilySansSerif FontFamily = "sans-serif"
	// FontFamilyMonospace gives equal width to all characters.
	FontFamilyMonospace FontFamily = "monospace"
	// FontFamilyCursive gives the characters a somewhat handwritten look.
	FontFamilyCursive FontFamily = "cursive"
	// FontFamilyFantasy is bit of a decorative kind of font.
	FontFamilyFantasy FontFamily = "fantasy"
)

// TextAlign aligns the text horizontally.
type TextAlign string

const (
	// TextAlignStart aligns text at the start of the line according to locale, it is the default.
	TextAlignStart TextAlign = "start"
	// TextAlignEnd aligns text at the end of the line according to locale.
	TextAlignEnd TextAlign = "end"
	// TextAlignLeft alignes text to the left.
	TextAlignLeft TextAlign = "left"
	// TextAlignRight alignes text to the right.
	TextAlignRight TextAlign = "right"
	// TextAlignCenter alignes text to the center.
	TextAlignCenter TextAlign = "center"
)

// TextBaseline aligns the text verically.
type TextBaseline string

const (
	// TextBaselineAlphabetic is the default.
	TextBaselineAlphabetic TextBaseline = "alphabetic"
	// TextBaselineTop puts the baseline at the top.
	TextBaselineTop TextBaseline = "top"
	// TextBaselineBottom puts the baseline at the bottom.
	TextBaselineBottom TextBaseline = "bottom"
	// TextBaselineHanging is the hanging baseline.
	TextBaselineHanging TextBaseline = "hanging"
	// TextBaselineMiddle puts the baseline in the middle.
	TextBaselineMiddle TextBaseline = "middle"
	// TextBaselineIdeographic is the bottom of the characters if they go beneath the alphabetic baseline.
	TextBaselineIdeographic TextBaseline = "ideographic"
)
