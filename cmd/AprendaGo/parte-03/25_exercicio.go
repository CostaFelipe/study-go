package main


import "fmt"

func main()  {
    movimentar := -1

    if movimentar == 1 {
        fmt.Println("mover")
    } else if (movimentar == 0) {
      fmt.Println("não mover")
    } else {
      fmt.Print("Erro de movimento")
    }
    
}
