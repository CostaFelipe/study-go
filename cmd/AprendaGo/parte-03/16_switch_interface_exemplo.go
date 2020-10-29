package main

import "fmt"

var x interface{}

func main() {
  x = true

  switch x.(type) {
      case bool:
        fmt.Println("é tipo booleano")
      case int:
        fmt.Println("é típo inteiro")
      case string:
        fmt.Println("é típo string")
      case float64:
        fmt.Println("é tipo float")
      default:
        fmt.Println("não tem tipo")  
  }

}
