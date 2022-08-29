package main

import "fmt"

func main() {

	var num int
	fmt.Print("Enter the num:")
	fmt.Scanf("%v", &num)
	
	if num % 7 == 0 && num % 5 == 0 && num % 3 == 0 {
		fmt.Println("Good this number devided")
	}else{
		fmt.Println("Not devided")
	}

}