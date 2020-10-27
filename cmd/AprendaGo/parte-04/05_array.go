package main

import "fmt"

var a [5]int = [5]int{1, 2, 7, 10, 5}

func main() {
    fmt.Println(a[0])
    fmt.Println(a[len(a)-1])

    fmt.Println("\n")

    for i, valor := range a {
      fmt.Printf("%d - %d\n", i, valor)
    }

    for _, v := range a {
      fmt.Printf("%d", v)
    }
}
