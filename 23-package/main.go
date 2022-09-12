package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("seafood", "food"))
	fmt.Println(strings.Count("gopher", "e"))
	fmt.Println(strings.ToUpper("gopher"))
}
