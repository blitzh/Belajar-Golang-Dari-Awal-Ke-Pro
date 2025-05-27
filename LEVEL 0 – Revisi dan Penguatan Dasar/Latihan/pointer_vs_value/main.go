
package main

import "fmt"

type Data struct {
    Value int
}

func modifyCopy(d Data) {
    d.Value = 100
}

func modifyPointer(d *Data) {
    d.Value = 100
}

func main() {
    d := Data{Value: 10}
    modifyCopy(d)
    fmt.Println("After modifyCopy:", d.Value) // 10

    modifyPointer(&d)
    fmt.Println("After modifyPointer:", d.Value) // 100
}
