package main

import (
    "fmt"
)

func main() {

// infinite loop
//   for {
//    fmt.Printf("Hello World\n")
//   } 

    //  counter is initialized to zero if we don't initialize it with
    //  a value
    var counter int

    for counter < 10 {
        fmt.Printf("Hello, World!\n")
        counter += 1
    }

    // A compact way of using for loops
    for i := 0; i < 10; i++ {
        fmt.Printf("Hello, Everybody!\n")
    }

    // Another example, but initializing multiple variables this time
    for i, j := 0, 1; i < 10; i, j = i+1, j*2 {
        fmt.Printf("%d Hello, World!\n", j)
    }

    var stop bool
    i := 0

    for !stop {
        fmt.Printf("Won't stop the loop!\n")

        i += 1
        if i == 10 {
            stop = true
        }
    }
}