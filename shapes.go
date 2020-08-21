package learn_go_with_tests

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base float64
	Height float64
}
func (r Rectangle) Perimeter() (perimeter float64) {
	return 2 * (r.Length + r.Width)
}

func (r Rectangle) Area() (area float64) {
	return r.Length * r.Width
}

func (c Circle) Perimeter() (perimeter float64) {
	return math.Pi * 2 * c.Radius
}

func (c Circle) Area() (area float64) {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() (area float64) {
	return t.Base * t.Height / 2.0
}
