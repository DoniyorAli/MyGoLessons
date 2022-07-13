package main      // SWITCH CASE

import "fmt"

func main() {

	fmt.Println("1.Your balance \n2.Your SMS\n3.Your MB")
	var choice int
	fmt.Println("Choose number: ")
	fmt.Scanf("%d", &choice)

	switch choice {
	case 1:
		fmt.Println("Your balance is 1200$")
		break   // --> break or NOT NEEDED
	case 2:
		fmt.Println("123 SMS came to you")
		break
	case 3:
		fmt.Println("Limit 12345 MB")
		break
	default:
		fmt.Println("This number is No! DANG 🤦‍♂️")
	}

}