package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	//* FMT
	fmt.Println("Hello Golang")

	//* ARND
	fmt.Println("RANDOM NUMBER:", rand.Intn(10))	//* RANDOM NUMBER: 1
	fmt.Println("RANDOM NUMBER:", rand.Intn(100))	//* RANDOM NUMBER: 87

	//* STRINGS
	fmt.Println(strings.Contains("Hello", "lo"))	//* true
	fmt.Println(strings.Contains("Hello", "la"))	//* false

	fmt.Println(strings.Count("testing", "t"))		//* 2

	fmt.Println(strings.HasPrefix("alikhanov@gmail.com", "ali"))	//* true
	fmt.Println(strings.HasSuffix("MyFileSystem.rar", "rar"))		//* true
	fmt.Println(strings.HasSuffix("What is this qale", "qale"))		//* true

	fmt.Println(strings.Index("Hello", "l"))	//* l --> 2 th index
	fmt.Println(strings.Index("Hello", "o"))	//* o --> 4 th index
	fmt.Println(strings.Index("Hello Golang", " "))	//* probel -->  5 th index

}