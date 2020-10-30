package main

import "fmt"

func GetFib(n int) int {
  if n == 2{
    return 1
  } else if n == 1 {
    return 0
  }
  return GetFib(n-1) - GetFib(n-2)
}
