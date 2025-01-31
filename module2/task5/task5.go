package main

import (
	"fmt"
	"math"
)

func main() {
    stringi := []string{"xuy","pip","noway"}
    stroka1 := "noway"
    cond := contains(stringi,stroka1)

    switch cond{
    case true:
        fmt.Println("yes, stroka in stringi")
    default:
        fmt.Println("Anlak")
    }

    fmt.Println("Максимальное целое число - это", getMax(5,2,6,6,7,878787))
}	

func contains(slice []string,x string) bool {
    for _,s := range slice {
        if s==x {
            return true
        }
    }
    return false;
}

func getMax(n ...int) int {
    maximum := math.MinInt
    for _,i := range n {
        if i>maximum{
            maximum=i
        }
    }
    return maximum
}