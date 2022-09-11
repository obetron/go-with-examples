package main

import "fmt"

var packageVar = "Package Scope Variable"

// packageVar1 := "Package Scope Variable" //non-decleration statement outside function body is not supported

func main() {

	if true {
		var ifBlockVar = "If Block Scope Variable"
		fmt.Println(ifBlockVar)
	}

	var funcVar = "Function Scope Variable"

	fmt.Println(funcVar)
	fmt.Println(packageVar)

	var name = "Test1"
	// name := "Test2" //cannot doing decleration again for name variable
	name, surname := "Test2", "Test3" //adding new value and just assigned new value for name variable.
	fmt.Println(name, surname)

	// fmt.Println(ifBlockVar)
	// myFunc()
}

// func myFunc() {
// 	fmt.Println(packageVar)
// }
