package main

import "fmt"

/*
*
defining variables, order is so important. If bool variable declared first, and int16 declared, it is a 1 byte val,
but keep 2 bytes in memory and 1 byte padding.

but first bool val declared and after int32 variable declared then 1 byte keeped for bool and 3 bytes padding and then
4 bytes keeped for int32 variable.
*/
func main() {
	name := "test"
	age := 21
	grade := 30.4
	active := true // bool variable 1 byte but they keep 2 bytes in memory,

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
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Printf("values: %v, %v, %v \n", "Eren", "Basaran", "cryptwise")
	fmt.Printf("values: %v, %v, %v %v\n", "Eren", "Basaran", "\n", "cryptwise")
	fmt.Printf("values: %#v, %#v, %#v, %#v\n", "Eren", "Basaran", "\n", "cryptwise")
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println()
	fmt.Printf("name variable type is %T and age variable type is %T", name, age)
	fmt.Println()
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Printf("String types: \t %T [%v]\n", name, name)
	fmt.Printf("int types: \t %T [%v]\n", age, age)
	fmt.Printf("float types: \t %T [%v]\n", grade, grade)
	fmt.Printf("bool types: \t %T [%v]\n", active, active)
}
