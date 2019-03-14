package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	document := js.Global.Get("document")
	document.Call("write", "Hello world!!")
}
