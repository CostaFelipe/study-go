package main

import "fmt"

func mult(a1, a2 int) int {
  multiplicao := a1 * a2
  return multiplicao
}

func main() {
  fmt.Println("resultado:", mult(2, 2))
}
