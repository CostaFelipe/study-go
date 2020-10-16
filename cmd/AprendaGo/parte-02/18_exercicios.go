package main

import "fmt"

const b int = 10
const a = 100

func main() {
    fmt.Printf("%v - %T\n", a, a)
    fmt.Printf("%v - %T", b, b)
}
