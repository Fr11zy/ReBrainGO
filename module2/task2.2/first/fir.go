package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var A *int
	var B int = rand.Intn(10000)
	A = &B
	fmt.Println(*A)
	*A = 5790
	fmt.Println(B)
}