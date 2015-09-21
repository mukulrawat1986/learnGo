package main

import (
    "fmt"
)

// A function that takes a pointer to a string
// any changes we make to the array here will be reflected back in our string in
// the main funciton
func printer(w *[4] string){
    for _, word := range *w {
        fmt.Printf("%s\n", word)
    }
    w[2] ="blue"
}

func main() {

    // An array of string of unknown number of values
    // An array is a fixed piece of memory
    words := [...] string{"the", "quick", "brown", "fox"}

    // send the pointer to the array to the function
    printer(&words)
    printer(&words)

    // Declaring and initializing a pointer to an array
    w := & words

    // Iterating over the pointer of string array
    for _, r := range w {
        fmt.Printf("%s\n", r)
    }

    // declaring and initializing a string and a pointer to it.
    i := "Hello"
    p := &i

    fmt.Printf("%s\n", *p)

    nums := [4] int {1,2,3,4}

    num_p := &nums

    // Iterating over the pointer to the string array and modifying
    // its value
    for val, _ := range *num_p{
        num_p[val] = val
    }

    // Iterating over the pointer to the string array and printing the values out
    for _, num := range *num_p{
        fmt.Printf("%d\n", num)
    }

}