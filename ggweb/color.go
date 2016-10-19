package ggweb

import (
	"image/color"
	"strconv"
)

// ColorToCSS converts a color to a CSS formatted color string, e.g. "rgba(0, 0, 0, 0)"
func ColorToCSS(c color.Color) string {
	r, g, b, a := c.RGBA()
	return "rgba(" + strconv.Itoa(mapColor(r)) + "," + strconv.Itoa(mapColor(g)) + "," +
		strconv.Itoa(mapColor(b)) + "," +
		strconv.FormatFloat(float64(mapColor(a))/255, 'f', -1, 64) + ")"
}

func mapColor(i uint32) int {
	return int(float64(i) / float64(0xffff) * 0xff)
}
