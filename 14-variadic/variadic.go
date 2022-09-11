package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func main() {

	sum(1, 3, 4)
	sum(1, 3, 5, 6)

	nums := []int{1, 5, 8, 0}
	sum(nums...)

}
