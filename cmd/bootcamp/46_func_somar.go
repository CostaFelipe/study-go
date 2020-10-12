package main

import "fmt"


func somar(a1, a2 int) int {
  add := a1 + a2
  return add
}

func main() {
  fmt.Println("soma:", somar(1, 1))
}
