package main

import "fmt"

func main() {  // 643

	var num int

	fmt.Println("Enter the number:")
	fmt.Scanln(&num)

	fmt.Println((num % 10) + (num % 100 / 10) + (num / 100))

}