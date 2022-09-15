package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//there another 2 go routines
	wg.Add(1)

	go printX()
	wg.Wait()
	printY()

}

func printX() {
	for i := 0; i < 20; i++ {
		fmt.Print("X")
	}
	fmt.Println()
	wg.Done()
}

func printY() {
	for i := 0; i < 20; i++ {
		fmt.Print("Y")
	}
	fmt.Println()
}
