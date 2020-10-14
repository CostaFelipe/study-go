package main

import "fmt"

var a [10]int = [10]int{1, 2, 3, 4, 5, 8, 3, 7, 2, 10}

func main()  {
    NumerosIguais()
}

func NumerosIguais() {
    for i, array := range a {
      for j, array2 := range a {
         if array == array2 && j > i {
            fmt.Println(array)
          }
       }
    }
}
