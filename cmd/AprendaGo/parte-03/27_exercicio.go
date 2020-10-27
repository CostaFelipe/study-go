package main

import "fmt"

func main()  {

    esporteFavorito := "futebol"

    switch esporteFavorito {
    case "futebol":
      fmt.Println(esporteFavorito)
    case "natacao":
      fmt.Println(esporteFavorito)
    case "beisebol":
      fmt.Println(esporteFavorito)
    default:
      fmt.Println("Esse esporte não está na lista")
    }
    
}
