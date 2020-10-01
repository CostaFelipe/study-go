package main

import "fmt"

type User struct {
  Id              int
  Name, Location  string
}

type Player struct {
  *User
  GameId          int
}

func main(){
  p := newPlayer(42, "Cuesta", "FOR", 90904)
  fmt.Println(p.Greeting())
}
// --------------- func ---------------------------------
func (u *User) Greeting() string {
  return fmt.Sprintf("Hi %s from %s", u.Name, u.Location)
}

func newPlayer(id int, name, location string, gameId int) *Player {
  return &Player {
    User: &User{id, name, location},
    GameId: gameId,
  }
}
