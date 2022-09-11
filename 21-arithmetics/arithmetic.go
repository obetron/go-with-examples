package main

import "fmt"

func main() {
	x, y := 15, 10

	fmt.Printf("%T, %v\n", x, x)             // int, 10
	fmt.Printf("%T, %v\n", y, y)             // int, 10
	fmt.Printf("%T, %v\n", (x + y), (x + y)) // int, 25
	fmt.Printf("%T, %v\n", (x - y), (x - y)) // int, 5
	fmt.Printf("%T, %v\n", (x * y), (x * y)) // int, 150
	fmt.Printf("%T, %v\n", (x / y), (x / y)) // int, 1

	z := 5.0 / 2 //float64, 2.5
	fmt.Printf("%T, %v\n", z, z)

	k := 10
	fmt.Println(k)     //10
	fmt.Println(k + 1) //11

	// fmt.Println(k++) //k++ also an expression so in one line can not define 2 expressions.
	k++
	fmt.Println(k) //11

	fmt.Println(k - 1) //10

	k--
	fmt.Println(k) //10

	// x = x + 10 // assignment statement
	// x -= 10 //assigment operation
}
