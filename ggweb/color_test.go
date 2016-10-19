package ggweb

import (
	"fmt"
	"image/color"
	"strconv"
	"testing"
)

/*
In previous design for the gogame package it was discovered that converting our custom
Color type into a CSS formatted string was a bottleneck. So this group of tests is for
discovering which method will be best. Though with the new design this may not be as big
of an issue.
*/
var c color.Color = color.RGBA{5, 10, 15, 20}
var s string

func BenchmarkColorSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r, g, b, a := c.RGBA()
		s = fmt.Sprintf("rgba(%d, %d, %d, %f)", mapColor(r), mapColor(g), mapColor(b), float64(mapColor(a))/255)
	}
}

func BenchmarkColorConvAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r, g, b, a := c.RGBA()
		s = "rgba(" + strconv.Itoa(mapColor(r)) + "," + strconv.Itoa(mapColor(g)) + "," +
			strconv.Itoa(mapColor(b)) + "," +
			strconv.FormatFloat(float64(mapColor(a))/255, 'f', -1, 64) + ")"
	}
}
