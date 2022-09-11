package main

import (
	"fmt"
	"math"
)

// const ---> compile time
// var 	 ---> run time

const s string = "constant string"

func main() {

	fmt.Println(s)

	const n = 500000000

	//A numeric constant has no type until it's given one, such as by explicit conversion
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))

	//in this example math.Sin method accepts float64 type so 'n' would be float64 type
	fmt.Println(math.Sin(n))

	//if want to give directly int var into math.Sin function, it would coverted to float64 as shown below.
	notConst := 50000
	fmt.Println(math.Sin(float64(notConst)))
}
