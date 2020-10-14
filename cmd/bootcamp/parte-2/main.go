package main

import "fmt"

var array = [4]int{1, 5, 4, 7}


func main()  {
    for _, s := range array {
        fmt.Println(s)
    }
}
