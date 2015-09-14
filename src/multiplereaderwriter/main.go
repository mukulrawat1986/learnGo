package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    _"os"
)


// A function that takes a url and returns the size and/or error
func getPage(url string) (int, error) {

    // use http Get function, returns a response and/or error
    resp, err := http.Get(url)

    if err != nil{
        return 0, err
    }

    // we need to close the response body after using it
    defer resp.Body.Close()

    // use ioutil ReadAll to read the response body, returns a byte slice and/or error
    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        return 0, err
    }

    // return size of body and no error
    return len(body), nil

}

// We will write a wrapper function that uses getPage, and transmits down the channel
func worker(urlChan string chan, size chan string, id int){
    for {
        url := <- urlChan
        length, err := getPage(url)
        if err == nil {
            size <- fmt.Sprintf("%s has length %d (%d)", url, length, id)
        }else {
            size <- fmt.Sprintf("Error getting %s: %s", url, err)
        }
    }
}

func generator(url string, urlChan chan string){
    urlChan <- url
}

func main() {

    urls := []string{"http://www.google.com", "http://yahoo.com", 
        "http://www.bing.com", "http://bbc.co.uk"}

    size := make(chan string)
    urlChan := make(chan string)

    // for _, url := range urls{
    //     pageLength, err := getPage(url)
    //     if err != nil {
    //         os.Exit(1)
    //     }
    // // pageLength, err := getPage(url)

    // // if err != nil {
    //     // os.Exit(1)
    // // }

    // fmt.Printf("%s is length %d\n", url, pageLength)
    // }


    for i := 0; i<10; i++ {
        go worker(urlChan, size, i)
    }



    for i, url := range urls {
        go generator(url, urlChan)
    }

    for i := 0; i < len(urls); i++{
        fmt.Printf("Length: %s\n", <-size)
    }
}