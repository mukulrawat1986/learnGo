package main

import (
    "fmt"
    "sync/atomic"
    "time"
    "math/rand"
)

var(
    running int64 = 0
)

func work(){
    atomic.AddInt64(&running, 1)
    // fmt.Printf("[")
    fmt.Printf("Starting: %d\n", running)
    time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
    // fmt.Printf("]")
    // atomic.AddInt64 (&running, -1)
    fmt.Printf("Closing: %d\n", running)
}

func worker(sema chan bool){

    // we will only do something when we recive something on this channel
    <- sema

    work()

    // once we are done working we will send true to the channel 
    sema <- true    
}

func main(){

    // While declaring a buffered channel we set up a capacity for the channel
    // This channel can store up to 10 ints, once it gets full, anyone sending on this channel 
    // will have to wait for there to be space. Anyone reciving on this channel can recive ints which
    // are buffered, even if nobody is transmitting at that time
    sema := make(chan bool, 10)

    for i:=0; i<1000; i++ {
        go worker(sema)
    }

    for i:=0; i < cap(sema);i++{
        // We can put 10 values here, the channel won't block since its buffered
        sema <- true
    }

    time.Sleep(30 * time.Second)

}