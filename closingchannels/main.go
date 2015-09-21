package main

import (
    "fmt"
    "time"
)

// We can use close on a channel to cause the work to happen or not happen
// We will use close on channel to coordinate goroutines to start operating at the same time

// When we have multiple goroutines listening on the same channel, the data that we send on the channel will be
// received by only one goroutine.
// We can make all goroutines receive something by using close


func printer(msg string, goCh chan bool){
    <- goCh 

    fmt.Printf("%s\n", msg)
}


func printer1(msg string, stopCh chan bool){
    for {
        select{
        case <- stopCh:
            return

        default:
            fmt.Printf("%s\n", msg)
        }
    }
}

func main(){

    goCh := make(chan bool)
    stopCh := make(chan bool)

    for i:=0; i<10;i++{
        go printer(fmt.Sprintf("printer %d", i), goCh)
    }

    time.Sleep(5 * time.Second)
    close(goCh)
    time.Sleep(5 * time.Second)


    for i := 0; i<10; i++{
        go printer1(fmt.Sprintf("printer %d", i), stopCh)
    }

    time.Sleep(5 * time.Second)
    close(stopCh)
    time.Sleep(5 * time.Second)
}