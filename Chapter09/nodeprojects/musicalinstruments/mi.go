package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Set("musicalInstruments", map[string]interface{}{
		"New": New,
	})

}

type MI struct {
	MIType string
	Price  float64
	Color  string
	Age    int
}

func (mi *MI) SetMIType(s string) {
	mi.MIType = s
}

func (mi *MI) GetMIType() string {
	return mi.MIType
}

func (mi *MI) SetPrice(f float64) {
	mi.Price = f
}

func (mi *MI) GetPrice() float64 {
	return mi.Price
}

func (mi *MI) SetColor(c string) {
	mi.Color = c
}

func (mi *MI) GetColor() string {
	return mi.Color
}

func (mi *MI) SetAge(a int) {
	mi.Age = a
}

func (mi *MI) GetAge() int {
	return mi.Age
}

func New() *js.Object {
	return js.MakeWrapper(&MI{})
}
