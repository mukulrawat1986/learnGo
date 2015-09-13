package main

import (
    "errors"
    "fmt"
)

var (
    erroEmptyString = errors.New("Unwilling to print an empty string")
)


// Golang is built around handling errors and not dealing with exceptions
// Here we will create custom error using fmt package
// This method is useful if we just want to return a string as an error
func printer(msg string) error {

    if msg == "" {
        return fmt.Errorf("Unwilling to print an empty string")
    }
    _, err := fmt.Printf("%s\n", msg)

    return err
}


// The errors package allows you to define an error and then you can do equality on it
func printer1(msg string) error {

    if msg == "" {
        return erroEmptyString
    }

    _, err := fmt.Printf("%s\n", msg)

    return err
}


func main() {

    if err := printer1(""); err != nil {
        if err == erroEmptyString{
            fmt.Printf("You tried to print an empty string!\n")
        }else{
            fmt.Printf("printer failed: %s\n", err)             
        }
    }
}