package main

import "fmt"

type User struct {
  Id              int
  Name, Location  string
}

type Player struct {
  User
  GameId int
}

func main()  {
  p := Player{}
  p.Id = 32
  p.Name = "Felipe"
  p.Location = "FOR"
  p.GameId = 90904
  fmt.Println("%+v", p)
}
