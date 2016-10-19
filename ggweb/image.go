package ggweb

// LoadImage loads the given image and returns it as a Surface. Important: This is a
// blocking call and so cannot be called in a js callback or main thread without wrapping
// it in a go routine.
func LoadImage(filename string) *Surface {
	// img := js.Global.Get("Image").New()
	// img.Set("src", filename)
	// c := make(chan Surface, 0)
	// img.Call("addEventListener", "load", func() {
	// 	s := NewSurface(img.Get("width").Int(), img.Get("height").Int())
	// 	s.Canvas().Call("getContext", "2d").Call("drawImage", img, 0, 0)
	// 	c <- s
	// 	close(c)
	// })
	// return <-c
	return nil
}
