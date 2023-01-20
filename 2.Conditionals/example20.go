package main

import "fmt"

func main() {

	year := 0
	fmt.Printf("Enter year:")
	fmt.Scanf("%d", &year)

	if year % 400 == 0 || (year % 4 == 0 && year % 100 != 0) {
		fmt.Println("Leap year")
	}else {
		fmt.Println("Not leap year!")
	}
	
}
