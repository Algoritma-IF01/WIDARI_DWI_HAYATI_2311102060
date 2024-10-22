package main

import (
    "fmt"
    "math"
)

// Prosedur untuk menghitung skor peserta
func hitungSkor(waktu []int) (soal, skor int) {
    for _, t := range waktu {
        if t <= 300 { // Hanya menghitung soal yang diselesaikan dalam waktu <= 300 menit
            soal++
            skor += t
        }
    }
    return soal, skor
}

func main() {
    var n int
    fmt.Print("Masukkan jumlah peserta: ")
    fmt.Scan(&n)

    pemenang := ""
    maxSoal := 0
    minSkor := math.MaxInt64

    for i := 0; i < n; i++ {
        var nama string
        var waktu [8]int

        fmt.Printf("Masukkan nama dan waktu penyelesaian soal peserta %d:\n", i+1)
        fmt.Scan(&nama)
        for j := 0; j < 8; j++ {
            fmt.Scan(&waktu[j])
        }

        soal, skor := hitungSkor(waktu[:])
        if soal > maxSoal || (soal == maxSoal && skor < minSkor) {
            pemenang = nama
            maxSoal = soal
            minSkor = skor
        }
    }

    fmt.Printf("Pemenang: %s, Jumlah Soal: %d, Nilai: %d\n", pemenang, maxSoal, minSkor)
}
