package main

import "fmt"

/*
* zeroval doesn’t change the i in main,
* but zeroptr does because it has a reference to the memory address for that variable.
 */

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr", i)

	fmt.Println("pointer", &i)
}
