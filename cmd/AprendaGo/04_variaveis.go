package main

import "fmt"

var y = 2

func main() {
    qualquercoisa(y)
}

func qualquercoisa(x int) {
  fmt.Println(y)
  fmt.Printf("%v, %T",x, x)
}
