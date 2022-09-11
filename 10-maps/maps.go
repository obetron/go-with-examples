package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 3
	m["k2"] = 12

	fmt.Println("map:", m)
	fmt.Println("len:", len(m))

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	delete(m, "k2")
	fmt.Println("map:", m)

	//present
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	//initialing new map in the same line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
	fmt.Println("len:", len(n))
	fmt.Println("foo val:", n["foo"])
}
