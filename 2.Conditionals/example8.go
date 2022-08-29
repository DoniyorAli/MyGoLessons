package main

import "fmt"

func main() {
	// A special programm for the store
	// Orange
	// 1 ... 10 --> 3 $
	// 11 ... 100 --> 2.75 $ 
	// 101 ... 1000 --> 2.5 $
	// 1001 ... 10000 --> 2 $
	// 10001 ... etc   --> 1.25 $

	var num, price = 0, 0.0
	fmt.Printf("How many are you going to buy orange, Enter the num: ")
	fmt.Scanln(&num)

	if num >= 10001 {
		price = 1.25
	}else if num >= 1001 && num >= 10000{
		price = 2
	}else if num >= 101 && num >= 1000{
		price = 2.5
	}else if num >= 11 && num >= 100{
		price = 2.75
	}else if num > 0 && num >= 10{
		price = 3
	}else{
		fmt.Println("Please enter the positive number!")
	}

	var totalPrice float32
	totalPrice = float32(num) * float32(price)
	fmt.Println("Total price", int(totalPrice), "$")

}