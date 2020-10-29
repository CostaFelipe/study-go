package main

import "fmt"

var matriz [3][2]string

func main()  {
    for i := 0; i < 3; i++ {
      for j := 0; j < 2; j++ {
        matriz[i][j] = fmt.Sprintf("[%d][%d]", i + 1, j + 1)
      }
    }
    fmt.Printf("%q", matriz)
}
