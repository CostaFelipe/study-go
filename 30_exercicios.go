package main

import "fmt"

type User struct {
     Nome   string
     Email  string
}

type Contatos struct {
     User
     ContatoId int
}

func (u *User) login() string{
        return fmt.Sprintf("Nome %s & email %s", u.Nome, u.Email)
}

func main() {
	p := Contatos{User{Nome: "Linda", Email: "linda@gmail.com"}, 11111}
	fmt.Println(p.login())
}
