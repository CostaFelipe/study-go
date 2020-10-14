package main

import "fmt"

var lista [10]string


func main() {
  lista[0]     = "Ruth"
  lista[1]     = "Felipe"
  lista[2]     = "Levi"

  fmt.Println(lista[0], lista[1], lista[2])
}
