
package main

import (
    "fmt"
)

func tambah(a, b int, ch chan int) {
    ch <- a + b
}

func kali(a, b int, ch chan int) {
    ch <- a * b
}

func bagi(a, b int, ch chan float64) {
    if b != 0 {
        ch <- float64(a) / float64(b)
    } else {
        ch <- 0.0
    }
}

func main() {
    var a, b int
    fmt.Print("Masukkan angka pertama: ")
    fmt.Scan(&a)
    fmt.Print("Masukkan angka kedua: ")
    fmt.Scan(&b)

    chInt := make(chan int)
    chFloat := make(chan float64)

    go tambah(a, b, chInt)
    go kali(a, b, chInt)
    go bagi(a, b, chFloat)

    fmt.Println("Hasil Penjumlahan:", <-chInt)
    fmt.Println("Hasil Perkalian:", <-chInt)
    fmt.Println("Hasil Pembagian:", <-chFloat)
}
