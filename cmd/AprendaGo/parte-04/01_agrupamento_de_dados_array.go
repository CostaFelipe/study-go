package main

import "fmt"

var x [5]int

func main()  {

    x[0] = 1
    x[1] = 2
    x[2] = 3
    x[3] = 4
    x[4] = 5

    fmt.Println(x)
    fmt.Println("O tamanho do array é:", len(x))
}
