package main

import "fmt"

type Jogador struct {
	Nome string
	Posicao string
}

func main() {
	j := Jogador{}
	j.Nome = "Neymar"
	j.Posicao = "Atacante"
	fmt.Println(j)
}
