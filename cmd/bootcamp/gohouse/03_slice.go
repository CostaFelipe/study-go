package main

import "fmt"

func main() {
  a1 := []string{"Felipe", "Ruth", "Levi"}
  a2 := append(a1, "Esther", "Linda")

  fmt.Println(a1, a2)
}
