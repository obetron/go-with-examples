package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	target := numRand(1, 100)

	fmt.Println("Guess a number between 1 and 100")
	reader := bufio.NewReader(os.Stdin)

	for attempts := 0; attempts < 10; attempts++ {
		fmt.Print(10-attempts, " guess left!")
		fmt.Println(" Enter your guess")

		input, err := reader.ReadString('\n')
		writeError(err)

		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		writeError(err)

		if num > target {
			fmt.Println("The number less than your guess, enter a number bigger than your guess")
		} else if num < target {
			fmt.Println("The number bigger than your guess, enter a number less than your guess")
		} else {
			fmt.Println("Correct Guess at ", attempts+1, " attemp(s)")
			break
		}
	}

}

func numRand(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func writeError(err error) {
	if err != nil {
		fmt.Println("HATA: ", err)
	}
	return
}
