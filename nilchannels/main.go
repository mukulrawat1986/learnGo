package main

import (
    "fmt"
    "time"
    "math/rand"
)

// A nil channel will not block when you try to recive from it
// In the context of select statement, once we set the channel to nil, the channel gets
// blocked in context of select statement, so that clause will not execute after it.
func reader(ch chan int){

    t := time.NewTimer(10 * time.Second)

    for {
        select{
        case i := <- ch:
            fmt.Printf("%d\n", i)

        case <- t.C:
            ch = nil
        }
    }
}

func writer(ch chan int){

    stopper := time.NewTimer(2 * time.Second)
    restarter := time.NewTimer(5 * time.Second)

    // save the channel to restart it later
    savedCh := ch

    for {
        select{
        case ch <- rand.Intn(42):
        case <- stopper.C:
            ch = nil
        case <-restarter.C:
            ch = savedCh
        }
    }
}


func main(){

    ch := make(chan int)

    go reader(ch)
    go writer(ch)
    
    time.Sleep(10 * time.Second)
}