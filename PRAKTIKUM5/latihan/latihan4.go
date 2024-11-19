package main

import (
    "fmt"
)

const NMAX int = 127
type tabel [NMAX]rune

// Prosedur untuk mengisi array dengan karakter dari masukan pengguna
func isiArray(t *tabel, n *int) {
    *n = 0
    var char rune
    fmt.Println("Masukkan karakter (akhiri dengan '.'): ")
    for *n < NMAX {
        fmt.Scanf("%c", &char)
        if char == '.' {
            break
        }
        t[*n] = char
        *n++
    }
}

// Prosedur untuk mencetak array karakter
func cetakArray(t tabel, n int) {
    for i := 0; i < n; i++ {
        fmt.Printf("%c", t[i])
    }
    fmt.Println()
}

// Prosedur untuk membalikkan urutan isi array
func balikanArray(t *tabel, n int) {
    for i := 0; i < n/2; i++ {
        t[i], t[n-1-i] = t[n-1-i], t[i]
    }
}

// Fungsi untuk memeriksa apakah array membentuk palindrom
func palindrom(t tabel, n int) bool {
    for i := 0; i < n/2; i++ {
        if t[i] != t[n-1-i] {
            return false
        }
    }
    return true
}

func main() {
    var tab tabel
    var m int
    
    // Isi array tab dengan memanggil prosedur isiArray
    isiArray(&tab, &m)
    
    // Cetak isi array tab sebelum dibalik
    fmt.Println("Isi array sebelum dibalik:")
    cetakArray(tab, m)
    
    // Balikan isi array tab dengan memanggil balikanArray
    balikanArray(&tab, m)
    
    // Cetak isi array tab setelah dibalik
    fmt.Println("Isi array setelah dibalik:")
    cetakArray(tab, m)
    
    // Periksa apakah array membentuk palindrom
    if palindrom(tab, m) {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }
}
