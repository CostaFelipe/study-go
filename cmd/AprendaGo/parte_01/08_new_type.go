package main

import "fmt"

type hotdog int
var b hotdog = 10

type palavra string
var nome palavra = "Linda"

func main()  {
  x := 10
  x = int(b)

  fmt.Printf("%v\n", x)
  fmt.Printf("%v, %T\n", b, b)
  fmt.Printf("%v, %T", nome, nome)
}
