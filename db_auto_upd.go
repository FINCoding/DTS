package main

import (
    "fmt"
    "os"
    "io"
)

func main() {
    file, err := os.Open("./data_from_parsing_files/lada.txt")
    print_err(err)


    defer file.Close()

    data := make([]byte, 64)

    for{
        n, err := file.Read(data)
        if err == io.EOF{
            break
        }
        fmt.Println(string(data[:n]))
    }
}

func print_err(err error) {
    if err != nil{
        fmt.Println(err)
        os.Exit(1)
    }
}