package ggweb

// DrawType specifies the method used for drawing.
type DrawType string

const (
	// Fill draws within the shape's boundaries.
	Fill DrawType = "fill"
	// Stroke draws only the shape's boundaries.
	Stroke DrawType = "stroke"
)

// RepeatType describes how to repeat.
type RepeatType string

const (
	// RepeatXY repeats in both horizontal and vertical directions.
	RepeatXY RepeatType = "repeat"
	// RepeatX repeats in the horizontal direction.
	RepeatX RepeatType = "repeat-x"
	// RepeatY repeats in the vertical direction.
	RepeatY RepeatType = "repeat-y"
	// NoRepeat doesn't repeat.
	NoRepeat RepeatType = "no-repeat"
)

// LineCap is a style of line cap.
type LineCap string

const (
	// LineCapButt draws a line with no ends.
	LineCapButt LineCap = "butt"
	// LineCapRound draws a line with rounded ends with radius equal to half its width.
	LineCapRound LineCap = "round"
	// LineCapSquare draws a line with the ends capped with a box that extends by an amount
	// equal to half the lines width.
	LineCapSquare LineCap = "square"
)

// LineJoin is the style for the point where two lines are connected.
type LineJoin string

const (
	// LineJoinRound joins lines with rounded corners.
	LineJoinRound LineJoin = "round"
	// LineJoinBevel joins lines by filling in the triangular gap between them.
	LineJoinBevel LineJoin = "bevel"
	// LineJoinMiter joins lines by extending the edges until they meet.
	LineJoinMiter LineJoin = "miter"
)

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

// Cursor is a style for the cursor. It can be any valid css value for the cursor property,
// most of which are predefined in this package.
type Cursor string

const (
	CursorDefault      Cursor = "default"
	CursorNone         Cursor = "none"
	CursorPointer      Cursor = "pointer"
	CursorText         Cursor = "text"
	CursorVerticalText Cursor = "vertical-text"
	CursorProgress     Cursor = "progress"
	CusorWait          Cursor = "wait"
	CursorAlias        Cursor = "alias"
	CursorAllScroll    Cursor = "all-scroll"
	CursorMove         Cursor = "move"
	CursorCell         Cursor = "cell"
	CursorCopy         Cursor = "copy"
	CursorCrosshair    Cursor = "crosshair"
	CursorNSResize     Cursor = "ns-resize"
	CursorEWResize     Cursor = "ew-resize"
	CursorNESWResize   Cursor = "nesw-resize"
	CursorNWSEResize   Cursor = "nwse-resize"
	CursorRowReszie    Cursor = "row-resize"
	CursorColResize    Cursor = "col-resize"
	CursorHelp         Cursor = "help"
	CursorNoDrop       Cursor = "no-drop"
	CursorNotAllowed   Cursor = "not-allowed"
)

// CompositeOp defines a type of compositing operation for controlling how to draw one
// surface to another.
type CompositeOp string

const (
	// SourceOver is the default behavior. It simply draws the new surface over the destination.
	SourceOver CompositeOp = "source-over"
	// SourceIn only draws where the surface overlap, everwhere else will be transparent.
	SourceIn CompositeOp = "source-in"
	// SourceOut only draws the new surface where it doesn't overlap with the destination..
	SourceOut CompositeOp = "source-out"
	// SourceAtop only draws the new surface where it overlaps with the destination.
	SourceAtop CompositeOp = "source-atop"
	// DestinationOver draws the new surface behind the destination.
	DestinationOver CompositeOp = "destination-over"
	// DestinationIn keeps the destination content only where it overlaps with the new surface.
	DestinationIn CompositeOp = "destination-in"
	// DestinationOut keeps the destination content only where it doesn't overlap.
	DestinationOut CompositeOp = "destination-out"
	// DestinationAtop keeps the destination content only where it overlaps and the new surface
	// is drawn behind.
	DestinationAtop CompositeOp = "destination-atop"
	// Lighter determines the color value of overlaping pixels by adding the color values.
	Lighter CompositeOp = "lighter"
	// Copy makes the destination surface a copy of the source.
	Copy CompositeOp = "copy"
	// Xor makes pixels transparent where both surfaces overlap, everwhere else is drawn normal.
	Xor CompositeOp = "xor"
	// Multiply multiplies the values of the corresponding pixels of both surfaces.
	Multiply CompositeOp = "multiply"
	// Screen inverts, multiplies, and inverts again (opposite of Multiply)
	Screen CompositeOp = "screen"
	// Overlay is a combination of multiply and screen.
	Overlay CompositeOp = "overlay"
	// Darken retains the darkest pixels of both surfaces.
	Darken CompositeOp = "darken"
	// Lighten retains the lightest pixels of both surfaces.
	Lighten CompositeOp = "lighten"
	// ColorDodge divides the destination surfaces by the inverted source surface.
	ColorDodge CompositeOp = "color-dodge"
	// ColorBurn divides the inverted destination surface by the source surface and inverts the result.
	ColorBurn CompositeOp = "color-burn"
	// HardLight is like overlay but with the source and destination swapped.
	HardLight CompositeOp = "hard-light"
	// SoftLight is a softer version of HardLight
	SoftLight CompositeOp = "soft-light"
	// Difference subtracts on surface from the other, whichever gives a positive value.
	Difference CompositeOp = "difference"
	// Exclusion is like difference but with lower contrast.
	Exclusion CompositeOp = "exclusion"
	// Hue preserves the luma and chroma of the destination while adopting the hue of the source.
	Hue CompositeOp = "hue"
	// Saturation preserves the luma and hue of the destination while adopting the chroma of the source.
	Saturation CompositeOp = "saturation"
	// Color preserves the luma of the destination while adopting the hue and chroma of the source.
	Color CompositeOp = "color"
	// Luminosity preserves the hue and chroma of the bottom layer while adopting the luma of the source.
	Luminosity CompositeOp = "luminosity"
)
