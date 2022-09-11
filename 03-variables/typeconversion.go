package main

import "fmt"

func main() {
	x := 10
	y := 10.0

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)

	// fmt.Println(x + y) //invalid operation for int and float64 type

	//Type Convertion for y and make addition
	fmt.Println(x + int(y))
	fmt.Println(float64(x) + y) //convertion for just this line x type still int
	fmt.Printf("x type: %T\n", x)

	z := int(y)
	fmt.Printf("%v %T\n", z, z)

	var k int16 = 129
	var l int8
	l = int8(k)
	fmt.Println(l)

	m := 106
	// n := string(m)
	n := fmt.Sprint(m)
	fmt.Printf("%v %T", m, m)
	fmt.Println()
	fmt.Printf("%v %T", n, n)
}
