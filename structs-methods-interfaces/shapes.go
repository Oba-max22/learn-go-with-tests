package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

func Perimeter(rectangle Rectangle) (p float64) {
	p = 2 * (rectangle.Width + rectangle.Height)
	return
}

func main() {
	fmt.Println(Perimeter(Rectangle{20.0, 10.0}))
	fmt.Println(Rectangle{20.0, 10.0}.Area())
	fmt.Println(Circle{10.0}.Area())
}