package main

import "fmt"

func main() {

	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)

	switch name {
	case "Alikhan":
		fmt.Println("Hello Alikhan")
	case "Temur":
		fmt.Println("Hello Temur")
	case "Jack":
		fmt.Println("Hello Jack")
	case "Anna":
		fmt.Println("Hello Anna")
	default:
		fmt.Println("No such name 🤦‍♂️")
	}

}