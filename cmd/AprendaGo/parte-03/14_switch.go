package main

import "fmt"

func main() {
  nomeDela := "Ruth"

  switch nomeDela {
      case "Esther":
        fmt.Println("Nome biblico")
      case "Ruth":
        fmt.Println("Outro nome biblico")
      case "Linda":
        fmt.Println("Nome de origem germânica")
      default:
        fmt.Println("Não existe significado")
  }

}
