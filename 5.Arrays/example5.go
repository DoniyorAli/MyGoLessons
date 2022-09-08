package main

import "fmt"

func main() {

	myArray := [10]int{11,22,33,44,55,}

	fmt.Println("LENGTH:", len(myArray))	// 10
	fmt.Println("LENGTH:", cap(myArray))	// 10

}