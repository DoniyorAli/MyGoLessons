package main

import "fmt"

func main() {

	var num int
	fmt.Print("Enter the number:")
	fmt.Scanln(&num)

	for ; num > 0; num-- {
		if num % 4 == 0 {
			fmt.Println("Devided:", num)
		}else{
			fmt.Println("Not devided:", num)
		}
	}
}