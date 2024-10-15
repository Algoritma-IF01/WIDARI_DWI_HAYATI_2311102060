package main

import "fmt"

func main() {
    var beratKiri, beratKanan float64
    var totalBerat float64

    for {
        fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
        fmt.Scan(&beratKiri, &beratKanan)
        
        if beratKiri < 0 || beratKanan < 0 {
            fmt.Println("Proses selesai.")
            break
        }

        totalBerat = beratKiri + beratKanan
        if totalBerat > 150 {
            fmt.Println("Proses selesai.")
            break
        }

        selisih := beratKiri - beratKanan
        if selisih < 0 {
            selisih = -selisih
        }

        oleng := selisih >= 9
        fmt.Println("Sepeda motor pak Andi akan oleng:", oleng)
    }
}
