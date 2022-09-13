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

}
