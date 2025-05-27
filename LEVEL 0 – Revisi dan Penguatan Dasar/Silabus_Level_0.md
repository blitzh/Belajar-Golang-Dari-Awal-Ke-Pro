
# Golang Level 0 - Revisi dan Penguatan Dasar

## Tujuan Utama
1. Memahami cara kerja pointer, slice, map, dan struct secara memory.
2. Menulis kode idiomatik Golang yang clean dan readable.
3. Mengerti concurrency dasar dengan goroutine dan channel.
4. Menghindari bug umum dalam pengembangan dengan Go.

---

## Materi & Contoh

### 1. Pointer dan Reference
```go
func update(val int) {
    val = 100
}

func updatePtr(val *int) {
    *val = 100
}

func main() {
    num := 10
    update(num)         // Tidak berpengaruh
    updatePtr(&num)     // Mengubah nilai
    fmt.Println(num)    // Output: 100
}
```

### 2. Slice dan Copy Semantik
```go
func modify(s []int) {
    s[0] = 99
}

func main() {
    data := []int{1, 2, 3}
    modify(data)
    fmt.Println(data) // Output: [99 2 3]
}
```

### 3. Map
```go
m := make(map[string]int)
m["a"] = 1
fmt.Println(m["b"]) // Output: 0 (tidak error)
```

### 4. Error Handling
```go
f, err := os.Open("file.txt")
if err != nil {
    log.Fatalf("error opening file: %v", err)
}
defer f.Close()
```

### 5. Interface & Type Assertion
```go
func describe(i interface{}) {
    fmt.Printf("(%v, %T)
", i, i)
}
```

### 6. Goroutine & Channel
```go
func say(msg string) {
    for i := 0; i < 3; i++ {
        fmt.Println(msg)
    }
}

func main() {
    go say("Hello")
    say("World")
}
```

---

## Latihan Mini Project

### CLI Kalkulator Paralel
- Input 2 angka dari CLI
- Jalankan 3 goroutine: tambah, kali, bagi
- Kirim hasil ke channel
- Cetak semua hasil setelah selesai

---

## Evaluasi
Checklist keterampilan Level 0 untuk memastikan pemahaman yang solid.
