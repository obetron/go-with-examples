package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set", s)
	fmt.Println("get", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	fmt.Println("Final Lenght: ", len(s))

	intSlices := make([]int, 4)
	intSlices[0] = 2
	intSlices[1] = 3
	intSlices[2] = 234

	intSlices = append(intSlices, 434)

	fmt.Println(intSlices)
	//output: [2 3 234 0 434]
}
