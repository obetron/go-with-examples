package main

import (
	"fmt"
	"time"
)


func main(){

	var now time.Time = time.Now()
	var year, month, day = now.Date()

	fmt.Println("Date:", day, month, year)

}