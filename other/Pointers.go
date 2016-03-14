package main

import "fmt"


func main()  {
  x := 1
  y := &x


  fmt.Println("The value at x is:", x)
  fmt.Println("The Address of x is:", &x)
  fmt.Println("The value of y is:", y)
  
  *y = 100

  fmt.Println("The value at x is:", x)
  fmt.Println("The Address of x is:", &x)
  fmt.Println("The value of y is:", y)
}
