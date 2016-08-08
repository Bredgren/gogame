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

	results = append(results, testRectInflate()...)
	results = append(results, testRectInflateIP()...)
	// results = append(results, testRectClamp()...)
	// results = append(results, testRectClampIP()...)
	// results = append(results, testRectIntersect()...)
	// results = append(results, testRectUnion()...)
	// results = append(results, testRectUnionIP()...)
	// results = append(results, testRectUnionAll()...)
	// results = append(results, testRectFit()...)
	// results = append(results, testRectNormalize()...)
	// results = append(results, testRectContains()...)
	// results = append(results, testRectCollidePoint()...)
	// results = append(results, testRectCollideRect()...)
	// results = append(results, testRectCollideList()...)
	// results = append(results, testRectCollideListAll()...)

	appendResultSection("Rect", results)
}

func testRectInflate() []*result {
	var results []*result
	res := result{TestName: "Inflate", Errors: []string{}}
	rect := gogame.Rect{X: 1, Y: 1, W: 5, H: 5}
	rect2 := rect.Inflate(2, 2)
	want := gogame.Rect{X: 0, Y: 0, W: 7, H: 7}
	if rect2 != want {
		res.Errors = append(res.Errors, fmt.Sprintf("got: %#v, expected: %#v", rect2, want))
	}
	results = append(results, &res)
	return results
}

func testRectInflateIP() []*result {
	var results []*result
	res := result{TestName: "InflateIP", Errors: []string{}}
	rect := gogame.Rect{X: 1, Y: 1, W: 5, H: 5}
	rect.InflateIP(2, 2)
	want := gogame.Rect{X: 0, Y: 0, W: 7, H: 7}
	if rect != want {
		res.Errors = append(res.Errors, fmt.Sprintf("got: %#v, expected: %#v", rect, want))
	}
	results = append(results, &res)
	return results
}
