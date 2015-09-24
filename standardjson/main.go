package main

import (
    "encoding/json"
_    "fmt"
    "log"
    "net/http"
    "os"
    "github.com/mukulrawat1986/learnGo/poetry"
)

// We can also use the module to read JSON
// Here we will create a struct type to read JSON
type config struct{
    Route string
    BindAddress string `json:"addr"`
}


// We can also encode structures using encoding/json
// Here we create an example struct to encode
// It still follows the rule, that is a structure member is private it won't show up on Response
type poemWithTitle struct{
    Title string
    Body  poetry.Poem
}

func poemHandler(w http.ResponseWriter, r *http.Request){
    err := r.ParseForm()
    if err != nil{
        log.Println("ParseForm Error", err)
    }

    poemName := r.Form["name"][0]
    p, err := poetry.LoadPoem(poemName)
    if err != nil{
        log.Println("LoadPoem Error", err)
    }else{
        pwt := poemWithTitle{poemName, p}
        enc := json.NewEncoder(w)
        enc.Encode(pwt)
    }
}

func main(){
    f, err := os.Open("config")
    if err != nil{
        os.Exit(1)
    }
    defer f.Close()
    // Create an object for our struct
    var c config
    dec := json.NewDecoder(f)
    err = dec.Decode(&c)
    if err != nil{
        os.Exit(1)
    }

    log.Println("Server starting")
    http.HandleFunc(c.Route, poemHandler)
    err = http.ListenAndServe(c.BindAddress, nil)
    if err != nil{
        log.Fatal("ListenAndServe Error", err)
    }
}
