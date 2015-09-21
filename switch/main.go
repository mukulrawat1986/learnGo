package main

import (
    "fmt"
    "os"
)

func main() {
    
    n, err := fmt.Printf("Hello, World!\n")

    // In go we don't need to use break statement since once a case matches, it won't fallthrough
    // We can make it happen manually by adding the fallthrough command at the end of a case

    switch {

    case err != nil:
       os.Exit(1)
   case n == 0:
        fmt.Printf("No bytes output!\n")
    case n != 14:
        fmt.Printf("Wrong number of characters\n")
        fallthrough
    default:
        fmt.Printf("OK!\n")
    }

    atoz := "the quick brown fox jumps over the lazy dog\n"

    // Another example of switch statement
    // Here will be switching over the characters

    vowels := 0
    consonants := 0
    zeds := 0

    for _, r := range atoz {

        switch r {

        // Since there is no fallthrough in go, we can use all the cases in one statement
        // separated by comms
        case 'a', 'e', 'i', 'o', 'u':
            vowels += 1
        case 'z':
            zeds += 1
            fallthrough
        default:
            consonants += 1
        }
    }

    fmt.Printf("Vowels: %d, Consonants: %d (Zeds: %d)\n", vowels, consonants, zeds)
}