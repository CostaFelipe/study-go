package main

import "fmt"

const(
  _     = iota
  KB    = 1 << (iota * 10)
  MB
  GB
  TB
)

func main()  {
    fmt.Printf("%b - %d\n", KB, KB)
    fmt.Printf("%b - %d\n", MB, MB)
    fmt.Printf("%b - %d\n", GB, GB)
    fmt.Printf("%b - %d\n", TB, TB)
}
