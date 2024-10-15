package main

import "fmt"

func main() {
    var beratParsel int
    fmt.Print("Berat parsel (gram): ")
    fmt.Scan(&beratParsel)

    totalKg := beratParsel / 1000
    sisaGram := beratParsel % 1000

    biayaTotal := totalKg * 10000
    biayaSisa := 0

    if totalKg > 10 {
        biayaSisa = 0
    } else {
        if sisaGram >= 500 {
            biayaSisa = sisaGram * 5
        } else {
            biayaSisa = sisaGram * 15
        }
    }

    totalBiaya := biayaTotal + biayaSisa

    fmt.Printf("Detail berat: %d kg + %d gr\n", totalKg, sisaGram)
    fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaTotal, biayaSisa)
    fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
