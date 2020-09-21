package main

import "fmt"

type User struct {
  Id             int
  Name, Location string
}

type Player struct {
  User
  GameId         int
}

func main() {
  p := Player {
    User{Id: 42, Name: "Felipe Costa", Location: "FOR"},
    90904,
  }
  p.Id = 12
  fmt.Println(p)
}
