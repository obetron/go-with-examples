package main

import "fmt"

func main() {

	myMap1 := map[string]int{
		"Ankara":    06,
		"Bursa":     16,
		"Eskisehir": 26,
	}
	fmt.Println(myMap1)
	fmt.Println("Value of key which included in myMap1 --> ", myMap1["Ankara"])
	fmt.Println("Value of key which not included in myMap1 -->", myMap1["Istanbul"]) // return 0 because there no key like Istanbul in myMap1

	fmt.Println("\nNIL & EMPTY MAP EXAMPLE")
	var myMap2 map[string]int
	fmt.Printf("%#v\n", myMap2)

	myMap3 := make(map[string]int)
	fmt.Printf("%#v\n", myMap3)

	fmt.Println()
	m := make(map[string]int)
	m["k1"] = 3
	m["k2"] = 12

	fmt.Println("map:", m)
	fmt.Println("len:", len(m))

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	delete(m, "k2")
	fmt.Println("map:", m)

	fmt.Println("\nPRESENT EXAMPLE")
	//present
	fmt.Println("m map:", m)
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	fmt.Println("\nINITIALIZING NEW MAP IN THE SAME LINE EXAPLE")
	//initialing new map in the same line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
	fmt.Println("len:", len(n))
	fmt.Println("foo val:", n["foo"])

	fmt.Println("\nMAP PASS BY REFERENCE EXAMPLE")
	myMap4 := map[string]int{
		"Ankara":    6,
		"Eskisehir": 26,
		"Bursa":     16,
		"Canakkale": 17,
	}
	fmt.Println("myMap4:", myMap4)
	myMap5 := myMap4
	fmt.Println("myMap5:", myMap5)

	fmt.Println("deleting Eskisehir from myMap4...")
	delete(myMap4, "Eskisehir")
	fmt.Println("myMap4:", myMap4)
	fmt.Println("myMap5:", myMap5)

	fmt.Println("\nKIND of COMPLEX MAP EXAMPLE")
	myComplexMap := map[string][]string{
		"Ankara":    []string{"Cankaya", "Yenimahalle", "Bala", "Haymana", "Etimesgut", "Eryaman"},
		"Eskisehir": []string{"Yenimahalle", "Odunpazari", "Palu", "Batikent"},
		"Bursa":     []string{"Nilufer"},
	}

	fmt.Println(myComplexMap["Ankara"])
	fmt.Println(myComplexMap["Eskisehir"][0])

	for key, value := range myComplexMap {
		fmt.Println("Il:", key)
		fmt.Println("Ilceler:")
		for _, v := range value {
			fmt.Println("\t", v)
		}
	}

}
