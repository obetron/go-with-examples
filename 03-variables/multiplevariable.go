package main

import "fmt"

func main() {
	var strVal1 string = "Name Value1"
	var intVal1 int = 1
	var boolVal1 bool = true
	fmt.Println(strVal1, intVal1, boolVal1)

	var strVal2, intVal2, boolVal2 = "Name Value2", 2, false
	fmt.Println(strVal2, intVal2, boolVal2)

	var (
		strVal3  string = "Name Value3"
		intVal3  int    = 3
		boolVal3 bool   = true
	)
	fmt.Println(strVal3, intVal3, boolVal3)

	strVal4, intVal4, boolVal4 := "Name Value4", 4, false
	fmt.Println(strVal4, intVal4, boolVal4)
}
