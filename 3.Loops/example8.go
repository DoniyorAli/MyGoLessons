package main

import "fmt"

func main() {

	var num int
	fmt.Print("Enter the number:")
	fmt.Scanf("%v", &num)  // 15

	for i := 1; i <= num; i += 1 {
		if i % 2 == 0 {
			fmt.Println(i) // 2,4,6,8,10,12,14
		}
	}
}