// Creating our own shuffle package

package shuffler

import (
    "math/rand"
)

// set up a counter
var (
    counter int
)

// Initialize the counter
func init(){
    counter = 0
}

// Get the counter value
func GetCount() int{
    return counter
}

type Shufflable interface{
    Len() int
    Swap(i, j int) 
}

type WeightedShufflable interface{
    Shufflable
    Weight(i int) int
}


func Shuffle(s Shufflable) {
    for i:=0; i<s.Len(); i++{
        j := rand.Intn(s.Len() - i)
        s.Swap(i,j)
        counter++
    }
}

func WeightedShuffle(w WeightedShufflable) {
    total := 0
    for i:=0; i < w.Len(); i++{
        total += w.Weight(i)
    }

    for i:=0; i<w.Len(); i++{
        pos := rand.Intn(total)
        cur := 0
        for j:=i; j<w.Len();j++{
            counter++
            cur += w.Weight(j)
            if pos >= cur{
                total -= w.Weight(j)
                w.Swap(i,j)
                break
             }
        }
    }
}
