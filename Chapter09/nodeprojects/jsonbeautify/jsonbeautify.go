package main

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
)

func main() {
	type MyType struct {
		Name     string
		Projects []string
	}
	value := MyType{Name: "mina", Projects: []string{"GopherJS", "ReactJS"}}
	prettyjson := js.Global.Call("require", "prettyjson")
	fmt.Println(prettyjson.Call("render", value))
}
