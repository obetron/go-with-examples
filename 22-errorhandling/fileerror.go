package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Dosyamiz,", file)
	}

	file1, err := os.Open("test1.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Dosyamiz,", file1)
	}
}
