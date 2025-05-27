// Penjelasan struktur folder idiomatik Golang
// cmd/appname/main.go
// internal/module_name
// pkg/shared

package main

import (
	"fmt"
	"myapp/greet"
)

func main() {
	fmt.Println(greet.SayHello("Gopher"))
}
