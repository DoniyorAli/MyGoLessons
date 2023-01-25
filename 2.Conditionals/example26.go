package main

import "fmt"

func main() {

	var num int
	fmt.Print("Enter the five rooms number:")
	fmt.Scanln(&num)

	one := ((num % 10) * 10000)
	two := (((num % 100) / 10) * 1000)
	three := (((num % 1000) / 100) * 100)
	four := (((num % 10000) / 1000) * 10)
	five := ((num / 10000))
	fmt.Println(one, two, three, four, five)

	num = ((num % 10) * 10000) + (((num % 100) / 10) * 1000) + (((num % 1000) / 100) * 100) + (((num % 10000) / 1000) * 10) + ((num / 10000))
	fmt.Println(num)

}
// Example number: 12345 ---> 54321
