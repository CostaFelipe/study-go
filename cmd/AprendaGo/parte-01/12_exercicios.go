package main

import "fmt"

var x int
var y string
var z bool

func main()  {
    c := float64(x)
    fmt.Println(x, "\n", y, "\n", z)
    fmt.Printf("\n%T, %T", c, x)
}
