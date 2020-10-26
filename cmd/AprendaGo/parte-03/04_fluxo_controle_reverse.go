package main

import "fmt"

func main()  {
  for j := 10; j > 0; j-- {
    fmt.Printf("%v , %b, %#X\n", j, j, j)
  }
}
