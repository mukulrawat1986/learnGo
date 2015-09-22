package main

import (
    "fmt"
)

// Sometimes in Go its useful to use an empty interface.
// That's an interface with no methods mentioned in it, and
// its satisfied by everything.
// Its helpful in having an unknown type that is determined at
// runtime.

func whatisThis(i interface{}){
    fmt.Printf("%T\n", i)

    switch v := i.(type ) {
    
    case string:
        fmt.Printf("It's a string %s \n", v)

    case uint32:
        fmt.Printf("Its an unsigned 32-bit integer %d \n", v)

    default:
        fmt.Printf("Don't know %v\n", v)

    }
}


func main(){
    whatisThis(uint32(43))

    whatisThis("Hello")

    whatisThis([...]string{"A", "B", "C"})

}