package main

import "fmt"

type User struct {
	Nome      string
	IdCode    int
}

type Gamer struct {
	User
	Posicao   string
	IdGamer   int
}

func (u *User) Login() string {
	return fmt.Sprintf("%s", u.Nome)
}

func main(){
	p := Gamer{User{Nome: "Robinho", IdCode: 90904},"Meio", 456}
	fmt.Println(p.Login())
}
