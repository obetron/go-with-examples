package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for t := range kvs {
		fmt.Println("key:", t)
	}

	//range on strings iterates over Unicode code points.
	//The first value is the starting byte index of the rune and the second the rune itself.
	for i, c := range "go" {
		fmt.Println(i, c)
	}

	//cities example
	cities := [...]string{"Ankara", "Eskisehir", "Bursa", "Balikesir", "Canakkale"}

	for _, city := range cities {
		fmt.Println(city)
	}

}
