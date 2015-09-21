package main

import (
    "fmt"
)

// A map has an index which does not have to be a number, it can be any arbitrary
// type. It maps that to a particular value.

func main(){

    dayMonths := make(map[string]int)

    dayMonths["Jan"] = 31
    dayMonths["Feb"] = 28
    dayMonths["Mar"] = 31
    dayMonths["Apr"] = 30
    dayMonths["May"] = 31
    dayMonths["Jun"] = 30
    dayMonths["Jul"] = 31
    dayMonths["Aug"] = 31
    dayMonths["Sep"] = 30
    dayMonths["Oct"] = 31
    dayMonths["Nov"] = 30
    dayMonths["Dec"] = 31

    fmt.Printf("Days in february: %d\n", dayMonths["Feb"])

    // If you access the element of a map which does not exist you get the zero type
    // for whatever the value type of map is. In this case its integer so we will 
    // get a 0

    // To differentiate between the zero value and an error, we have the comma-ok method
    days, ok := dayMonths["January"]

    if !ok {
        fmt.Printf("Can't get the days for January\n")
    }else{
        fmt.Printf("No. of days: %d\n", days)
    }

    // iteration over the map
    // Here the key and values will be printed in no particular order
    // this way is just to access all elements of a map
    // To get the key, values in a particular order we need to sort it.
    for month, days := range dayMonths {
        fmt.Printf("%s has %d days\n", month, days)
    }

    // Delete an item from the map
    // If we accidentally delete an item twice, Go won't complain
    delete(dayMonths, "Feb")

    // static allocation of map
    temp := map[string]int {
        "a":1,
        "b":2,
    }

    for key, value := range temp {
        fmt.Printf("%s, %d\n", key, value)
    }
}