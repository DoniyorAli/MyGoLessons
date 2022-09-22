package main

import "fmt"

func main() {

	var x [5]int = [5]int{1, 2, 3, 4, 5}
	var s []int = x[1:3]
	fmt.Println("My Slice:", s[:])		//* My Slice: [2 3]
	fmt.Println("LENGTH:", len(s))		//* LENGTH: 2
	fmt.Println("CAPACITY:", cap(s))	//* CAPACITY: 4
	fmt.Println(s[:cap(s)])	//* [2 3 4 5]
	fmt.Println(s[1:])	//* [3]

	s = append(s, 123)
	fmt.Println("Append num:", s)
	fmt.Println("CAPACITY:", cap(s))	//* CAPACITY: 4

//*==============================================================================================
	fmt.Println()

	nums := make([]int, 5)
	nums[1] = 11
	fmt.Println(nums)	// [0 11 0 0 0]
	nums = append(nums, 789)
	fmt.Println(nums)	// [0 11 0 0 0 789]

	fmt.Println("LENGTH:", len(nums))		//* LENGTH: 6
	fmt.Println("CAPACITY:", cap(nums))		//* CAPACITY: 10
}