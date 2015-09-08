package main

import(
    "fmt",
    "os"
)

// The Printf functions returns number of bytes printed and/or the error
// encountered.

func main() {
    // Here we declare two variables inline, and then check the second variable for a value
    // If no error occurs, the error value returned is nil
    // The variables declared inline are limited to the if scope, so they can't be accessed
    // outside of the if-else statement
    if numberBytes, theError := fmt.Printf("Hello, World!\n"); theError != nil{
        os.Exit(1)
    } else {
        fmt.Prinf("Printed %d bytes\n", numberBytes)        
    }
}