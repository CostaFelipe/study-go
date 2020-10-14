package main

import "fmt"

func main()  {
  s := "Hello Esther"

  for _, v := range s {
    fmt.Printf("%v - %#U - %#x\n", v, v, v)
  }

  fmt.Println(" ")

  for i := 0; i < len(s); i++ {
    fmt.Printf("%v - %T - %#U - %#x\n", s[i], s[i], s[i], s[i])
  }
}
