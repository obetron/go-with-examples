package main

import "fmt"

func main() {
	i := 7
	if i%2 == 0 {
		fmt.Println(i, " is event")
	} else {
		fmt.Println(i, " is odd")
	}

	j := 8
	if j%4 == 0 {
		fmt.Println(j, " is divided by 4")
	} else {
		fmt.Println(j, " is not divided by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, " is negative")
	} else if num < 10 {
		fmt.Println(num, " has 1 digit")
	} else {
		fmt.Println(num, " has multiple digits")
	}
}
