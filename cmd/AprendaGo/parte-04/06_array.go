package main

import "fmt"

type Currency int

const (
  USD Currency = iota
  EUR
  GBP
  RMB
)

func main()  {
  symbol := [...]string{USD: "$", EUR: "9", GBP: "!", RMB: ""}
  fmt.Println(USD, symbol[USD]) // "3 ""
}
