// We will use our shuffler package over here

package main

import (
    "fmt"
    "github.com/mukulrawat1986/learnGo/shuffler"
)


// We will create our type that satisfies the shuffable interface 
// and use it here
type intSlice []int

func (is intSlice) Len() int{
    return len(is)
}

func (is intSlice) Swap(i, j int){
    is[i], is[j] = is[j], is[i]
}

// Create a new string slice that satisfied weighted shuffling

type weightString struct{
    weight int
    s string
}

type stringSlice []weightString

func (is stringSlice) Len() int{
    return len(is)
}

func (is stringSlice) Swap(i, j int){
    is[i], is[j] = is[j], is[i]
}

func (is stringSlice) Weight(i int) int{
    return is[i].weight
}

func main(){
    i := intSlice{1,2,3,4,5,6,7}

    shuffler.Shuffle(i)
    fmt.Printf("%v\n", i)

    x := stringSlice{weightString{100,"Hello"}, weightString{200,"World"},
    weightString{10,"Goodbye"}}

    shuffler.Shuffle(x)
    fmt.Printf("%v\n",x)

    shuffler.WeightedShuffle(x)
    fmt.Printf("%v\n",x)

    fmt.Printf("Loop: %d\n", shuffler.GetCount())
}
