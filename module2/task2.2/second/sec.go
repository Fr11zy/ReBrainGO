package main

import (
	"fmt"
	"math"
)

func main() {
	var l float64 = 35
	var R *float64
	radius := l/2/3.14
	R = &radius

	S := 3.14*(*R)*(*R)
	fmt.Println(math.Round(S*100)/100)
}