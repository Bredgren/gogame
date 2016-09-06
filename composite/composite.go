// Package composite defines composite operations.
package composite

// Op defines a type of compositing operation for controlling how to draw one
// surface to another.
type Op string

const (
	// SourceOver is the default behavior. It simply draws the new surface over the destination.
	SourceOver Op = "source-over"
	// SourceIn only draws where the surface overlap, everwhere else will be transparent.
	SourceIn Op = "source-in"
	// SourceOut only draws the new surface where it doesn't overlap with the destination..
	SourceOut Op = "source-out"
	// SourceAtop only draws the new surface where it overlaps with the destination.
	SourceAtop Op = "source-atop"
	// DestinationOver draws the new surface behind the destination.
	DestinationOver Op = "destination-over"
	// DestinationIn keeps the destination content only where it overlaps with the new surface.
	DestinationIn Op = "destination-in"
	// DestinationOut keeps the destination content only where it doesn't overlap.
	DestinationOut Op = "destination-out"
	// DestinationAtop keeps the destination content only where it overlaps and the new surface
	// is drawn behind.
	DestinationAtop Op = "destination-atop"
	// Lighter determines the color value of overlaping pixels by adding the color values.
	Lighter Op = "lighter"
	// Copy makes the destination surface a copy of the source.
	Copy Op = "copy"
	// Xor makes pixels transparent where both surfaces overlap, everwhere else is drawn normal.
	Xor Op = "xor"
	// Multiply multiplies the values of the corresponding pixels of both surfaces.
	Multiply Op = "multiply"
	// Screen inverts, multiplies, and inverts again (opposite of Multiply)
	Screen Op = "screen"
	// Overlay is a combination of multiply and screen.
	Overlay Op = "overlay"
	// Darken retains the darkest pixels of both surfaces.
	Darken Op = "darken"
	// Lighten retains the lightest pixels of both surfaces.
	Lighten Op = "lighten"
	// ColorDodge divides the destination surfaces by the inverted source surface.
	ColorDodge Op = "color-dodge"
	// ColorBurn divides the inverted destination surface by the source surface and inverts the result.
	ColorBurn Op = "color-burn"
	// HardLight is like overlay but with the source and destination swapped.
	HardLight Op = "hard-light"
	// SoftLight is a softer version of HardLight
	SoftLight Op = "soft-light"
	// Difference subtracts on surface from the other, whichever gives a positive value.
	Difference Op = "difference"
	// Exclusion is like difference but with lower contrast.
	Exclusion Op = "exclusion"
	// Hue preserves the luma and chroma of the destination while adopting the hue of the source.
	Hue Op = "hue"
	// Saturation preserves the luma and hue of the destination while adopting the chroma of the source.
	Saturation Op = "saturation"
	// Color preserves the luma of the destination while adopting the hue and chroma of the source.
	Color Op = "color"
	// Luminosity preserves the hue and chroma of the bottom layer while adopting the luma of the source.
	Luminosity Op = "luminosity"
)
