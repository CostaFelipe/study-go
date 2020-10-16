package main

import "fmt"

func main()  {
  var array1 [6]int = [6]int{1, 2, 3, 4, 5, 6}
  fmt.Println(array1)

  var array2 = [6]int{1, 2, 3, 4, 5, 6}
  fmt.Println(array2)

  var array3 = [...]int{1, 2, 3, 4, 5, 6}
  fmt.Println(array3)
}
