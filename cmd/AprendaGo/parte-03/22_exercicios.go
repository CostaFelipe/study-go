package main

import "fmt"

func main()  {
  anoquenasci := 1989
  anoqueestou := 2020

  for {
    if  anoquenasci > anoqueestou {
      break
    }
    fmt.Println(anoquenasci)
    anoquenasci++
  }

}
