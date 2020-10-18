package main

import "fmt"

func main()  {
  for i := 0; i < 5; i++ {
    fmt.Printf("%b - %v - %#x\n", i, i, i)
  }
}
