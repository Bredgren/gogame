package main

import (
	"fmt"

	"github.com/Bredgren/gogame"
)

func testRect() {
	var results []*result

	results = append(results, testRectInflate())
	results = append(results, testRectInflateIP())
	results = append(results, testRectClamp())
	results = append(results, testRectClampIP())
	results = append(results, testRectIntersect())
	results = append(results, testRectUnion())
	// results = append(results, testRectUnionIP())
	// results = append(results, testRectUnionAll())
	// results = append(results, testRectFit())
	// results = append(results, testRectNormalize())
	// results = append(results, testRectContains())
	// results = append(results, testRectCollidePoint())
	// results = append(results, testRectCollideRect())
	// results = append(results, testRectCollideList())
	// results = append(results, testRectCollideListAll())

	appendResultSection("Rect", results)
}

func testRectInflate() *result {
	res := result{TestName: "Inflate", Errors: []string{}}

	rect := gogame.Rect{X: 1, Y: 1, W: 5, H: 5}
	rect2 := rect.Inflate(2, 2)
	want := gogame.Rect{X: 0, Y: 0, W: 7, H: 7}
	if rect2 != want {
		res.Errors = append(res.Errors, fmt.Sprintf("got: %#v, want: %#v", rect2, want))
	}

	return &res
}

func testRectInflateIP() *result {
	res := result{TestName: "InflateIP", Errors: []string{}}

	rect := gogame.Rect{X: 1, Y: 1, W: 5, H: 5}
	rect.InflateIP(2, 2)
	want := gogame.Rect{X: 0, Y: 0, W: 7, H: 7}
	if rect != want {
		res.Errors = append(res.Errors, fmt.Sprintf("got: %#v, want: %#v", rect, want))
	}

	return &res
}

func testRectClamp() *result {
	res := result{TestName: "Clamp", Errors: []string{}}

	rect := gogame.Rect{X: 1, Y: 1, W: 5, H: 5}

	rect2 := gogame.Rect{X: 0, Y: 0, W: 1, H: 1}
	want := gogame.Rect{X: 1, Y: 1, W: 1, H: 1}
	got := rect2.Clamp(&rect)
	if got != want {
		res.Errors = append(res.Errors, fmt.Sprintf("Top left: got: %#v, want: %#v", got, want))
	}

	rect2 = gogame.Rect{X: 7, Y: 6, W: 1, H: 1}
	want = gogame.Rect{X: 5, Y: 5, W: 1, H: 1}
	got = rect2.Clamp(&rect)
	if got != want {
		res.Errors = append(res.Errors, fmt.Sprintf("Bottom right: got: %#v, want: %#v", got, want))
	}

	rect2 = gogame.Rect{X: 7, Y: 6, W: 7, H: 7}
	want = gogame.Rect{X: 0, Y: 0, W: 7, H: 7}
	got = rect2.Clamp(&rect)
	if got != want {
		res.Errors = append(res.Errors, fmt.Sprintf("Too big: got: %#v, want: %#v", got, want))
	}

	return &res
}

func testRectClampIP() *result {
	res := result{TestName: "ClampIP", Errors: []string{}}

	rect := gogame.Rect{X: 1, Y: 1, W: 5, H: 5}

	got := gogame.Rect{X: 0, Y: 0, W: 1, H: 1}
	want := gogame.Rect{X: 1, Y: 1, W: 1, H: 1}
	got.ClampIP(&rect)
	if got != want {
		res.Errors = append(res.Errors, fmt.Sprintf("Top left: got: %#v, want: %#v", got, want))
	}

	got = gogame.Rect{X: 7, Y: 6, W: 1, H: 1}
	want = gogame.Rect{X: 5, Y: 5, W: 1, H: 1}
	got.ClampIP(&rect)
	if got != want {
		res.Errors = append(res.Errors, fmt.Sprintf("Bottom right: got: %#v, want: %#v", got, want))
	}

	got = gogame.Rect{X: 7, Y: 6, W: 7, H: 7}
	want = gogame.Rect{X: 0, Y: 0, W: 7, H: 7}
	got.ClampIP(&rect)
	if got != want {
		res.Errors = append(res.Errors, fmt.Sprintf("Too big: got: %#v, want: %#v", got, want))
	}

	return &res
}

func testRectIntersect() *result {
	res := result{TestName: "Intersect", Errors: []string{}}

	rect := gogame.Rect{X: 1, Y: 1, W: 5, H: 5}

	rect2 := gogame.Rect{X: 0, Y: 0, W: 2, H: 3}
	want := gogame.Rect{X: 1, Y: 1, W: 1, H: 2}
	got := rect.Intersect(&rect2)
	if got != want {
		res.Errors = append(res.Errors, fmt.Sprintf("Top left: got: %#v, want: %#v", got, want))
	}

	rect2 = gogame.Rect{X: 2, Y: 3, W: 4, H: 5}
	want = gogame.Rect{X: 2, Y: 3, W: 4, H: 3}
	got = rect.Intersect(&rect2)
	if got != want {
		res.Errors = append(res.Errors, fmt.Sprintf("Bottom right: got: %#v, want: %#v", got, want))
	}

	rect2 = gogame.Rect{X: 2, Y: 2, W: 2, H: 2}
	want = gogame.Rect{X: 2, Y: 2, W: 2, H: 2}
	got = rect.Intersect(&rect2)
	if got != want {
		res.Errors = append(res.Errors, fmt.Sprintf("Inside: got: %#v, want: %#v", got, want))
	}

	rect2 = gogame.Rect{X: 6, Y: 6, W: 2, H: 2}
	want = gogame.Rect{X: 2, Y: 2, W: 0, H: 0}
	got = rect.Intersect(&rect2)
	if got.W != 0 && got.H != 0 {
		res.Errors = append(res.Errors, fmt.Sprintf("Outside: got: %#v, want: %#v", got, want))
	}

	return &res
}

func testRectUnion() *result {
	res := result{TestName: "Union", Errors: []string{}}

	rect := gogame.Rect{X: 1, Y: 1, W: 5, H: 5}

	rect2 := gogame.Rect{X: 0, Y: 0, W: 1, H: 1}
	want := gogame.Rect{X: 0, Y: 0, W: 6, H: 6}
	got := rect2.Union(&rect)
	if got != want {
		res.Errors = append(res.Errors, fmt.Sprintf("Top left: got: %#v, want: %#v", got, want))
	}

	rect2 = gogame.Rect{X: 4, Y: 3, W: 3, H: 3}
	want = gogame.Rect{X: 1, Y: 1, W: 6, H: 5}
	got = rect2.Union(&rect)
	if got != want {
		res.Errors = append(res.Errors, fmt.Sprintf("Bottom right: got: %#v, want: %#v", got, want))
	}

	rect2 = gogame.Rect{X: 2, Y: 2, W: 2, H: 2}
	want = gogame.Rect{X: 1, Y: 1, W: 5, H: 5}
	got = rect2.Union(&rect)
	if got != want {
		res.Errors = append(res.Errors, fmt.Sprintf("Inside: got: %#v, want: %#v", got, want))
	}

	rect2 = gogame.Rect{X: 7, Y: 6, W: 7, H: 7}
	want = gogame.Rect{X: 1, Y: 1, W: 13, H: 12}
	got = rect2.Union(&rect)
	if got != want {
		res.Errors = append(res.Errors, fmt.Sprintf("Outside: got: %#v, want: %#v", got, want))
	}

	return &res
}
