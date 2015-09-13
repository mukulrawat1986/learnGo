// Another try at the problem using select, faulty program

package main

import (
        "fmt"
        _"time"
)

func emit(c chan string, count chan int) {

        words := []string{"The", "quick", "brown", "fox"}

        total := 0
        for {
                fmt.Printf("Printing\n")
                // count <- 0
                select {

                case <-count:
                        total += 1
                        fmt.Printf("total is %d\n", total)
                        if total == 4 {
                                fmt.Printf("Closing channel\n")
                                close(c)
                                return
                        }


                default:
                        for _, word := range words {
                            c <- word
                        }
                        count <- 1
                        // fmt.Printf("Done sending words\n")
                        // time.Sleep(5*time.Second)
                        // return
                        // count <- 1
                }
        }
}

func main() {

        wordChannel := make(chan string)
        stopChannel := make(chan int)

        for i := 0; i < 5; i++ {
                go emit(wordChannel, stopChannel)
        }

        // for word := range wordChannel {
                // fmt.Printf("%s\n", word)
        // }

        for {
                word := <- wordChannel
                fmt.Printf("%s\n", word)
        }

}
