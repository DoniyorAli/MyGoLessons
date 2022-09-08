package main

import "fmt"

func main() {

	myCars := make(map[string]int)
	myCars["MercedesBens"] = 90.000
	myCars["BMW"] = 80.000
	myCars["Bentley"] = 75.000
	myCars["Genesis"] = 65.000
	myCars["Tesla"] = 110.000

	for k, v := range myCars {
		fmt.Printf("Car:%v. Price:%v\n", k, v)
	}
//*============================================================
	fmt.Println()

	keys := make([]string, len(myCars))

	i := 0
	for k := range myCars {
		keys[i] = k
		i += 1
	}
	fmt.Println("KEYS SLICE:", keys)

}