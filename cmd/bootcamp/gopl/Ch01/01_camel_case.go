package main

import "fmt"

func main()  {
  nomes := `
    userID not userId
    productAPI not productApi
  `

  funcoes := `
  // Do is an exported method and is accessible in other packages
    func (a Awesomeness) Do() string {
        return a.doMagic("Awesome")
    }

    // doMagic is where magic happens and only visible inside awesome
    func (a Awesomeness) doMagic(input string) string {
        return input
    }

    QuoteRuneToASCII or parseRequestLine
  `
  fmt.Println(funcoes, "\n", nomes)

}
