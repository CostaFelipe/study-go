package main

import "fmt"

type User struct {
  Id              int
  Name, Location  string
}

func (u *User) Greeting() string {
  return fmt.Sprintf("Hi %s from %s",u.Name, u.Location)
}

type Player struct {
  User
  GameId int
}

func main() {
  p := Player{}
  p.Id = 32
  p.Name = "Felipe Costa"
  p.Location = "FOR"
  fmt.Println(p.Greeting())
}
