package main

import "fmt"

type Person struct {
	id        int
	name      string
	age       int
	isMarried bool
	kidsName  []string
}

func main() {

	var employee1 struct {
		name      string
		age       int
		isMarried bool
	}

	fmt.Println(employee1)
	fmt.Printf("%#v\n", employee1)

	employee1.name = "Test"
	employee1.age = 23
	employee1.isMarried = false

	fmt.Printf("%#v\n", employee1)
	fmt.Println(employee1.name)
	fmt.Println(employee1.age)
	fmt.Println(employee1.isMarried)

	type employee2 struct {
		name      string
		age       int
		isMarried bool
	}

	var e1 employee2 //zero value defining with construction
	e1.name = "Test1"
	e1.age = 34
	e1.isMarried = true

	var e2 employee2
	e2.name = "Test3"
	e2.age = 54
	e2.isMarried = false

	fmt.Printf("%#v\n", e1)
	fmt.Println(e1)
	fmt.Printf("%#v\n", e2)
	fmt.Println(e2)

	fmt.Println("\nPERSON DEFINITION")
	person1 := Person{1, "Eren", 35, true, []string{"Asil"}}
	person2 := Person{2, "Ahmet", 54, false, []string{}}
	fmt.Println(person1)
	fmt.Println(person2)

	person3 := Person{} //zero value defining with literal constraction
	fmt.Println(person3)

}
