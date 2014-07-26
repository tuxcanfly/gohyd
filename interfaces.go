package main

import "fmt"
import "math"

type shape interface {
	area() float64
}

type square struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.width * s.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func printer(sh shape) {
	fmt.Println("Area of shape:", sh.area())
}

func main() {
	s := square{width: 3, height: 4}
	c := circle{radius: 5}

	printer(s)
	printer(c)
}
