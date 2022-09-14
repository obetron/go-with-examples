package main

import "fmt"

func main() {
	name := "George"
	fmt.Println("name Value:", name)
	fmt.Println("name Pointer", &name)

	x := 35
	fmt.Println("x Value:", x)             // --> value
	fmt.Println("x Pointer:", &x)          // --> address of x in memory (pointer x)
	fmt.Println("x Pointer Value:", *(&x)) // --> value of address in memory (dereferencing))

	y1 := 5
	y2 := &y1
	fmt.Println("y1:", y1, "y2:", y2)
	fmt.Println("y1:", y1, "&y2:", *y2)

	*y2 = 3
	fmt.Println("y1:", y1, "*y2", *y2)
}
