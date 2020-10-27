package main

import "fmt"

func main() {

  x := 4

  switch {
    case (x == 4), (x == 2):
      fmt.Println("Numero um ou dois")
    case (x == 10), (x == 5):
      fmt.Println("numero dez ou cinco")
  }

}
