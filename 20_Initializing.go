package main

import (
  "fmt"
)

type Bootcamp struct {
  Lat, Lon   float64
}

func main() {
    x := new(Bootcamp)
    y := &Bootcamp{
      Lat: 32,
      Lon: 43,
    }

    fmt.Println(*x == *y)
}
