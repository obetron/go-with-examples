package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	var p2 person
	p2 = person{"Yasin", 24}
	fmt.Println(p2)

	var p3 = person{"Ceylan", 34}
	fmt.Println(p3)

	p1 := person{"Eren", 25}
	fmt.Printf("%#v\n", p1)

	fmt.Println(person{"Eren", 35})

	fmt.Println(person{name: "Anil", age: 36})

	fmt.Println(person{name: "Asil"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	fmt.Println("\nSTRUCT PASS BY VALUE EXAMPLE")
	type phone struct {
		brand     string
		model     string
		serialNum string
	}

	phone1 := phone{
		brand:     "Apple",
		model:     "iPhone",
		serialNum: "A5312",
	}
	fmt.Println(phone1)
	phone2 := phone1
	fmt.Println(phone2)

	phone1.brand = "Samsung"
	fmt.Println(phone1)
	fmt.Println(phone2)

}
