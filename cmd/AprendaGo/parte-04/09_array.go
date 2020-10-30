package main

import "fmt"

func main() {
    array1 := []int{10, 20, 30}
    array2 := []float64{3.50, 2.25, 1.00}
    array3 := []string{"Linda", "Esther"}

    array := []interface{}{array1, array2, array3}

    fmt.Println(array)
}
