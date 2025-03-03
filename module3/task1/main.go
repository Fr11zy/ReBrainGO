package main

import (
	"fmt"
	"task1/packages/pack"
	"github.com/huandu/xstrings"
)

func main() {
	fmt.Println(xstrings.Shuffle(pack.City()))
	fmt.Println(pack.Digit())
}