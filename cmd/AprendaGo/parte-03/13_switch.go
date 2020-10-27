package main

import "fmt"

func main() {
  x := 5
  switch {
      case x < 10:
        fmt.Println("Numero menor que dez")
      case x == 10:
        fmt.Println("Numero igual a dez")
      case x > 10:
        fmt.Println("Numero maior que dez")
  }
}
