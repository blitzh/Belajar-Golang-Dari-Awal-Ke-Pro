
package main

import "fmt"

func main() {
    m := map[string]int{"a": 1}
    val, ok := m["b"]
    if ok {
        fmt.Println("b exists:", val)
    } else {
        fmt.Println("b not found")
    }
}
