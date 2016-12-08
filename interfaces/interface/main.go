package main

import "fmt"

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

func main() {
	s := Square{10}
	info(s)
}
