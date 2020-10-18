package main

import "fmt"

func main()  {
  for horas := 0; horas <= 12; horas++ {
    fmt.Print("hora:", horas)
    for i := 0; i < 60; i++ {
      fmt.Print(" ", i)
    }
    fmt.Print("\n")
  }
}
