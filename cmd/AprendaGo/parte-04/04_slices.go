package main

import "fmt"

func main()  {
    slice := []int{1, 4, 8, 3}
    soma := 0

    for _, valor := range slice {
      soma += valor
    }

    fmt.Println(soma)
}
