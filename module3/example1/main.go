package main

import (
	"fmt"
	"github.com/fatih/color"
	"example1/wordz"
)

func main() {
	fmt.Println("Hello World")
	color.Red("Hello world yoy")
	fmt.Println(wordz.Random())
}