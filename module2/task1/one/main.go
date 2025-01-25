package main

import (
	"fmt"
	"strconv"
)

func main() {
	str1 := "104"
	str2 := 35
	newstr1,_ := strconv.Atoi(str1)
	newstr2 := strconv.Itoa(str2)
	fmt.Println(newstr1, newstr2)
}