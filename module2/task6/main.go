package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("06_task_in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	outfile, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		i += 1
		line := scanner.Text()
		outfile.WriteString(strconv.Itoa(i) + ") " + line + "\n")
	}
}
