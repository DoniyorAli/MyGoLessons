package main

import "fmt"

func main() {

	var height float32
	var weight float32
	fmt.Printf("Enter your height: ")
	fmt.Scanf("%f", &height)
	fmt.Printf("Enter your weight:")
	fmt.Scanf("%f", &weight)

	index := weight / (height * height)
	fmt.Println(index)

	if index < 18.5 {
		fmt.Println("Weak")
	}else if 18.5 < index && index < 25 {
		fmt.Println("Normal")
	}else if 25 < index && index < 30 {
		fmt.Println("Generous")
	}else if index > 30 {
		fmt.Println("Heavy")
	}else {
		fmt.Println("Please enter correctly!")
	}
}