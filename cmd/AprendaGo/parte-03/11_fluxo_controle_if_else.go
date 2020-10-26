package main

import "fmt"

func main()  {
  if x := 1000; x < 100 {
    fmt.Println("X não é maior que 100")
  } else {
    fmt.Println("X é maior que 100")
  }

  fmt.Println("\n")

  if x := 1000; (x > 100) {
    fmt.Println("X é maior que cem")
  } else if (x < 10) {
    fmt.Println("X é menor que dez")
  } else {
    fmt.Println("x não é menor que dez e maior que cem")
  }
}
