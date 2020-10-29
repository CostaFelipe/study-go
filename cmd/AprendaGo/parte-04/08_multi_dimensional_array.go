package main

import "fmt"

var a [2][3]string

func main()  {
    for i := 0; i < 2; i++ {
      for j := 0; j < 3; j++ {
        a[i][j] = fmt.Sprintf("row - %d - column %d", i + 1, j + 1)
      }
    }
    fmt.Printf("\n%q", a)
}
