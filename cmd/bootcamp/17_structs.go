package main

import (
  "fmt"
  "time"
)

type Bootcamp struct {
  //Latitude
  Lat float64
  //Longitude
  Long float64
  //Date
  Data time.Time
}

type Person struct {
  Nome string
  Idade int16
  Data time.Time
}

func main(){
  fmt.Println(Bootcamp{
    Lat: 34.012836,
    Long: -118.495338,
    Data: time.Now(),
  })

  fmt.Println(Person{
    Nome: "Felipe",
    Idade: 32,
    Data: time.Now(),
  })
}
