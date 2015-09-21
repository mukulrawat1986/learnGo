package main

import (
    "fmt"
    "os"
)

// Simple function that takes two arguments
// since both arguments are of same type we only declare the type once
func printer(msg, msg2 string, repeat int){

    for repeat > 0 {
        fmt.Printf("%s", msg)
        fmt.Printf("%s\n", msg2)
        repeat -= 1
    }
}

//  A function that returns error
func printer1 (msg string) (string, error){

    // When we give a defer statement it gets queued up to execute when the funciton returns
    // As the function returns with the message and/or the error, it executes the defer statement
    defer fmt.Printf("\n")

    _, err := fmt.Printf(msg)
    msg += "\n"
    return msg, err
}

// Another function to show defer and opening of files
func printer2(msg string) error {

    f, err := os.Create("helloworld.txt")
    
    // use defer to close the file after creating it
    defer f.Close()

    if err != nil{
        return err
    }

    // Write to the file
    // We can't write strings to a file, so we convert it to [] byte
    f.Write([] byte(msg+"\n"))
    return err
}

// Example of named return
// Notice we don't name the error during return
func printer3(msg string) (e error){
    _, e = fmt.Printf("%s\n", msg)
    return
}

// Function to take variable arguments
func printer4(msgs ...string){

    for _, msg := range msgs {
        fmt.Printf("%s\n", msg)
    }
}

// Another example of a function with variable arguments
func printer5(format string, msgs ...string){

    for _, msg := range msgs{
        fmt.Printf(format, msg)
    }
}

func main() {

    printer("Hello", "World", 6)
    appendedMessage, err := printer1("Hello, World!")
    if err == nil {
        fmt.Printf("%q\n", appendedMessage) // %q prints everything in the string without escaping anything
        fmt.Printf("% x\n", appendedMessage) // %x prints it in hex form
    }

    // call printer2
    printer2("Hello World!")

    e := printer3("Hello World!")

    if e != nil {
        os.Exit(1)
    }

    // sending variable arguments
    printer4("hello", "world", "fdfd")

    printer5("% x\n", "hello", "world", "this", "is", "dog")
}