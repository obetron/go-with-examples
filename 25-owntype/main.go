package main

import "fmt"

type mile float64
type kilometer float64
type myString string

func main() {

	var m1 mile
	m1 = 3.2
	fmt.Println(m1)
	fmt.Printf("%T, %v\n", m1, m1)

	fmt.Println(3.2 + 4.4)
	f1 := 4.4
	fmt.Println(m1 + mile(f1))
	fmt.Println(float64(m1) + f1)
	fmt.Println(m1 + 2.1)

	k1 := kilometer(7.8)
	fmt.Printf("%T, %v\n", k1, k1)

	fmt.Println("\nVariable Type Example")
	fmt.Printf("Mile Variable Type: \t\t%T,\nKilometer Variable Type: \t%T\n", m1, k1)

	m2 := mile(10)
	k2 := toKilometer(m2)
	fmt.Println(k2)

}

func toKilometer(m mile) kilometer {
	return kilometer(m * 1.6)
}
