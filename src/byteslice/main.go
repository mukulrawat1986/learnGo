package main

import (
    "fmt"
    "os"
)

// Byte slice is a very common slice in Golang
// A lot of i/o function uses byte slices when they read data from files
// or somewhere else

func main(){

    // Open a file
    f, err := os.Open("test.txt")

    if err != nil {
        // err is not of string type, but it gets converted into a string
        // by the Printf function
        fmt.Printf("%s\n", err)
        os.Exit(1)
    }
    defer f.Close()

    // When we read from a file, it reads into a byte slice so we create a byte slice
    b := make([]byte, 100)

    n, err := f.Read(b) 

    fmt.Printf("%d %c\n", n, b)

    //  converting the byte slice into a string
    stringVersion := string(b)

    fmt.Printf("%d %s\n", n, stringVersion)
}