package main

import "fmt"

// create shape interface
type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
