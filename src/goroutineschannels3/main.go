package main

import (
    "fmt"
)

func makeID(idChan chan int){
    var id int
    id = 0

    for {
        idChan <- id
        id ++
    }
}

func receiveId(idChan chan int){
    x := <- idChan

    fmt.Printf("id is %d\n", x)
}

func main(){

    idChan := make(chan int)

    go makeID(idChan)
    for i:=0; i<5; i++{
        go receiveId(idChan)
    }
    // fmt.Printf("%d\n", <-idChan)
    for {

    }
}