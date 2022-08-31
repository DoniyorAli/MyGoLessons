package main

import "fmt"

func main() {

	num := 0
	fmt.Print("Enter the number:")
	fmt.Scanf("%v", &num)
	total := 0
	factorial := 1
	for i := 1; i <= num; i++ {
		total += i
		factorial *= i
	}
	fmt.Println("Total:", total)
	fmt.Println("Factorial:", factorial)
}