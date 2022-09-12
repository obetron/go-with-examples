package main

import "fmt"

func greeting() {
	fmt.Println("Hello World")
}

func greetingWithName(name string) string {
	greetingStr := "Hello, " + name
	return greetingStr
}

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	greeting()
	greeting := greetingWithName("Test")
	fmt.Println(greeting)

	sum := plus(2, 5)
	fmt.Println("2 + 5 = ", sum)

	sumOfThree := plusPlus(2, 3, 5)
	fmt.Println("2 + 3 + 5 =", sumOfThree)

	x := 5
	y := 10
	z := 15

	total := plusPlus(x, y, z)
	fmt.Printf("%d + %d + %d = %d\n", x, y, z, total)
}
