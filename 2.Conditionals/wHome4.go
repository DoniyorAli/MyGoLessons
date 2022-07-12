package main

import "fmt"

func main() {

  fmt.Println("which do you choice a rectangle or triangle ?")
  fmt.Printf("1.Rectangle \n2.Triangle\n")

  var choice int
  fmt.Scanln(&choice)

  switch choice {
    case 1: fmt.Println("Rectangle, Enter the sides of rectangle a,b,c,d:")
      a1 := 0
      b1 := 0
      a2 := 0
      b2 := 0

      fmt.Printf("a-->")
      fmt.Scanf("%d", &a1)
      fmt.Printf("b-->")
      fmt.Scanf("%d", &b1)
      fmt.Printf("c-->")
      fmt.Scanf("%d", &a2)
      fmt.Printf("d-->")
      fmt.Scanf("%d", &b2)

      fmt.Print("a-->",a1,"sm","\nb-->",b1,"sm","\nc-->",a2,"sm","\nd-->",b2,"sm","\n")

      // rectangle := a1 * b1

      if (a1 == a2 && b1 == b2) && a1 == b1{
        fmt.Println("Square")
      }else if (a1 == a2 && b1 == b2) && a1 != b1{
        fmt.Println("Rectangle")
      }else{
        fmt.Println("It is neither Square nor Rectangle!")
      }

    case 2: fmt.Println("Triangle, Enter the angles of the triangle a,b,c:")
      a := 0
      b := 0
      c := 0
      fmt.Printf("a-->")
      fmt.Scanf("%d", &a)
      fmt.Printf("b-->")
      fmt.Scanf("%d", &b)
      fmt.Printf("c-->")
      fmt.Scanf("%d", &c)

      fmt.Print("a-->",a,"sm","\nb-->",b,"sm","\nc-->",c,"sm","\n")

      if a == b && b == c && a == c{
        fmt.Println("All sides are equal triangle")
      }else if a == c {
        fmt.Println("Triangle with two equal sides triangle")
      }else {
        fmt.Println("Simple triangle")
      }
  }
}