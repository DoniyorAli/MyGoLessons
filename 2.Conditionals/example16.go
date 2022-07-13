package main

import "fmt"

func main() {

	var firstBall float32
	var secondBall float32
	var thirdBall float32

	fmt.Println("Enter your exam balls and final ball:")
	fmt.Printf("1. Exam --> ")
	fmt.Scanln(&firstBall)
	fmt.Printf("2. Exam --> ")
	fmt.Scanln(&secondBall)
	fmt.Printf("3 Final --> ")
	fmt.Scanln(&thirdBall)

	totalBall := firstBall + secondBall + thirdBall

	switch {
	case 85 <= totalBall && totalBall <= 100: 
		fmt.Println("Your ball --> 5")
	case 80 <= totalBall && totalBall < 85: 
		fmt.Println("Your ball --> 4.5")
	case 75 <= totalBall && totalBall < 80: 
		fmt.Println("Your ball --> 4")
	case 70 <= totalBall && totalBall < 75: 
		fmt.Println("Your ball --> 3.5")
	case 65 <= totalBall && totalBall < 70: 
		fmt.Println("Your ball --> 3")
	case 60 <= totalBall && totalBall < 65: 
		fmt.Println("Your ball --> 2.5")
	case 55 <= totalBall && totalBall < 60: 
		fmt.Println("Your ball --> 2")
	case 55 > totalBall: 
		fmt.Println("Your failed the exam!!!")
	}

}