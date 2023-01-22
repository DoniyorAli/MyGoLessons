package main

import "fmt"

func main() {
	
	var numbers int
	fmt.Printf("Enter three digit number:")
	fmt.Scan(&numbers)
	if numbers > 0 {
		total := (numbers / 100) + (numbers / 10 % 10) + (numbers % 10)
		fmt.Println(total)
	}else {
		fmt.Println("Only three digit numbers enter!")
	}
}
