// To run the tests run "gopherjs serve $GOPATH/github.com/Bredgren/gogame/test" and navigate
// to http://localhost:8080/github.com/Bredgren/gogame/test/ in your browser.
package main

import (
	"fmt"

	"github.com/Bredgren/gogame"
	"github.com/gopherjs/jquery"
)

var jq = jquery.NewJQuery

var resultList jquery.JQuery

type result struct {
	TestName string
	Errors   []string
}

func main() {
	resultList = jq("#test-results")
	testRect()

	// ready := gogame.Ready()
	// go func() {
	// 	<-ready
	// 	start()
	// }()
}

func appendResultSection(sectionName string, results []*result) {
	section := jq("<li>").AddClass("result-section")
	section.Append(jq("<h1>").SetText(sectionName))
	sectionRes := jq("<ul>").AddClass("result-section-list")
	section.Append(sectionRes)
	for _, result := range results {
		res := jq("<li>").AddClass("result")
		res.Append(jq("<h2>").SetText(result.TestName))
		if len(result.Errors) == 0 {
			res.AddClass("result-pass")
		} else {
			res.AddClass("result-fail")
			errors := jq("<ul>").AddClass("result-error-list")
			for _, e := range result.Errors {
				errors.Append(jq("<li>").AddClass("result-error").SetText(e))
			}
			res.Append(errors)
		}
		sectionRes.Append(res)
	}
	resultList.Append(section)
}

func testRect() {
	var results []*result

	results = append(results, testRectInflate())
	results = append(results, testRectInflateIP())
	results = append(results, testRectClamp())
	results = append(results, testRectClampIP())
	// results = append(results, testRectIntersect())
	// results = append(results, testRectUnion())
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
