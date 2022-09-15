package main

import "fmt"

func writeHello(chan1 chan string) {
	chan1 <- "Hello World"
}

func main() {
	myChannel := make(chan string)

	go writeHello(myChannel)
	fmt.Println(<-myChannel)
}
