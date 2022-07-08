package main

import "fmt"

func main() {

  fmt.Println("which do you choice a rectangle or triangle ?")
  fmt.Printf("1.Rectangle \n2.Triangle\n")

  var choice int
  fmt.Scanln(&choice)

  switch choice {
    case 1: fmt.Println("Rectangle, Enter the sides of rectangle a,b,c,d:")
      a := 0
      b := 0
      c := 0
      d := 0
      fmt.Scanln(&a,&b,&c,&d)
      fmt.Print("a-->",a,"sm","\nb-->",b,"sm","\nc-->",c,"sm","\nd-->",d,"sm","\n")
      if a == b {

      }
    case 2: fmt.Println("Triangle, Enter the sides of triangle a,b,c:")
      a := 0
      b := 0
      c := 0
      fmt.Scanln(&a,&b,&c)
      fmt.Print("a-->",a,"sm","\nb-->",b,"sm","\nc-->",c,"sm","\n")
  }




  // var rectangle int
  // var triangle int
  // fmt.Scanln(&rectangle)
  
  // if 1 == rectangle {
  //   a := 0
  //   b := 0
  //   c := 0
  //   d := 0
  //   fmt.Println("Enter the sides of rectangle a,b,c,d:")
  //   fmt.Scanln(&a,&b,&c,&d)
  //   fmt.Printf("a-->%d\nb-->%d\nc-->%d\nd-->%d\n", &a, &b, &c, &d)
  // }

  // fmt.Scanln(&triangle)

}