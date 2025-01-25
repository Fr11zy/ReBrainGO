package main

import (
	"fmt"
)

func main() {
	days := []string{"Mon","Tue","Wed","Thur","Fri","Sat","Sun"}
	work := make([]string,5)
	
	workdays := copy(work,days)
	fmt.Println(workdays)

	days = days[5:]
	fmt.Println(days)
}