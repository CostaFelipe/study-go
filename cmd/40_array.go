package main

import "fmt"

func main() {
  var a [2][2]string
  for i := 0; i < 2; i++ {
    for j := 0; j < 2; j++ {
      a[i][j] = fmt.Sprintf("linha %d - coluna %d", i+1, j+1)
    }
  }
  fmt.Printf("%q", a)
}
