package main

import "fmt"

func main()  {
  a := 200
  fmt.Printf("%b - %#x - %v\n", a, a, a)
  y := a << 1
  fmt.Printf("%b - %#x - %v\n", y, y, y)

}
