package main

import "fmt"

func main()  {

    x := 10

    switch {
        case x < 10:
          fmt.Println("menor do que dez")
        case x == 10:
          fmt.Println("igual a dez")
    }
}
