package main

import "fmt"

func main() {

	var seconds int32
	fmt.Printf("Enter the seconds: ")
	fmt.Scan(&seconds)

	// Minute := seconds / 60
	// fmt.Println(Minute, "Minute:")

	// Hour := seconds / 3600
	// fmt.Println(Hour, "Hour:")

	// Minute := seconds / 60
	// Seconds := seconds % 60
	// fmt.Println(Minute, "Minute:", Seconds, "Second:")

	// Hour := seconds / 3600
	// Seconds := seconds % 3600 
	// fmt.Println(Hour,"Hour:", Seconds,"Seconds:")

	hour := seconds / 3600
	min := seconds % 3600 / 60
	sec := seconds % 3600 % 60
	fmt.Println(hour,".hour", min,".minute", sec,".second")

}
