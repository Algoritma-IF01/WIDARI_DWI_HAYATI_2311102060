package main

import (
    "fmt"
)

// Konstanta dan tipe data
const nMax int = 7919

type Buku struct {
    id, judul, penulis, penerbit string
    eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

// Fungsi untuk mendeklarasikan prosedur Daftarkan Buku
func DaftarkanBuku(pustaka *DaftarBuku, n int) {
    for i := 0; i < n; i++ {
        var buku Buku
        fmt.Scan(&buku.id, &buku.judul, &buku.penulis, &buku.penerbit, &buku.eksemplar, &buku.tahun, &buku.rating)
        pustaka[i] = buku
    }
}

// Fungsi untuk mencetak buku terfavorit
func CetakTerfavorit(pustaka *DaftarBuku, n int) {
    if n == 0 {
        fmt.Println("Tidak ada buku dalam pustaka.")
        return
    }
    terfavorit := pustaka[0]
    for i := 1; i < n; i++ {
        if pustaka[i].rating > terfavorit.rating {
            terfavorit = pustaka[i]
        }
    }
    fmt.Printf("Terfavorit: %s, %s, %s, %d\n", terfavorit.judul, terfavorit.penulis, terfavorit.penerbit, terfavorit.tahun)
}

// Fungsi untuk mengurutkan buku menggunakan insertion sort
func UrutBuku(pustaka *DaftarBuku, n int) {
    for i := 1; i < n; i++ {
        key := pustaka[i]
        j := i - 1
        for j >= 0 && pustaka[j].rating < key.rating {
            pustaka[j+1] = pustaka[j]
            j = j - 1
        }
        pustaka[j+1] = key
    }
}

// Fungsi untuk mencetak 5 buku dengan rating tertinggi
func Cetak5Terbaru(pustaka *DaftarBuku, n int) {
    limit := 5
    if n < 5 {
        limit = n
    }
    fmt.Println("5 Buku dengan rating tertinggi:")
    for i := 0; i < limit; i++ {
        fmt.Printf("%s\n", pustaka[i].judul)
    }
}

// Fungsi untuk mencari buku dengan rating tertentu menggunakan pencarian biner
func CariBuku(pustaka *DaftarBuku, n int, r int) {
    left, right := 0, n-1
    for left <= right {
        mid := (left + right) / 2
        if pustaka[mid].rating == r {
            fmt.Printf("Buku ditemukan: %s, %s, %s, %d, %d, %d\n", pustaka[mid].judul, pustaka[mid].penulis, pustaka[mid].penerbit, pustaka[mid].tahun, pustaka[mid].eksemplar, pustaka[mid].rating)
            return
        } else if pustaka[mid].rating < r {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    fmt.Println("Tidak ada buku dengan rating seperti itu.")
}

func main() {
    var n int
    fmt.Scan(&n)

    var pustaka DaftarBuku
    DaftarkanBuku(&pustaka, n)

    CetakTerfavorit(&pustaka, n)
    UrutBuku(&pustaka, n)
    Cetak5Terbaru(&pustaka, n)

    var rating int
    fmt.Scan(&rating)
    CariBuku(&pustaka, n, rating)
}
