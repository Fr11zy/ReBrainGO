package main

import (
    "fmt"
)

func main() {
    biblio := map[string]map[string]int{
        "Chel": {
            "books": 4,
            "period": 3},
        "Chel2": {
            "books": 7,
            "period": 1},
        "Chel3": {
            "books": 6,
            "period": 0}}
    
    kol := 0
    summ_period := 0
    for chelik := range biblio{
        if biblio[chelik]["period"]>0 {
            kol+=1
        }
        summ_period += biblio[chelik]["period"]
    }
    fmt.Println("Chelik with period: ",kol)
    fmt.Println("All periods: ", summ_period)
 
}