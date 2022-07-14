package main

import "fmt"

func main() {
	
	// Military admission
	// Your gender Male, and you must be taller than 180
	// Your gender Woman, and you must be taller than 160
	// Output: Hello name, Your gender is male, Your height 180
	// Output: Hello name, Your gender is woman, Your height 160

	var name, gender, height = "", "", 0
	fmt.Printf("Enter your name: ")
	fmt.Scanf("%v", &name)
	fmt.Printf("Enter your gender: ")
	fmt.Scanf("%v", &gender)
	fmt.Printf("Enter your height: ")
	fmt.Scanf("%d", &height)

	if (gender == "male" || gender == "Male") && height >= 180 && height <= 210{
		fmt.Printf("Hello %v, Your gender is %v, Your height %d\n", name, gender, height)
		fmt.Println("Good you are accepted!")
	}else if (gender == "male" || gender == "Male") && height < 180{
		fmt.Printf("Hello %v, Your gender is %v, Your height %d\n", name, gender, height)
		fmt.Println("Unfor you are not accepted!!!")
	}else if (gender == "woman" || gender == "Woman") && height >= 160{
		fmt.Printf("Hello %v, Your gender is %v, Your height %d\n", name, gender, height)
		fmt.Println("Good you are accepted!!!")
	}else if (gender == "woman" || gender == "Woman") && height < 160{
		fmt.Printf("Hello %v, Your gender is %v, Your height %d\n", name, gender, height)
		fmt.Println("Unfor ou are not accepted!!!")
	}else{
		fmt.Println("Please fill out the questionnaire correctly or wrong entered height!")
	}

}