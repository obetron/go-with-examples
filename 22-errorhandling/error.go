package main

import (
	"errors"
	"fmt"
)

func main() {

	result, err := isEvenNum(10)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func isEvenNum(num int) (string, error) {
	if num%2 != 0 {
		return "", errors.New("HATA: number should be even")
	}

	return "The number is even", nil
}
