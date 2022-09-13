package main

import "fmt"

func main() {

	cities := [4]string{"Ankara", "Istanbul", "Sanliurfa"}
	fmt.Println("Cities:")
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

	cars := [...]string{"Volkswagen", "Mercedes", "BMW", "Audi"}
	fmt.Println("Cars:")
	for i := 0; i < len(cars); i++ {
		fmt.Println(cars[i])
	}

	var myArr1 [5]int
	var myArr2 [4]int

	fmt.Println()
	fmt.Printf("type of myArr1 is: %T\n", myArr1)
	fmt.Printf("type of myArr2 is: %T\n", myArr2)

}
