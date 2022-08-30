package main

import "fmt"

func main() {

	var num int
	fmt.Printf("Enter the num:")
	fmt.Scanln(&num)  // 10

	for i := num; i > 0; i-- {
		fmt.Println(i)   // 10,9,8,7,6,5,4,3,2,1
	}

}