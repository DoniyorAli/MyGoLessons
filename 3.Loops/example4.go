package main      // RANGE

import "fmt"

func main() {

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, value := range pow {
		fmt.Println(i,"===>", value)
	}
//========================================================
	fmt.Print("=============\n")
	var users = []string{"Ali", "Dani", "John", "Leyla"}
	for _, v := range users {
		fmt.Println("Value===>", v)
	}
//========================================================
	fmt.Print("============\n")
	var cities = map[string]string{"Uzb":"Tashkent", "USA":"NewYork", "KR":"Korea", "Ros":"Moskva", "GR":"Berlin"}
	for key, value := range cities {
		fmt.Print(key,":", value,"\n")
	}

}