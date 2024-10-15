package main

import "fmt"

func hitungAkar2(K int) float64 {
    result := 1.0
    for k := 0; k <= K; k++ {
        numerator := (4*float64(k) + 2) * (4*float64(k) + 2)
        denominator := (4*float64(k) + 1) * (4*float64(k) + 3)
        result *= numerator / denominator
    }
    return result
}

func main() {
    var K int
    fmt.Print("Nilai K = ")
    fmt.Scan(&K)
    
    result := hitungAkar2(K)
    fmt.Printf("Nilai akar 2 = %.10f\n", result)
}
