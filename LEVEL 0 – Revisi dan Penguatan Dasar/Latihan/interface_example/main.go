
package main

import "fmt"

type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func describe(s Shape) {
    fmt.Println("Area:", s.Area())
}

func main() {
    c := Circle{Radius: 5}
    describe(c)
}
