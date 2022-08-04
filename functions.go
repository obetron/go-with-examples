package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	sum := plus(2, 5)
	fmt.Println("2 + 5 = ", sum)

	sumOfThree := plusPlus(2, 3, 5)
	fmt.Println("2 + 3 + 5 =", sumOfThree)
}
