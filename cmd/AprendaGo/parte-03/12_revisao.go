package main

import "fmt"

func main()  {
  s := "Ruth"
  sb := []byte(s)

  for _, v:= range sb {
    fmt.Printf(string(v))
  }
}
