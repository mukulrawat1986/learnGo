package main

import (
    "fmt"
    "math/rand"
)

// Interfaces can be used a little bit like type. But they define behavior 
// rather than individual type

// we create an interface here
// Anything that satisfies this interface or has these methods defined on it
// is something that can be shuffled
type shuffler interface {
    Len() int
    Swap(i, j int)
}

// now we will define a function that takes the interface.
// so anything that satisfies the interface can be passed to the function
func shuffle (s shuffler){
    for i := 0; i < s.Len(); i++{
        j := rand.Intn(s.Len() - i)
        s.Swap(i,j)
    }
}

// Now we will create a type that satisfies the interface
type intSlice []int

// And here we define the methods associated with this type

func (i intSlice) Len() int{
    return len(i)
}

func (is intSlice) Swap(i, j int) {
    is[i], is[j] = is[j], is[i]
}

// Now we will create a completely different type that satisfied the interface shuffler too
type stringSlice []string

// Again we create methods with this type to satisfy the interface
func (s stringSlice) Len() int{
    return len(s)
}

func (s stringSlice) Swap(i, j int){
    s[i], s[j] = s[j], s[i]
}

func main(){

    is := intSlice{1,2,3,4,5,6}

    shuffle(is)

    fmt.Printf("%v\n", is)

    s := stringSlice{"The", "quick", "brown", "fox"}

    shuffle(s)
    
    fmt.Printf("%v\n", s)

}
