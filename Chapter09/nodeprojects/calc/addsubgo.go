package main

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
)

func main() {
	exports := js.Module.Get("exports")
	exports.Set("Add", Add)
	exports.Set("Sub", Sub)
	//Make the FormatNumbers function exportable as a Javascript module
	exports.Set("FormatNumbers", FormatNumbers)
}

func Add(i, j int) int {
	return i + j
}

func Sub(i, j int) int {
	return i - j
}

type Obj struct {
	*js.Object     //For any struct type expected to be processed by Gopherjs, we need to embed the *js.Object type to it
	First      int `js:"first"`  //struct tag represents the field name in Javascript
	Second     int `js:"second"` //struct tag represents the field name in Javascript
}

func FormatNumbers(o Obj) string {
	return fmt.Sprintf("First number: %d second number: %d", o.First, o.Second)
}
