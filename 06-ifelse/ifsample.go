package main

import "fmt"

func main() {

	x := 20

	if x%2 == 0 {
		fmt.Println(x, "cift bir sayidir")
		return
	}

	fmt.Println("x tek bir sayidir.")

	fmt.Println("burada sonraki islemler yapilmaya devam edilecektir.")

}
