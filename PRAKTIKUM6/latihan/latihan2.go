package main

import (
    "fmt"
)

type Ikan struct {
    Berat float64
}

type Wadah struct {
    Ikan   []Ikan
    Total  float64
}

func main() {
    var x, y int
    fmt.Scan(&x, &y)
	fmt.Print("Masukkan berat ikan dalam kg: ")
    ikan := make([]Ikan, x)
    for i := 0; i < x; i++ {
        fmt.Scan(&ikan[i].Berat)
    }

    wadahCount := (x + y - 1) / y
    wadah := make([]Wadah, wadahCount)
    for i := 0; i < x; i++ {
        index := i / y
        wadah[index].Ikan = append(wadah[index].Ikan, ikan[i])
        wadah[index].Total += ikan[i].Berat
    }

    for _, w := range wadah {
        fmt.Printf("%.2f ", w.Total)
    }
    fmt.Println()

    var totalBerat float64
    for _, w := range wadah {
        totalBerat += w.Total
    }

    beratRataRata := totalBerat / float64(wadahCount)
    fmt.Printf("%.2f\n", beratRataRata)
}
