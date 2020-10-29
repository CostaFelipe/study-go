package main

import "fmt"

const boilingF = 212.0

func main() {
    f := boilingF
    c := (f - 32) * 5 / 9
    fmt.Printf("Ponto de ebulição = %g°F ou %g°C\n", f, c)
}
