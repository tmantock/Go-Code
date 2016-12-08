package main

import (
	"fmt"
	"math"
)

//Square type impelements a new type of Square
type Square struct {
	side float64
}

func (z Square) area() float64 {
	return z.side * z.side
}

//Shape type implements a new interface
type Shape interface {
	area() float64
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

//Circle type implements a new type of Circle
type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	s := Square{10}
	c := Circle{5}
	info(s)
	info(c)
}
