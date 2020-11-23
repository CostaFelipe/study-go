package main

import "fmt"

var x = 2

func main() {
  qualquercoisa(x)
}

func qualquercoisa(y int) {
  fmt.Println(y)
  fmt.Printf("%v, %T", y, y)
}
