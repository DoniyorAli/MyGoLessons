package main

import "fmt"

func main() {

	var number = 11
	var user int

	for true {
		fmt.Println("Enter the number between [1....20] and find what kinf of number was entered:")
		fmt.Scanf("%v", &user)

		if number == user {
			fmt.Println("Good you found the number")
			break
		}else {
			fmt.Println("Dangngng!")
		}
	}
}