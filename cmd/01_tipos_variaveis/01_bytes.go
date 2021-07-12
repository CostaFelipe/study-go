package main

import "fmt"

func main() {
	bytes, erros := fmt.Print("ok")
	fmt.Println(bytes, erros)
}
