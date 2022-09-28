package main

import "fmt"

func main() {
	fmt.Println("PASS BY VALUE EXAMPLE")
	passByValue()

	fmt.Println("\nPASS BY REFERENCE EXAMPLE")
	passByReference()

	fmt.Println("\nPASS BY VALUE ADDRESS EXAMPLE")
	passByValueAddress()
}

func passByValue() {
	x := 5
	fmt.Println("Before doubleInt method x is", x)
	doubleInt(x)
	fmt.Println("After doubleInt method x is", x)
}

func doubleInt(x int) {
	x *= 2
	fmt.Println("x double:", x)
}

func passByValueAddress() {
	x := 5
	fmt.Println("Before doubleIntAddress method called, x is", x)
	doubleIntAddress(&x)
	fmt.Println("After doubleIntAddress method called, x is", x)
}

func doubleIntAddress(x *int) {
	*x *= 2
	fmt.Println("memory address of x", x)
	fmt.Println("x double:", *x)
}

func passByReference() {
	mySlice := []int{1, 2, 3}
	fmt.Println("Before doubleSlice method mySlice is", mySlice)
	doubleSlice(mySlice)
	fmt.Println("After doubleSlice method mySlice is", mySlice)
}

func doubleSlice(mySlice []int) {
	for i := 0; i < len(mySlice); i++ {
		mySlice[i] *= 2
	}
	fmt.Println("mySlice double:", mySlice)
}
