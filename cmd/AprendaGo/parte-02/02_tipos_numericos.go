package main

import "fmt"

func main()  {
  a := "e"
  b := "é"
  c := "#23446"
  fmt.Printf("%v, %v, %v\n", a, b, c)

  d := []byte(a)
  e := []byte(b)
  f := []byte(c)

  fmt.Printf("%v, %v, %v", d, e, f)
}
