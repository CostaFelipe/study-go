package main

import "fmt"

func main() {
  a := 1
  b := "hello, World!"
  c := true
  
  fmt.Println(a, b, c)
  fmt.Printf("a:= %v a:= %T\n", a, a)
  fmt.Printf("b:= %v b:= %T", b, b)
}
