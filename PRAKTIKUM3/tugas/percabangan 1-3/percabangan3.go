package main

import "fmt"

func main() {
    var b int
    fmt.Print("Bilangan: ")
    fmt.Scan(&b)

    if b <= 0 {
        fmt.Println("Bilangan harus lebih besar dari 0")
        return
    }

    fmt.Print("Faktor: ")
    count := 0
    for i := 1; i <= b; i++ {
        if b % i == 0 {
            fmt.Print(i, " ")
            count++
        }
    }
    fmt.Println()
    
    if count == 2 {
        fmt.Println("Prima: true")
    } else {
        fmt.Println("Prima: false")
    }
}
