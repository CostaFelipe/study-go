package main

import (
  "fmt"
  "runtime"
)

func main()  {
    fmt.Println(runtime.GOOS, runtime.GOARCH)

    s := "100"
    for _, v := range s {
      fmt.Println(v)
    }

    a := [5]int{1, 2, 3, 4 , 5}
    for i := 0; i < len(a); i++ {
      fmt.Printf("%v,", i)
    }
}
