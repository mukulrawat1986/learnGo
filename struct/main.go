package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

// We can build our own type in Go using the type keyword, we can also attach
// functions called methods to a type
// Type can be structs or aliases for built in types
type webPage struct{
    url string
    body []byte
    err error
}


// the expression written before the get() function is called the receiver
// It defines the type on which we are operating
func (w *webPage) get(){
    resp, err := http.Get(w.url)
    if err != nil{
        return
    }
    defer resp.Body.Close()

    w.body, err = ioutil.ReadAll(resp.Body)
    if err != nil{
        w.err = err
    }
}

func (w *webPage) isOK() bool{
    return w.err == nil
}

type summable []int

func (s summable) sum() int{
    sum := 0
    for _, r := range s{
        sum += r
    }

    return sum
}

func main(){

    w := &webPage{url: "http://www.oreilly.com"}

    w.get()
    if w.isOK() {
        fmt.Printf("URL: %s length %d\n", w.url, len(w.body))
    }else{
        fmt.Printf("Something went wrong\n")
    }

    var s summable = summable{1, 1, 2, 3, 5, 8, 13}

    fmt.Printf("%d\n", s.sum())
} 