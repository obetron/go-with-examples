package main

import "fmt"

func main() {

	//SHADOWING
	x := 5
	if true {
		x := 10
		fmt.Println(x)
	}
	fmt.Println(x)

	//ASSIGNMENT
	y := 5
	if true {
		y = 10
		y++
		fmt.Println(y)
	}
	fmt.Println(y)
}
