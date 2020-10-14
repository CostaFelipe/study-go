package main

import "fmt"

func main()  {
  s := "Ruth, Levi"
  for i := 0; i < len(s); i++ {
    fmt.Println(s[i])
  }

  fmt.Println(" ")

  for _, v := range s {
    fmt.Printf("%b\n", v)
  }
}
