package main

import "fmt"

const (
  a = iota * 10
  b
  c
)

func main()  {
  fmt.Println(a, b, c)
}
