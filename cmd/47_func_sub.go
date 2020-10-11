package main

import "fmt"

func sub(a1, a2 int) int{
  subtracao := a1 - a2
  return subtracao
}

func main() {
  fmt.Println("valor é:", sub(5, 2))
}
