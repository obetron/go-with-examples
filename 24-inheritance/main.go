package main

import "fmt"

type person struct {
	name      string
	surname   string
	age       int
	isMarried bool
}

type employee struct {
	person
	department string
}

type manager struct {
	employee
	manageDept string
}

func main() {

	p1 := person{"Michael", "George", 32, true}
	fmt.Println(p1.name, p1.surname, p1.age, p1.isMarried)

	p2 := person{}
	p2.name = "Van"
	p2.surname = "Chris"
	p2.age = 24
	p2.isMarried = false
	fmt.Println(p2.name, p2.surname, p2.age, p2.isMarried)

	e1 := employee{
		person: person{
			name:      "Stewe",
			surname:   "John",
			age:       36,
			isMarried: true,
		},
		department: "Research and Development",
	}
	fmt.Println(e1.name, e1.surname, e1.age, e1.isMarried, e1.department)

	e2 := employee{}
	e2.name = "Brad"
	e2.surname = "Pitt"
	e2.age = 70
	e2.isMarried = false
	e2.department = "Film Industry"
	fmt.Println(e2.name, e2.surname, e2.age, e2.isMarried, e2.department)

	m1 := manager{
		employee: employee{
			person: person{
				name:      "Chris",
				surname:   "Web",
				age:       54,
				isMarried: false,
			},
			department: "Research and Developement",
		},
		manageDept: "Information Department",
	}
	fmt.Println(m1.name, m1.surname, m1.age, m1.isMarried, m1.department, m1.manageDept)

	m2 := manager{}
	m2.name = "Ronaldo"
	m2.surname = "Messi"
	m2.age = 37
	m2.isMarried = true
	m2.department = "Manchester United"
	m2.manageDept = "Barcelona"
	fmt.Println(m2.name, m2.surname, m2.age, m2.isMarried, m2.department, m2.manageDept)

	//if struct needed only one time, than created anonim struct is more sense
	fmt.Println("\nANONIM STRUCT EXAMPLE")
	boss := struct {
		name  string
		admin bool
	}{
		name:  "Angela",
		admin: true,
	}

	fmt.Println(boss.name, boss.admin)

}
