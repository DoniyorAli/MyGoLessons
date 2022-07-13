package main // Scanln  and TIME

import "fmt"
import "time"


func main() {

	var name string   // DECLARE --> "" bo'sh
	fmt.Scanln(&name)
	fmt.Println("Hello --> ", name) // Hello --> Alxan


	vaqt := time.Date(2022, time.February, 11, 23, 47, 78, 45, time.Local)
	fmt.Println(vaqt)

	fmt.Println(time.Now())  // Hozirgi vaqt

}