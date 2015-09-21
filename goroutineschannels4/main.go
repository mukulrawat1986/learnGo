package main

import (
    "fmt"
)


func test_ch(ch chan string, timer chan bool, exit chan int){

    words := []string{"the", "quick", "brown", "fox"}

    for {
        select{
        case <-timer :
            for _, word := range words{
                ch <- word
                exit <- 0
            }
            timer = nil
        
        default:
            ch <- ""
            exit <- 1        

        }
            
    }
    
}

func main() {
    ch := make(chan string)
    timer := make(chan bool)
    exit := make(chan int)

    for i := 0; i<5; i++{
        go test_ch(ch, timer, exit)
    }

    close(timer)

    total := 0

    for {
        c := <- ch
        if c != ""{
            fmt.Printf("%s\n", c)
        }
        total += <- exit
        if total == 5{
            break
        }
    }
}