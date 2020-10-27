package main

import "fmt"

func main()  {

    x := 3

    if (x == 2 || x == 3) {
      fmt.Println("Deu certo")
    }

    y := 6
    if (y % 2 == 0 && y % 3 == 0) {
      fmt.Println("o numero", y, "é multiplo de 2 e 3")
    }
}
