package main

import (
  "fmt"
)

/*
Inside a function, the := short assignment statement can be used in
place of a var declaration with implicit type.
*/
func main()  {
  name, location := "Princess Yaruka", "São Paulo"
  idade := 32
  fmt.Printf("%s (%d) of %s",name, location, idade)
}
