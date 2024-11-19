package main

import (
    "fmt"
)

// Struct untuk menyimpan informasi pertandingan
type Pertandingan struct {
    klubA   string
    klubB   string
    skorA   int
    skorB   int
    hasil   []string
}

// Method untuk mengisi data pertandingan
func (p *Pertandingan) isiData() {
    fmt.Print("Klub A: ")
    fmt.Scan(&p.klubA)

    fmt.Print("Klub B: ")
    fmt.Scan(&p.klubB)

    for i := 1; ; i++ {
        fmt.Printf("Pertandingan %d: ", i)
        fmt.Scan(&p.skorA, &p.skorB)

        // Berhenti jika salah satu skor tidak valid (negatif)
        if p.skorA < 0 || p.skorB < 0 {
            break
        }

        // Menentukan hasil pertandingan
        var hasilPertandingan string
        if p.skorA > p.skorB {
            hasilPertandingan = p.klubA
        } else if p.skorB > p.skorA {
            hasilPertandingan = p.klubB
        } else {
            hasilPertandingan = "Draw"
        }

        p.hasil = append(p.hasil, hasilPertandingan)
    }
}

// Method untuk menampilkan pemenang
func (p *Pertandingan) tampilkanHasil() {
    if len(p.hasil) > 0 {
        for i, hasil := range p.hasil {
            fmt.Printf("Hasil %d: %s\n", i+1, hasil)
        }
    } else {
        fmt.Println("Tidak ada hasil pertandingan yang valid.")
    }
    fmt.Println("Pertandingan selesai")
}

func main() {
    var pertandingan Pertandingan

    pertandingan.isiData()
    pertandingan.tampilkanHasil()
}
