package main

import "fmt"

func main() {
	name := "test"
	age := 21

	// %s -> string
	// %f -> float
	// %d -> int
	// %t -> boolean
	// %v -> any value
	// %#v -> any value
	// %T -> type of value

	fmt.Print("My name is ", name, " and I am ", age, " year(s) old.")
	fmt.Println()
	fmt.Println("My name is", name, "and I am", age, "year(s) old.")
	fmt.Printf("My name is %v and I am %v year(s) old.", name, age)
	fmt.Println()
	fmt.Println()
	fmt.Println("---------------------------------------")
	fmt.Printf("values: %v, %v, %v \n", "Eren", "Basaran", "cryptwise")
	fmt.Printf("values: %v, %v, %v %v\n", "Eren", "Basaran", "\n", "cryptwise")
	fmt.Printf("values: %#v, %#v, %#v, %#v\n", "Eren", "Basaran", "\n", "cryptwise")
	fmt.Println("---------------------------------------")
	fmt.Println()
	fmt.Printf("name variable type is %T and age variable type is %T", name, age)
	fmt.Println()
}
