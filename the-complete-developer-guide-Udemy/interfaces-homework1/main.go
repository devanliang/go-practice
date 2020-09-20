package main

import "fmt"

type shape interface {
	getAres() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {
	tr := triangle{2, 3}
	sq := square{5}

	printArea(tr)
	printArea(sq)
}

func (t triangle) getAres() float64 {
	area := 0.5 * t.base * t.height
	return area
}

func (s square) getAres() float64 {
	area := s.sideLength * s.sideLength
	return area
}

func printArea(s shape) {
	fmt.Println(s.getAres())
}
