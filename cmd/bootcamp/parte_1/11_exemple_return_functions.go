package main

import "fmt"

func location(city string) (string, string) {
  var region string
  var continent string
  switch city {
  case "New York", "NYC":
    region, continent = "New York", "Noth America"
  case "Fortaleza", "FOR":
    region, continent = "Fortaleza", "Brazil"
  default:
    region, continent = "Unknown", "Unknow"
  }
  return region, continent
}

func main(){
  region, continent := location("Fortaleza")
  fmt.Printf("Felipe lives in %s, %s", region, continent)
}
