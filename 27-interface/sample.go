package main

import "fmt"

type shape interface {
	area() int
}

type rectangle struct {
	a int
	b int
}

type triangle struct {
	a int
	h int
}

type square struct {
	a int
}

func (r rectangle) area() int {
	return r.a * r.b
}

func (t triangle) area() int {
	return (t.a * t.h) / 2
}

func (s square) area() int {
	return s.a * s.a
}

func writeShapesInfo(shapes ...shape) {
	for _, s := range shapes {
		fmt.Println("Area:", s.area())
	}
}

func main() {
	r := rectangle{4, 5}
	t := triangle{3, 6}
	s := square{10}
	writeShapesInfo(r, t, s)
}
