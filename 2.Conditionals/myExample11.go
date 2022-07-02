package main

import "fmt"

func main() {

    fmt.Print("Enter the second: ")
    var sec int = 4000 
    fmt.Scan(&sec)
    
//    var hour int 
//    var min int
//    var sek int 
    
    hour := sec / 3600
    min := sec % 3600 / 60
    sek := sec % 3600 % 60
    fmt.Println(hour, "hour",min, "minut", sek, "sekunt")
}