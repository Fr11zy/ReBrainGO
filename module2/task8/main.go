package main

import (
	"bufio"
	"log"
	"os"
	"fmt"
)

func main() {
	f,err := os.Open("08_task_in.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func(){
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(f)
	totalrow := 0

	for scanner.Scan() {
		totalrow+=1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total strings: %d\n",totalrow)
}