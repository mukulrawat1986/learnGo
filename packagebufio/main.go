package main

import (
    "fmt"
    "github.com/mukulrawat1986/learnGo/poetry"
)

func main() {

    p, err := poetry.LoadPoem("wordsworth.txt")
    if err != nil{
        fmt.Printf("Error loading poem: %s\n", err)
    }
    fmt.Printf("%s\n", p)
}
