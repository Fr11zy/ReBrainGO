package main

import (
	"fmt"
)

type AmericanVelocity float64
type EuropeanVelocity float64

func main() {	
	var speed1 EuropeanVelocity = 120.4*3.6
	var speed2 AmericanVelocity = 130*3600/1609	
	fmt.Println(speed1)
	fmt.Println(speed2)
}