package main

import (
    "fmt"
    "math/rand"
)

// Interfaces can be used like types, but they define behaviors rather than 
// specific types
// Any type that has the following functions can be shuffled
type shuffler interface{
    Len() int
    Swap(i, j int)
}

// Anything that satisfies the shuffler interface can be passed to this function
func shuffle(s shuffler){
    for i:=0; i<s.Len(); i++{
        j := rand.Intn(s.Len() - i)
        s.Swap(i, j)
    }
}

// We will make this type shufflable by defining the interface shuffler on it
type intSlice []int

func (is intSlice) Len() int{
    return len(is)
}

func (is intSlice) Swap(i, j int){
    is[i], is[j] = is[j], is[i]
}

type stringSlice []string

func (s stringSlice) Len() int{
    return len(s)
}

func (s stringSlice) Swap(i, j int){
    s[i], s[j] = s[j], s[i]
}

func main(){
    is := intSlice{1,2,3,4,5,6}

    shuffle(is)
    fmt.Printf("%d\n", is)

    s := stringSlice{"The", "quick", "brown", "fox"}

    shuffle(s)
    fmt.Printf("%v\n", s)
}
