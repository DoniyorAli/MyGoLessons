package main

import "fmt"

func main() {

	numbers := [...]int{1,2,3,4,5,6,7,8,9,10,12,13,14,15,16,17,18,19,20}

	for _, v := range numbers {
		if v % 2 == 0 {
			fmt.Println("An even number:", v)
		}
	}
}