package main

import (
    "fmt"
)


// The global const and var can be used in any function of the package
// Since the global const and var declared here start with small letters
// they can't be used outside of this package. If there names began with a 
// capital letter, they could be used outside of this package.

// declaring a global const
// We do not use := over here because we are declaring it by
// using const here
const(
    message_const = "The answer to life is %d\n"
    answer_const = 42  
)

// declaring a global variable here
// Again we do not use := over here because we are declaring it by
// using var here
var(
    message_var = "The answer to life is %d\n"
    answer_var = 42
)

// Example of items that are automatically created
const(
    message1_const = "%d %d\n"
    answer1_const = iota * 2 // the value of iota is 1
    answer2_const // This automatically increments the previous value
)

func main(){
    // declaring a variable string
    var message string
    message = "Hello, World!\n"

    // short way to declaring and setting a variable
    message1 := "Hello, Everybody!\n"

    // Using format specifiers in Printf
    message2 := "The answer to life is %d\n"
    answer := 42

    // change the global variable
    answer_var += 1

    fmt.Printf(message)
    
    fmt.Printf(message1)

    fmt.Printf(message2, answer)

    fmt.Printf(message_const, answer_const)

    fmt.Printf(message_var, answer_var)

    fmt.Printf(message1_const, answer1_const, answer2_const)

    // Floating point numbers
    var pi float64 = 3.14
    fmt.Printf("Value: %.2f\n", pi)

    // Another way of declaring the variable
    pi1 := float64(3.14)

    fmt.Printf("Value: %f\n", pi1)

    nine := uint8(9)
    fmt.Printf("Value: %d\n", nine)

    // Boolean variables
    // In go when we don't give a variable an initial value it gets the zero value
    // by default.
    isTrue := !true

    fmt.Printf("Value %t\n", isTrue)

    // Bytes
    b := byte(65)

    // printing out the hex value of the byte
    fmt.Printf("Value: %x\n", b)


}