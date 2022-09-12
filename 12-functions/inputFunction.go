package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	// name, error := reader.ReadString('\n')
	name, _ := reader.ReadString('\n') // _ blank identifier

	fmt.Println("Hello", name)
}
