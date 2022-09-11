package main

import "fmt"

func main() {
	name := "test"
	age := 21

	fmt.Print("My name is ", name, " and I am ", age, " year(s) old.")
	fmt.Println()
	fmt.Println("My name is", name, "and I am", age, "year(s) old.")
	fmt.Printf("My name is %v and I am %v year(s) old.", name, age)
	fmt.Println()

	fmt.Printf("name variable type is %T and age variable type is %T", name, age)
	fmt.Println()
}
