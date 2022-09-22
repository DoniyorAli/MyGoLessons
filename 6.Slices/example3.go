package main

import "fmt"

func main() {
	//* New Slice
	var x []int = []int{8, 11, 32, 47, 51}
	fmt.Println(cap(x[:]))		//* 5
	fmt.Println(cap(x[:3]))		//* 5

	//* Make Slice
	box := make([]int, 7)
	fmt.Println(box)	//* [0 0 0 0 0 0 0]
	fmt.Printf("%T \n", box) 	//* Type ===> []int

}