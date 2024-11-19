package main

import (
    "fmt"
    "math"
)

type ArrayData struct {
    data []int
}

func (a *ArrayData) isiArray(n int) {
    a.data = make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Printf("Masukkan nilai untuk elemen ke-%d: ", i)
        fmt.Scan(&a.data[i])
    }
}

func (a *ArrayData) tampilkanSeluruhArray() {
    fmt.Println("Isi seluruh array:", a.data)
}

func (a *ArrayData) tampilkanIndeksGanjil() {
    fmt.Println("Elemen-elemen dengan indeks ganjil:")
    for i := 1; i < len(a.data); i += 2 {
        fmt.Printf("%d ", a.data[i])
    }
    fmt.Println()
}

func (a *ArrayData) tampilkanIndeksGenap() {
    fmt.Println("Elemen-elemen dengan indeks genap:")
    for i := 0; i < len(a.data); i += 2 {
        fmt.Printf("%d ", a.data[i])
    }
    fmt.Println()
}

func (a *ArrayData) tampilkanKelipatanX(x int) {
    fmt.Printf("Elemen-elemen dengan indeks kelipatan %d:\n", x)
    for i := 0; i < len(a.data); i++ {
        if i%x == 0 {
            fmt.Printf("%d ", a.data[i])
        }
    }
    fmt.Println()
}

func (a *ArrayData) hapusElemen(indeks int) {
    if 0 <= indeks && indeks < len(a.data) {
        a.data = append(a.data[:indeks], a.data[indeks+1:]...)
    }
    fmt.Println("Isi array setelah penghapusan:", a.data)
}

func (a *ArrayData) rataRata() float64 {
    sum := 0
    for _, value := range a.data {
        sum += value
    }
    rata := float64(sum) / float64(len(a.data))
    fmt.Printf("Rata-rata: %.2f\n", rata)
    return rata
}

func (a *ArrayData) standarDeviasi(rata float64) {
    varianceSum := 0.0
    for _, value := range a.data {
        varianceSum += math.Pow(float64(value)-rata, 2)
    }
    variance := varianceSum / float64(len(a.data))
    stdDev := math.Sqrt(variance)
    fmt.Printf("Standar deviasi: %.2f\n", stdDev)
}

func (a *ArrayData) frekuensiBilangan(bilangan int) {
    freq := 0
    for _, value := range a.data {
        if value == bilangan {
            freq++
        }
    }
    fmt.Printf("Frekuensi %d dalam array: %d\n", bilangan, freq)
}

func main() {
    var arrayData ArrayData
    var n, x, indeksHapus, bilanganCari int

    fmt.Print("Masukkan jumlah elemen dalam array: ")
    fmt.Scan(&n)

    arrayData.isiArray(n)

    arrayData.tampilkanSeluruhArray()
    arrayData.tampilkanIndeksGanjil()
    arrayData.tampilkanIndeksGenap()

    fmt.Print("Masukkan bilangan untuk kelipatan x: ")
    fmt.Scan(&x)
    arrayData.tampilkanKelipatanX(x)

    fmt.Print("Masukkan indeks untuk dihapus: ")
    fmt.Scan(&indeksHapus)
    arrayData.hapusElemen(indeksHapus)

    rata := arrayData.rataRata()
    arrayData.standarDeviasi(rata)

    fmt.Print("Masukkan bilangan untuk mencari frekuensinya: ")
    fmt.Scan(&bilanganCari)
    arrayData.frekuensiBilangan(bilanganCari)
}
