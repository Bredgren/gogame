// Package ggweb is wrapper around gopherjs that makes it more convenient to work with for making games.
//
// A few lines to help get started, assuming there is a canvas element in the page with
// ID "main-canvas".
//	func main() {
//		ggweb.Init(onInit)
//	}
//
//	func onInit() {
//		width, height := 100, 100
//		display := ggweb.NewSurfaceFromID("main-canvas")
//		display.SetSize(width, height)
//		display.StyleColor(ggweb.Fill, color.Black)
//		display.DrawRect(ggweb.Fill, display.Rect())
//	}
//
// Check out the examples directory for more detailed examples.
//
// TODO
//  - Touch input
//  - Network
package ggweb
