package main

import "fmt"

func main() {

	fmt.Println("1. For Loop")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("2. For Loop")
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	fmt.Println("3. For Loop")
	for {
		fmt.Println("loop")
		break
	}

	fmt.Println("4. For Loop")
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	fmt.Println("5. For Loop")
	j := 0

	for ; j < 11; j += 5 {
		fmt.Println(j)
	}

	fmt.Println(j)
}
