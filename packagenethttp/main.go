package main

import (
    "fmt"
    "github.com/mukulrawat1986/learnGo/poetry"
    "log"
    "net/http"
)


func poemHandler(w http.ResponseWriter, r *http.Request){
    err := r.ParseForm()
    if err != nil{
        log.Println("ParseForm Error", err)
    }
    poemName := r.Form["name"][0]

    p, err := poetry.LoadPoem(poemName)

    if err != nil{
        log.Println("Loading Poem Error", err)
    }else{
        // Since we have defined the Stringer interface on the poem we can just print it out
        fmt.Fprintf(w,"%s", p)
        log.Println("Poem requested and served ", poemName)
    }
}


func main(){

    log.Println("Server started")
    http.HandleFunc("/poem", poemHandler)
    err := http.ListenAndServe(":8080", nil)
    if err != nil{
        log.Fatal("ListenAndServe Error: ", err)
    }
}
