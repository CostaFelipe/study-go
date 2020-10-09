package main

import "fmt"

func main() {
  var a [5][2]string
  for i := 0; i < 5; i++ {
    for j := 0; j < 2; j++ {
      a[i][j] = fmt.Sprintf("Linha %d - Coluna %d", i+1, j+1)
    }
  }

  fmt.Printf("%q", a)
}
