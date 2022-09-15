package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	printXYInSeq()

	printXYConcurrently()
	time.Sleep(time.Second)

}

func printXYInSeq() {
	printX()
	printY()
}

func printXYConcurrently() {
	go printX()
	go printY()
}

func printXYWaitGroup() {

	//always print X first
	go printX()

}

func printX() {
	for i := 0; i < 20; i++ {
		fmt.Print("X")
	}
	fmt.Println()
}

func printY() {
	for i := 0; i < 20; i++ {
		fmt.Print("Y")
	}
	fmt.Println()
}
