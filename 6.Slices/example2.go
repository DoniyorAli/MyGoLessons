package main

import "fmt"

func main() {

	var x [5]int = [5]int{1, 2, 3, 4, 5}
	var b []int = x[:]
	fmt.Println("My Slice:", b[:])		//* My Slice: [1 2 3 4 5]
	fmt.Println("LENGTH:", len(b))		//* LENGTH: 5
	fmt.Println("CAPACITY:", cap(b))	//* CAPACITY: 5

	b = append(b, 123)
	fmt.Println(b)	//* [1 2 3 4 5 123]
	fmt.Println("CAPACITY:", cap(b))	//* CAPACITY: 10

	var d []int = x[0:]
	fmt.Println(d) 		//* 5

	var c []int = x[0:]
	c = append(c, 123)
	fmt.Println(cap(c))		//* 10
}
//? QUESTION
