package main

import (
	"fmt"
)

func main() {
	days := []string{"Mon","Tue","Wed","Thur","Fri","Sat","Sun"}
	workdays := make([]string,5)
	
	copy(workdays,days)
	days = days[5:]

	var all_days []string
	all_days = append(all_days, workdays...)
	all_days = append(all_days, days...)

	fmt.Println(all_days)
}