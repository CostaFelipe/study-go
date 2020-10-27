package main

import(
  "fmt"
  "crypto/sha256"
)

func main()  {
    c1 := sha256.Sum256([]byte("felipe"))
    //c2 := sha256.Sum256([]byte("felipe"))

    fmt.Println(c1)
}
