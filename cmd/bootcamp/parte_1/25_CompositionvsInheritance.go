package main

import "fmt"

type User struct {
  Id int
  Name, Email, Senha string
}

type Conta struct {
  User
  TypeConta int
  IdConta   int
}

func main() {
  p := Conta{
    User{Id: 21, Name: "Felipe Costa", Email: "costadefelipe@gmail.com", Senha: "asdf145236"},
    1, 90904,
    }

    fmt.Println(p, p.login())
}

func (u *User) login() string {
  return fmt.Sprintf("email is %s and senha is %s", u.Email, u.Senha)
}
