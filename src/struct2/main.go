package main

import (
    "fmt"
    "math"
)

// A struct is a type which contains named fields
// The type keyword introduces a new type.
type Circle struct{
    x, y, r float64    
}

type Rectangle struct{
    x, y float64
}

// Since in functions arguments are always copied in Go
// So we send the pointer to the struct type
func circleArea(c *Circle) float64{
    return math.Pi * c.r * c.r
}

func distance(x1, y1, x2, y2 float64) float64{
    a := x2 - x1
    b := y2 - y1
    return math.Sqrt(a*a + b*b)
}

func rectangleArea(r1 *Rectangle, r2 *Rectangle) float64{

    l := distance(r1.x, r1.y, r1.x, r2.y)
    w := distance(r1.x, r1.y, r2.x, r1.y)
    return l*w
}

func main(){

    // This allocates memory for all the fields, sets each
    // of them to their zero value and returns a pointer
    // (*Circle)
    c := new(Circle)
    

    // Initialize the fields in the struct
    c.x = 0
    c.y = 0
    c.r = 5

    // here we send the address of the struct type
    // since we used new to create the instance of struct circle
    // c contains (*Circle), hence we do not need to send the address
    // just have to send c
    fmt.Println(circleArea(c))

    // We can also create a instance of the struct in this way
    var r1 Rectangle

    // Another way to create a instance of the struct
    r2 := Rectangle{x: 0, y: 10}

    // Initialize the fields in the struct
    r1.x = 0
    r1.y = 10

    fmt.Println(rectangleArea(&r1, &r2))

}