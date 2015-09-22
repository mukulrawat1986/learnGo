package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

// We can define our own type using the keyword type, we can also attach functions to 
// this type, called methods. This is kind of similar to object oriented programming

// we define our own type to receive a webpage
// This consists of a url, body and error type
type webPage struct{
    url string
    body []byte
    err error
}

// We create a function that acts on the type webpage
// The (w *webPage) part is called the receiver and it means the function is
// operating on that type. Since we are using pointer over here, it means we can 
// make changes to our struct in the function.
func (w *webPage) get(){
    resp, err := http.Get(w.url)
    if err != nil{
        w.err = err
    }

    // close the body of response when we are done with it
    defer resp.Body.Close()

    w.body, err  = ioutil.ReadAll(resp.Body)
    if err != nil{
        w.err = err
    }
}

// We create another method on our webPage type 
func (w *webPage) isOk() bool{
    return w.err == nil
}

// We can also create aliases using type
type SummableSlice []int

// We create a method with our new type as a receiver to it. Here we are passing by value, so
// any changes made to our type in the method will not be reflected outside of the method
func (s SummableSlice) sum() int{
    sum := 0
    for _, i := range s{
        sum += i
    }
    return sum
}

func main(){
    // one way to create the object to our struct is
    w := &webPage{url: "http://www.oreilly.com"}
    w.get()
    if w.isOk() {
        fmt.Printf("URL: %s Length: %d\n", w.url, len(w.body))
    }else{
        fmt.Printf("Error: %s\n", w.err)
    }

    // Now we create our summable type
    var s SummableSlice = SummableSlice{1,1,2,3,5,8,13}

    fmt.Printf("The sum is %d\n", s.sum())
}

