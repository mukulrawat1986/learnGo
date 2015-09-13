// This code is problametic. I have introduce by running multiple goroutines sending on same channel.
// Thus raising the problem of either sending on a closed channel or closing on a closed channel.

package main

import (
    "fmt"
_    "time"
)

// Goroutines are functions that run in parallel to the main program, and we can communicate 
// with these goroutines

//  A channel is a go's fundamental communication mechanism. Used to send messages between goroutines
func emit(c chan string){
    words := []string{"The", "quick", "brown", "fox"}

    for _, word := range words {
        // we are sending the word to the channel
        // time.Sleep(10*time.Millisecond)
        c <- word
      }

    // close the channel
    // defer close(c)

    // fmt.Printf("Closed\n")
}


func main(){

    wordChannel := make(chan string)

    // start a go routine
    for i:=0; i<5; i++ {
        go emit(wordChannel)
    }

    // for word := range wordChannel {
        // fmt.Printf("%s ", word)
    // }

    count := 0
    for {
        if count == 20 {
            break
        }
        word, ok := <- wordChannel
        if ok{
            fmt.Printf("%s\n", word)
            count += 1
        }else{
            fmt.Printf("Channel closed\n")
            break
        }
    }


    fmt.Printf("\n")

}