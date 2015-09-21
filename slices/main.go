package main

import (
    "fmt"
)

func printer(w []string){

    for _, word := range w{
        fmt.Printf("%s", word)
    }
    fmt.Printf("\n")


}


func main(){

    // A slice is kind of a window into an underlying array. 
    // A slice can store the position in the array, and the number of elements in the
    // array

    // declare a slice. This is static declaration
    words := []string{"the", "quick", "brown", "fox", "jumped", "over", "the", "lazy", "dog"}

    fmt.Printf("Length: %d\n", len(words))

    // slices are passed by reference in a function
    // so if you pass a slice to a function, modifying them will change things
    printer(words[0:2])

    // We can declare a slice dynamically by using the keyword make
    // Here we create a slice of size 4. All initial strings are empty 
    words1 := make([]string, 4)
    words1[0] = "The"
    words1[1] = "quick"
    words1[2] = "brown"
    words1[3] = "fox"

    printer(words1)

    // We can also create a slice in the way where we can set an initial capacity 
    // and a max capacity
    // Here the initial size of slice is 0, but we can add 4 strings(max capacity) to it
    words2 := make([]string, 0, 4)
    fmt.Printf("len: %d capacity: %d\n", len(words2), cap(words2))
    words2 = append(words2, "The")
    words2 = append(words2, "quick")
    words2 = append(words2, "brown")
    words2 = append(words2, "fox")
    fmt.Printf("len: %d capacity: %d\n", len(words2), cap(words2))

    printer(words2)
    words2 = append(words2, "Jumps")
    // We see that as we append more to the slice and go beyond the max capacity, the
    // slice makes room for the new string and increases the max capacity
    fmt.Printf("len: %d capacity: %d\n", len(words2), cap(words2))

    // We pass slices through functions when we want to use them as reference
    // But sometimes we want to make a copy of the slice, for that we have the keyword copy
    newWords := make([]string, len(words2))

    // This makes a copy of words2 in newWords
    copy(newWords, words2)

    printer(newWords)

    // If we modify newWords, then the change will not be reflected in words2
    // By using := we will create a new slice which will be referring to the original slice
}