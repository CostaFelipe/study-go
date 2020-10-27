package main

import "fmt"

func main()  {
    slice := []string{"Felipe", "Esther", "Linda", "Ruth", "Levi", "Paulo"}

    for posicao, nome := range slice {
      fmt.Println(posicao, "-" , nome)
    }

    fmt.Println("\n")

    slice = append(slice, "Luiz")
    for posicao, nome := range slice {
      fmt.Println(posicao, "-" , nome)
    }
}
