package main

import (
	"fmt"

	"github.com/Bredgren/gogame/geo"
)

func testRect() {
	var results []*result

	results = append(results, testRectInflate())
	results = append(results, testRectInflateIP())
	results = append(results, testRectClamp())
	results = append(results, testRectClampIP())
	results = append(results, testRectIntersect())
	results = append(results, testRectUnion())
	results = append(results, testRectUnionIP())
	results = append(results, testRectUnionAll())
	results = append(results, testRectFit())
	results = append(results, testRectNormalize())
	results = append(results, testRectContains())
	results = append(results, testRectCollidePoint())
	results = append(results, testRectCollideRect())

	appendResultSection("Rect", results)
}

func testRectInflate() *result {
	res := result{TestName: "Inflate", Errors: []string{}}

	rect := geo.Rect{X: 1, Y: 1, W: 5, H: 5}
	rect2 := rect.Inflate(2, 2)
	want := geo.Rect{X: 0, Y: 0, W: 7, H: 7}
	if rect2 != want {
		res.Errors = append(res.Errors, fmt.Sprintf("got: %#v, want: %#v", rect2, want))
	}

	return &res
}

func testRectInflateIP() *result {
	res := result{TestName: "InflateIP", Errors: []string{}}

	rect := geo.Rect{X: 1, Y: 1, W: 5, H: 5}
	rect.InflateIP(2, 2)
	want := geo.Rect{X: 0, Y: 0, W: 7, H: 7}
	if rect != want {
		res.Errors = append(res.Errors, fmt.Sprintf("got: %#v, want: %#v", rect, want))
	}

	return &res
}

func testRectClamp() *result {
	res := result{TestName: "Clamp", Errors: []string{}}

	rect := geo.Rect{X: 1, Y: 1, W: 5, H: 5}

	rect2 := geo.Rect{X: 0, Y: 0, W: 1, H: 1}
	want := geo.Rect{X: 1, Y: 1, W: 1, H: 1}
	got := rect2.Clamp(&rect)
	if got != want {
		res.Errors = append(res.Errors, fmt.Sprintf("Top left: got: %#v, want: %#v", got, want))
	}

	rect2 = geo.Rect{X: 7, Y: 6, W: 1, H: 1}
	want = geo.Rect{X: 5, Y: 5, W: 1, H: 1}
	got = rect2.Clamp(&rect)
	if got != want {
		res.Errors = append(res.Errors, fmt.Sprintf("Bottom right: got: %#v, want: %#v", got, want))
	}

	rect2 = geo.Rect{X: 7, Y: 6, W: 7, H: 7}
	want = geo.Rect{X: 0, Y: 0, W: 7, H: 7}
	got = rect2.Clamp(&rect)
	if got != want {
		res.Errors = append(res.Errors, fmt.Sprintf("Too big: got: %#v, want: %#v", got, want))
	}

	return &res
}

func testRectClampIP() *result {
	res := result{TestName: "ClampIP", Errors: []string{}}

	rect := geo.Rect{X: 1, Y: 1, W: 5, H: 5}

	got := geo.Rect{X: 0, Y: 0, W: 1, H: 1}
	want := geo.Rect{X: 1, Y: 1, W: 1, H: 1}
	got.ClampIP(&rect)
	if got != want {
		res.Errors = append(res.Errors, fmt.Sprintf("Top left: got: %#v, want: %#v", got, want))
	}

	got = geo.Rect{X: 7, Y: 6, W: 1, H: 1}
	want = geo.Rect{X: 5, Y: 5, W: 1, H: 1}
	got.ClampIP(&rect)
	if got != want {
		res.Errors = append(res.Errors, fmt.Sprintf("Bottom right: got: %#v, want: %#v", got, want))
	}

	got = geo.Rect{X: 7, Y: 6, W: 7, H: 7}
	want = geo.Rect{X: 0, Y: 0, W: 7, H: 7}
	got.ClampIP(&rect)
	if got != want {
		res.Errors = append(res.Errors, fmt.Sprintf("Too big: got: %#v, want: %#v", got, want))
	}

	return &res
}

func testRectIntersect() *result {
	res := result{TestName: "Intersect", Errors: []string{}}

	rect := geo.Rect{X: 1, Y: 1, W: 5, H: 5}

	rect2 := geo.Rect{X: 0, Y: 0, W: 2, H: 3}
	want := geo.Rect{X: 1, Y: 1, W: 1, H: 2}
	got := rect.Intersect(&rect2)
	if got != want {
		res.Errors = append(res.Errors, fmt.Sprintf("Top left: got: %#v, want: %#v", got, want))
	}

	rect2 = geo.Rect{X: 2, Y: 3, W: 4, H: 5}
	want = geo.Rect{X: 2, Y: 3, W: 4, H: 3}
	got = rect.Intersect(&rect2)
	if got != want {
		res.Errors = append(res.Errors, fmt.Sprintf("Bottom right: got: %#v, want: %#v", got, want))
	}

	rect2 = geo.Rect{X: 2, Y: 2, W: 2, H: 2}
	want = geo.Rect{X: 2, Y: 2, W: 2, H: 2}
	got = rect.Intersect(&rect2)
	if got != want {
		res.Errors = append(res.Errors, fmt.Sprintf("Inside: got: %#v, want: %#v", got, want))
	}

	rect2 = geo.Rect{X: 6, Y: 6, W: 2, H: 2}
	want = geo.Rect{X: 2, Y: 2, W: 0, H: 0}
	got = rect.Intersect(&rect2)
	if got.W != 0 && got.H != 0 {
		res.Errors = append(res.Errors, fmt.Sprintf("Outside: got: %#v, want: %#v", got, want))
	}

	return &res
}

func testRectUnion() *result {
	res := result{TestName: "Union", Errors: []string{}}

	rect := geo.Rect{X: 1, Y: 1, W: 5, H: 5}

	rect2 := geo.Rect{X: 0, Y: 0, W: 1, H: 1}
	want := geo.Rect{X: 0, Y: 0, W: 6, H: 6}
	got := rect2.Union(&rect)
	if got != want {
		res.Errors = append(res.Errors, fmt.Sprintf("Top left: got: %#v, want: %#v", got, want))
	}

	rect2 = geo.Rect{X: 4, Y: 3, W: 3, H: 3}
	want = geo.Rect{X: 1, Y: 1, W: 6, H: 5}
	got = rect2.Union(&rect)
	if got != want {
		res.Errors = append(res.Errors, fmt.Sprintf("Bottom right: got: %#v, want: %#v", got, want))
	}

	rect2 = geo.Rect{X: 2, Y: 2, W: 2, H: 2}
	want = geo.Rect{X: 1, Y: 1, W: 5, H: 5}
	got = rect2.Union(&rect)
	if got != want {
		res.Errors = append(res.Errors, fmt.Sprintf("Inside: got: %#v, want: %#v", got, want))
	}

	rect2 = geo.Rect{X: 7, Y: 6, W: 7, H: 7}
	want = geo.Rect{X: 1, Y: 1, W: 13, H: 12}
	got = rect2.Union(&rect)
	if got != want {
		res.Errors = append(res.Errors, fmt.Sprintf("Outside: got: %#v, want: %#v", got, want))
	}

	return &res
}

func testRectUnionIP() *result {
	res := result{TestName: "UnionIP", Errors: []string{}}

	rect := geo.Rect{X: 1, Y: 1, W: 5, H: 5}

	got := geo.Rect{X: 0, Y: 0, W: 1, H: 1}
	want := geo.Rect{X: 0, Y: 0, W: 6, H: 6}
	got.UnionIP(&rect)
	if got != want {
		res.Errors = append(res.Errors, fmt.Sprintf("Top left: got: %#v, want: %#v", got, want))
	}

	got = geo.Rect{X: 4, Y: 3, W: 3, H: 3}
	want = geo.Rect{X: 1, Y: 1, W: 6, H: 5}
	got.UnionIP(&rect)
	if got != want {
		res.Errors = append(res.Errors, fmt.Sprintf("Bottom right: got: %#v, want: %#v", got, want))
	}

	got = geo.Rect{X: 2, Y: 2, W: 2, H: 2}
	want = geo.Rect{X: 1, Y: 1, W: 5, H: 5}
	got.UnionIP(&rect)
	if got != want {
		res.Errors = append(res.Errors, fmt.Sprintf("Inside: got: %#v, want: %#v", got, want))
	}

	got = geo.Rect{X: 7, Y: 6, W: 7, H: 7}
	want = geo.Rect{X: 1, Y: 1, W: 13, H: 12}
	got.UnionIP(&rect)
	if got != want {
		res.Errors = append(res.Errors, fmt.Sprintf("Outside: got: %#v, want: %#v", got, want))
	}

	return &res
}

func testRectUnionAll() *result {
	res := result{TestName: "UnionAll", Errors: []string{}}

	rect := geo.Rect{X: 1, Y: 1, W: 5, H: 5}

	rects := []*geo.Rect{
		&geo.Rect{X: 0, Y: 2, W: 3, H: 6},
		&geo.Rect{X: 4, Y: -1, W: 4, H: 4},
	}

	want := geo.Rect{X: 0, Y: -1, W: 8, H: 9}
	got := rect.UnionAll(rects)
	if got != want {
		res.Errors = append(res.Errors, fmt.Sprintf("got: %#v, want: %#v", got, want))
	}

	return &res
}

func testRectFit() *result {
	res := result{TestName: "Fit", Errors: []string{}}

	rect1 := geo.Rect{X: 1, Y: 1, W: 6, H: 3}
	rect2 := geo.Rect{X: 2, Y: 2, W: 3, H: 6}
	rect3 := geo.Rect{X: 3, Y: 3, W: 8, H: 2}
	rect4 := geo.Rect{X: 4, Y: 4, W: 2, H: 8}

	cases := []struct {
		r1, r2, want *geo.Rect
	}{
		{&rect1, &rect2, &geo.Rect{X: 2, Y: 2, W: 3, H: 1.5}},
		{&rect1, &rect3, &geo.Rect{X: 3, Y: 3, W: 4, H: 2}},
		{&rect1, &rect4, &geo.Rect{X: 4, Y: 4, W: 2, H: 1}},

		{&rect2, &rect1, &geo.Rect{X: 2, Y: 1, W: 1.5, H: 3}},
		{&rect2, &rect3, &geo.Rect{X: 3, Y: 3, W: 1, H: 2}},
		{&rect2, &rect4, &geo.Rect{X: 4, Y: 4, W: 2, H: 4}},

		{&rect3, &rect1, &geo.Rect{X: 1, Y: 4 - 6.0/4.0, W: 6, H: 6.0 / 4.0}},
		{&rect3, &rect2, &geo.Rect{X: 2, Y: 3, W: 3, H: 3.0 / 4.0}},
		{&rect3, &rect4, &geo.Rect{X: 4, Y: 4, W: 2, H: 2.0 / 4.0}},

		{&rect4, &rect1, &geo.Rect{X: 4, Y: 1, W: 3.0 / 4.0, H: 3}},
		{&rect4, &rect2, &geo.Rect{X: 5 - 6.0/4.0, Y: 2, W: 6.0 / 4.0, H: 6}},
		{&rect4, &rect3, &geo.Rect{X: 4, Y: 3, W: 2.0 / 4.0, H: 2.0}},
	}

	for i, c := range cases {
		got := c.r1.Fit(c.r2)
		if got != *c.want {
			res.Errors = append(res.Errors, fmt.Sprintf("%d: got: %#v, want: %#v", i, got, c.want))
		}
	}

	return &res
}

func testRectNormalize() *result {
	res := result{TestName: "Normalize", Errors: []string{}}

	got := geo.Rect{X: 1, Y: 1, W: -4, H: -2}
	want := geo.Rect{X: -3, Y: -1, W: 4, H: 2}
	got.Normalize()
	if got != want {
		res.Errors = append(res.Errors, fmt.Sprintf("got: %#v, want: %#v", got, want))
	}

	return &res
}

func testRectContains() *result {
	res := result{TestName: "Contains", Errors: []string{}}

	rect1 := geo.Rect{X: 1, Y: 1, W: 5, H: 5}
	rect2 := geo.Rect{X: 2, Y: 2, W: 5, H: 2}
	rect3 := geo.Rect{X: 2, Y: 2, W: 4, H: 2}

	if rect1.Contains(&rect2) {
		res.Errors = append(res.Errors, fmt.Sprintf("got: rect1 contains rect2, want: rect1 DOESN'T contain rect2"))
	}

	if rect1.Contains(&rect3) {
		res.Errors = append(res.Errors, fmt.Sprintf("got: rect1 doesn't contain rect3, want: rect1 DOES contain rect3"))
	}

	return &res
}

func testRectCollidePoint() *result {
	res := result{TestName: "CollidePoint", Errors: []string{}}

	rect := geo.Rect{X: 1, Y: 1, W: 5, H: 5}

	cases := []struct {
		x, y float64
		want bool
	}{
		{0, 0, false},
		{1, 1, true},
		{4, 4, true},
		{5, 5, true},
		{6, 6, false},
	}

	for i, c := range cases {
		got := rect.CollidePoint(c.x, c.y)
		if got != c.want {
			res.Errors = append(res.Errors, fmt.Sprintf("%d: got: %#v, want: %#v", i, got, c.want))
		}
	}

	return &res
}

func testRectCollideRect() *result {
	res := result{TestName: "CollideRect", Errors: []string{}}

	rect := geo.Rect{X: 1, Y: 1, W: 5, H: 5}

	cases := []struct {
		r    geo.Rect
		want bool
	}{
		{geo.Rect{X: 0, Y: 0, W: 7, H: 1}, false},
		{geo.Rect{X: 0, Y: 0, W: 1, H: 7}, false},
		{geo.Rect{X: 6, Y: 0, W: 2, H: 7}, false},
		{geo.Rect{X: 0, Y: 6, W: 7, H: 2}, false},
		{geo.Rect{X: 0, Y: 0, W: 2, H: 2}, true},
		{geo.Rect{X: 5, Y: 5, W: 2, H: 2}, true},
	}

	for i, c := range cases {
		got := rect.CollideRect(&c.r)
		if got != c.want {
			res.Errors = append(res.Errors, fmt.Sprintf("%d: got: %#v, want: %#v", i, got, c.want))
		}
	}

	return &res
}
