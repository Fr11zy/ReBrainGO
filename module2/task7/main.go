package main

import (
	"bufio"
	"fmt"
	"log"
    "os"
    "strings"
)

func main() {
    f,err := os.Open("07_task_in.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    output,err := os.Create("data/data_out.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer output.Close()

    scanner := bufio.NewScanner(f)
    writer := bufio.NewWriter(output)
    defer writer.Flush()

    defer func(){
        if r := recover(); r != nil {
            fmt.Println("Error was: ", r)
            
            data,err := os.ReadFile("data/data_out.txt")
            if err != nil {
                log.Fatal(err)
            }

            fmt.Println("Содержимое data_out.txt до ошибки:")
            fmt.Println(string(data))
        }
    }()

    row := 1
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line,"|")

        if len(parts)<3 || parts[0] == "" || parts[1] == "" || parts[2] == "" {
            panic(fmt.Sprintf("parse error: empty field on string %d",row))
        }
        
        finalLine := fmt.Sprintf("Row: %d\nName: %s\nAddress: %s\nCity : %s\n\n\n",
            row, parts[0], parts[1], parts[2])
        
        _, err := writer.WriteString(finalLine)
        if err != nil {
            log.Fatal(err)
        }

        row+=1
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
