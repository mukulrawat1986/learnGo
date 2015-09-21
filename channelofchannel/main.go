package main

import (
    "fmt"
    "time"
)

// We can use a channel of channel to send to the receive the channel on which the data will be sent
// We are also using a timeout over here, so the data emission will stop after 3 seconds
func emit(chanChannel chan chan string, doneCh chan bool){

    // Create a channel of string and send it to our channel of channel
    wordChannel := make(chan string)
    chanChannel <- wordChannel

    defer close(wordChannel)

    words := []string{"the", "quick", "brown", "fox"}

    i := 0

    t := time.NewTimer(1 * time.Second)

    for {
        select {
        case wordChannel <- words[i]:
            i += 1
            if i == len(words){
                i = 0
            }

        case <- doneCh:
            doneCh <- true
            return

        case <- t.C:
            return
        }
    }
}

func main() {

    channelCh := make(chan chan string)
    doneCh := make(chan bool)

    go emit(channelCh, doneCh)

    wordCh := <- channelCh

    for word := range wordCh {
        fmt.Printf("%s\n", word)
    }
}