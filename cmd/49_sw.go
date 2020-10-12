package main

import "fmt"

func calc(a1, a2 float64, unit ... string) float64 {
  switch unit[0] {
  case "ADD":
    return a1 + a2
  case "SUB":
    return a1 + a2
  case "MULT":
    return a1 * a2
  case "DIV":
    return a1 / a2
  default:
    return 0
  }
}

func main() {
  fmt.Println("valor:", calc(1,2, "DIV"))
}
