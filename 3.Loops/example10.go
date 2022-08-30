package main

import "fmt"

func main() {

	tickets := 500
	yes := "Y"
	for yes == "y" || yes == "Y" {
		fmt.Print("Do you want to buy ticktes? [Y/N]:")
		fmt.Scanln(&yes)
		if yes == "n" || yes == "N" {
			fmt.Println("Thank you for your attention")
			break
		}

		fmt.Print("How many do you want to buy tickets!")
		boughtTicket := 0
		fmt.Scanln(&boughtTicket)
		if boughtTicket <= tickets {
			tickets -= boughtTicket
			fmt.Println(boughtTicket," was boughten")
			fmt.Println("Number of remaining tickets:", tickets)
		}else if tickets == 0 {
			fmt.Println("Tickets are not avalable!")
			break
		}else{
			fmt.Println("There are not many tickets available!")
		}
	}
}