package main

import (
    "fmt"
)

// We may have multiple channels in our goroutine, Select, is a way of chosing a channel out of 
// multiple channels
// Channels are bidirectional so we can send on a channel we received on
func emit(wordCh chan string, doneCh chan bool){

    words := []string{"the", "quick", "brown", "fox"}

    i := 0

    for {
        select {
        case wordCh <- words[i]:
            i += 1
            if i == len(words) {
                i = 0
            }

        case <- doneCh:
            fmt.Printf("Got a done message\n")
            close(doneCh)
            return
        }
    }
}


func main() {

    wordCh := make(chan string)
    doneCh := make(chan bool)

    go emit(wordCh, doneCh)

    for i:=0; i<100; i++ {
        fmt.Printf("%s\n", <- wordCh)
    }

    doneCh <- true
}