package main

import "fmt"

func main()  {
  s := "Hello"
  sb := []byte(s)
  fmt.Printf("%v , %T\n", s, s)
  fmt.Printf("%v, %T\n", sb, sb)

  for _, v := range sb {
    fmt.Printf("%v - %#U - %#X\n", v, v, v)
  }
}
