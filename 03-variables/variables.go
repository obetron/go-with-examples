package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	otherA := "other initial"
	fmt.Println(otherA)

	var b, c int = 1, 2
	fmt.Println(b, c)

	otherB, otherC := 3, 4
	fmt.Println("other b and c:", otherB, otherC)

	var d = true
	fmt.Println(d)

	otherD := true
	fmt.Println("other boolean:", otherD)

	var e bool = false
	fmt.Println("another boolean: ", e)
}
