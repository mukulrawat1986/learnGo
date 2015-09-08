package main

import (
    "fmt"
)

func main() {

    atoz := "the quick brown fox jumps over the lazy dog\n"

    // making substring of the first 9 characters
    // we don't need to add 0, by default it starts from beginning.
    fmt.Printf("%s\n", atoz[0:9])

    // Going through string roon by roon
    for i, r := range atoz {
        fmt.Printf("%d %c\n", i, r)
    }

    // length of the string
    fmt.Printf("%d\n", len(atoz))

    // In go we can use backquotes, and anything inside backquotes is taken literally
    // as it is.
    atoz1 := `the quick brown fox jumps over the lazy dog\n`

    fmt.Printf("%s\n", atoz1)
}