package main

import (
	"fmt"
	"unsafe"
)

type size2 struct {
	id   int
	name string
}

type size3 struct {
	id   int
	name string
	age  int16
}

type size4 struct {
	id     int
	name   string
	age    int16
	gender string
}

type size1 struct {
	id     int
	name   string
	age    int16
	gender string
	active bool
}

type size5 struct {
	id      int
	name    string
	age     int16
	gender  string
	active  bool
	passive bool
	a       bool
	b       bool
	c       bool
	d       bool
	e       bool
	f       bool
	g       bool
}

type size6 struct {
	var1 bool
	var2 int
	var3 string
}

type size7 struct {
	var1 string
	var2 int
	var3 bool
}

type size8 struct {
	var1 string
	var2 int
	var3 bool
}

type size9 struct {
	flag    bool
	counter int16
	pi      float32
}

type size10 struct {
	flag    bool
	counter int32
	pi      float32
}

type size11 struct {
	flag    bool
	counter int64
	pi      float32
}

type size12 struct {
	flag    bool
	counter int64
	flag1   bool
	pi      float32
}

type size13 struct {
	counter int64
	pi      float32
	flag    bool
	flag1   bool
}

func main() {

	var s1 size1
	var s2 size2
	var s3 size3
	var s4 size4
	var s5 size5

	var s6 size6
	var s7 size7
	var s8 size8

	var s9 size9
	var s10 size10
	var s11 size11
	var s12 size12
	var s13 size13

	var intSize int
	var stringSize string
	var int16Size int16
	var boolSize bool

	fmt.Println("Size 6:", unsafe.Sizeof(s6))
	fmt.Println("Size 7:", unsafe.Sizeof(s7))
	fmt.Println("Size 8:", unsafe.Sizeof(s8))
	fmt.Println("**************************")
	fmt.Println()

	fmt.Println("Int Size:", unsafe.Sizeof(intSize))
	fmt.Println(" --------")
	fmt.Println("|xxxxxxxx|")
	fmt.Println(" --------")
	fmt.Println("String Size:", unsafe.Sizeof(stringSize))
	fmt.Println(" -------- --------")
	fmt.Println("|ssssssss|ssssssss|")
	fmt.Println(" -------- --------")
	fmt.Println("Int16 Size:", unsafe.Sizeof(int16Size))
	fmt.Println(" --")
	fmt.Println("|xx|")
	fmt.Println(" --")
	fmt.Println("Bool Size:", unsafe.Sizeof(boolSize))
	fmt.Println(" -")
	fmt.Println("|x|")
	fmt.Println(" -")
	fmt.Println("**************************")
	fmt.Println()

	fmt.Println("int Size: ", unsafe.Sizeof(intSize))
	fmt.Println("int, string Size: ", unsafe.Sizeof(s2))
	fmt.Println(" -------- -------- --------")
	fmt.Println("|xxxxxxxx|ssssssss|ssssssss")
	fmt.Println(" -------- -------- --------")
	fmt.Println("int, string, int16 Size:", unsafe.Sizeof(s3))
	fmt.Println(" -------- -------- -------- --------")
	fmt.Println("|xxxxxxxx|ssssssss|ssssssss|xx------|")
	fmt.Println(" -------- -------- -------- --------")
	fmt.Println("int, string, int16, string Size:", unsafe.Sizeof(s4))
	fmt.Println(" -------- -------- -------- -------- -------- --------")
	fmt.Println("|xxxxxxxx|ssssssss|ssssssss|xx------|ssssssss|ssssssss")
	fmt.Println(" -------- -------- -------- -------- -------- --------")
	fmt.Println("int, string, int16, string, bool Size: ", unsafe.Sizeof(s1))
	fmt.Println(" -------- -------- -------- -------- -------- -------- --------")
	fmt.Println("|xxxxxxxx|ssssssss|ssssssss|xx------|ssssssss|ssssssss|b-------")
	fmt.Println(" -------- -------- -------- -------- -------- -------- --------")
	fmt.Println("int, string, int16, string, bool, bool Size:", unsafe.Sizeof(s5))
	fmt.Println(" -------- -------- -------- -------- -------- -------- -------- --------")
	fmt.Println("|xxxxxxxx|ssssssss|ssssssss|xx------|ssssssss|ssssssss|bbbbbbbb|b-------")
	fmt.Println(" -------- -------- -------- -------- -------- -------- -------- --------")
	fmt.Println("****************************************************************************")
	fmt.Println()

	fmt.Println("bool, int16, float32 Size:", unsafe.Sizeof(s9), "bits")
	fmt.Println("bool, int32, float32 Size:", unsafe.Sizeof(s10), "bits")
	fmt.Println("bool, int64, float32 Size:", unsafe.Sizeof(s11), "bits")
	fmt.Println("bool, int64, bool, float32 Size:", unsafe.Sizeof(s12), "bits") //14 bit padding
	fmt.Println("int64, float32, bool, bool Size:", unsafe.Sizeof(s13), "bits") //6 bit padding

}
