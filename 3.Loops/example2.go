package main       // FOR ---> WHILE. In the WHILE --> FOR view
			       // WHILE ---> FOR ko'rinishida
import "fmt"

func main() {

	i := 0
	for i < 10 {
		fmt.Println("-->", i) // 0, 1, 2, 3, 4, 5, 6, 7, 8, 9
		i++
	}
//_________________________________________________________________
	fmt.Print("=======\n")

	var j int = 10
	for j >= 1 {
		fmt.Println("-->", j) // 10, 9, 8, 7, 6, 5, 4, 3, 2, 1
		j -= 1
	}
//_________________________________________________________________
	fmt.Print("========\n")

	var num = 0     //  var num int = 0
	var total = 0   //  var total int =0
	for num <= 10 {
		fmt.Println("-->", num) // 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
		total += num  // 0+1+2+3+4+5+6+7+8+9+10 = 55
		num += 1
	}
	fmt.Println("Total:", total) // output Total: 55
//_________________________________________________________________
	fmt.Print("========\n")

	box := 10

	for box < 100 {
		fmt.Println(box) // 10,15,20,25, ... 95
		box += 5     // 10+5 ... 15+5 ... 20+5 ... 25+5 ...... 95
	}
}