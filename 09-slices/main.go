package main

import "fmt"

func main() {

	var mySlice1 []int
	fmt.Println("Slice:", mySlice1) //output: []

	var mySliceWithMake = make([]int, 5)
	fmt.Println("Slice with created make:", mySliceWithMake) //output: [0 0 0 0 0]

	// ARRAYS are Pass By Value
	fmt.Println("\nARRAY")
	var arr1 = [4]int{1, 2, 3, 4}
	fmt.Println("arr1", arr1)

	arr2 := arr1
	fmt.Println("arr2", arr2)

	arr2[1] = 100
	fmt.Println("arr2", arr2) //output: [1, 100, 3, 4]
	fmt.Println("arr1", arr1) //output: [1, 2, 3, 4]

	// SLICES are Pass By Reference
	fmt.Println("\nSLICE")
	slc1 := []int{1, 2, 3, 4}
	fmt.Println("slc1:", slc1)

	slc2 := slc1
	fmt.Println("slc2:", slc2)

	slc2[1] = 100
	fmt.Println("slc2:", slc2) //output: [1, 100, 3, 4]
	fmt.Println("slc1:", slc1) //output: [1, 100, 3, 4]

	fmt.Println("\n\nPASS BY REFERENCE EXAMPLE")
	underArray := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(underArray)

	mySlc1 := underArray[2:5]
	fmt.Println(mySlc1)
	mySlc2 := underArray[:6]
	fmt.Println(mySlc2)
	mySlc3 := underArray[3:]
	fmt.Println(mySlc3)
	mySlc4 := underArray[:]
	fmt.Println(mySlc4)

	//mySlc first element, mySlc2 and mySlc4 third elements changed to 100
	//because of pass by reference
	mySlc1[0] = 100
	fmt.Println(mySlc1, "<-- first element changed to 100")
	fmt.Println(mySlc2, "<-- third element changed to 100")
	fmt.Println(mySlc4, "<-- third element changed to 100")

	fmt.Println("\nSLICE APPEND EXAMPLE")
	myAppendSlc1 := []int{1, 2, 3}
	fmt.Print("myAppendSlc1: ")
	fmt.Println(myAppendSlc1)

	myAppendSlc2 := append(myAppendSlc1, 4, 5)
	fmt.Print("myAppendSlc2: ")
	fmt.Println(myAppendSlc2)

	myAppendSlc1[0] = 100
	fmt.Print("myAppendSlc1: ")
	fmt.Println(myAppendSlc1)

	fmt.Print("myAppendSlc2: ")
	fmt.Println(myAppendSlc2)

	fmt.Println("\nAPPEND SLICE TO ANOTHER SLICE")
	myAppendSlc3 := []int{1, 2, 3}
	fmt.Println("Slice 1", myAppendSlc3)
	myAppendSlc4 := []int{4, 5}
	fmt.Println("Slice 2", myAppendSlc4)

	myAppendSlc3 = append(myAppendSlc3, myAppendSlc4...)
	fmt.Println(myAppendSlc3)

}
