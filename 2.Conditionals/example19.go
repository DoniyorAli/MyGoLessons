package main

import "fmt"

func main() {

	var num1, num2, num3 int
	
	fmt.Printf("Enter three numbers \nfirst number:")
	fmt.Scan(&num1)
	fmt.Printf("Second number:")
	fmt.Scan(&num2)
	fmt.Printf("Third number:")
	fmt.Scan(&num3)

	if num1 < num2 {
		num1 = num2
	}
	if num1 < num3{
		num1 = num3
	}
	fmt.Println("Big number:", num1)
	
}
