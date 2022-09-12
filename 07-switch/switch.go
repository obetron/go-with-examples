package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is the weekend")
	default:
		fmt.Println("It is a weekaday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am a int")
		case string:
			fmt.Println("I am a string")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("Eren")
	whatAmI(4.5)

	switch x := 25; {
	case x < 20:
		fmt.Printf("%d kucuktur 20\n", x)
		fallthrough

	case x < 50:
		fmt.Printf("%d kucuktur 50\n", x)
		fallthrough

	case x < 75:
		fmt.Printf("%d kucuktur 75\n", x)
		fallthrough

	case x < 100:
		fmt.Printf("%d kucuktur 100\n", x)

	}

}
