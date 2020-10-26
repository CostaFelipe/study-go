package main

import "fmt"

func main()  {

    x := 10
    if x < 100 {
      fmt.Println("Hello, World!")
    }

    //operador de negação " ! "
    if !(x > 100) {
      fmt.Println("Hello")
    }

    if x := 10; (x < 100) {
      fmt.Println("Hello, World!")
    }
}
