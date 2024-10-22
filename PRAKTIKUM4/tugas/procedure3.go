package main

import "fmt"

// Prosedur untuk mencetak deret Skiena
func cetakDeret(n int) {
    for n != 1 {
        fmt.Printf("%d ", n)
        if n%2 == 0 {
            n = n / 2
        } else {
            n = 3*n + 1
        }
    }
    fmt.Println(1)
}

func main() {
    var n int

    fmt.Print("Masukkan satu bilangan bulat positif yang lebih kecil dari 1000000: ")
    fmt.Scan(&n)

    if n > 0 && n < 1000000 {
        cetakDeret(n)
    } else {
        fmt.Println("Nilai tidak valid. Masukkan bilangan antara 1 dan 999999.")
    }
}
