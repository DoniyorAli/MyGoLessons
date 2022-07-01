package main  // TYPE CHANGING (CONVERSION)

import "strconv"

func main() {

	box := "1"
	num, _ := strconv.Atoi(box)
	println(num + 10)     // -->  11

	one := 11
	bag := float32(one)
	println(bag)         //  -->  +1.100000e+001

	n := 15
	println(float64(n))  //  -->  +1.500000e+001

	w := 22.5
	println(int(w))      //  -->  22

	b := -45.6
	println(uint(b))     //  -->  18446744073709551571

}