package main

import "fmt"

func Array() {
  a := [2]string{"Ruth", "Linda"}
  fmt.Printf("%v", a)
}

func Slices() {
  a := []int{1, 2, 3, 4, 5, 6}
  fmt.Println(a)
}

func Array_01(a []string) []string{
    return a
}

func main()  {
  Array()
  Slices()
  a := []string{"Felipe", "Ruth Costa"}
  Array_01(a)
  fmt.Println(a)
}
