package main

import (
	"fmt"
	"reflect"
)

func main() {

	var A int = 11
	var B int = 10

	if reflect.TypeOf(A) != reflect.TypeOf(B){
		fmt.Println("Good")
	}

}