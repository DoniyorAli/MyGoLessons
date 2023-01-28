package main

import "fmt"

func main() {

	var q,w,e,r,t int
	fmt.Print("Enter 5 numbers:")
	fmt.Scanln(&q,&w,&e,&r,&t)

	bigNum := q
	if bigNum < w {
		bigNum = w
	}
	if bigNum < e {
		bigNum = e
	}
	if bigNum < r {
		bigNum = r
	}
	if bigNum < t {
		bigNum = t
	}
	fmt.Println(bigNum)

}
