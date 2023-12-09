package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
)

func main(){

	fmt.Println("Enter grade of Student")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n') // seklinde yapilacak bir tnaimlama sonucunda eger err degeri handle edilmiyor yada kullanillmiyorsa kod hata verecek ve calismayacaktir.
	// input, _ := reader.ReadString('\n') // _ seklinde verdigimizde aslinda go ya bu degeri kullanmayacagimizi belirtmis oluyoruz.
	if err != nil{
		log.Fatal("Error: ", err)
	}

	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 32)

	if err != nil {
		log.Fatal("Error: ", err)
	}

	fmt.Println("Student Grade is: ", grade)
	if grade == 100 {
		fmt.Println("Perfect")
	} else if grade >= 60 && grade < 100 {
		fmt.Println("Pass")
	} else if grade >= 40 && grade < 60 {
		fmt.Println("Average")
	} else if grade >= 0 && grade < 40 {
		fmt.Println("Bad")
	} else {
		fmt.Println("Fail")
	}

}


