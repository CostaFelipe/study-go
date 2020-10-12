package main

import "fmt"

var (
  a int
  b float32
  c string
  d bool
)

func main()  {
  fmt.Printf("%v, %T\n", a, a)
  fmt.Printf("%v, %T\n", b, b)
  fmt.Printf("%v, %T\n", c, c)
  fmt.Printf("%v, %T", d, d)
}
