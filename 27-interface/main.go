package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	circumference() float64
}

type rectangle struct {
	a, b float64
}

type circle struct {
	r float64
}

func (r rectangle) area() float64 {
	return r.a * r.b
}

func (r rectangle) circumference() float64 {
	return 2 * (r.a + r.b)
}

func (r rectangle) display() {
	fmt.Println("Kenar 1:", r.a, "\nKenar 2:", r.b)
}

func (c circle) display() {
	fmt.Println("Yari cap:", c.r)
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c circle) circumference() float64 {
	return 2 * math.Pi * c.r
}

func interfaceFunc(i shape) {
	fmt.Println(i)
	fmt.Println(i.area())
	fmt.Println(i.circumference())

	fmt.Printf("%T\n", i)
}

func main() {

	r1 := rectangle{5, 4}
	fmt.Println("\nRECTANGLE SHAPE EXAMPLE")
	interfaceFunc(r1)

	c1 := circle{12}
	fmt.Println("\nCIRCLE SHAPE EXAPLE")
	interfaceFunc(c1)

}
