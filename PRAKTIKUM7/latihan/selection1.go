package main

import (
    "fmt"
)

func selectionSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        minIndex := i
        for j := i + 1; j < n; j++ {
            if arr[j] < arr[minIndex] {
                minIndex = j
            }
        }
        arr[i], arr[minIndex] = arr[minIndex], arr[i]
    }
}

func main() {
    var n int
    fmt.Scan(&n)
    
    for i := 0; i < n; i++ {
        var m int
        fmt.Scan(&m)
        
        arr := make([]int, m)
        for j := 0; j < m; j++ {
            fmt.Scan(&arr[j])
        }
        
        selectionSort(arr)
        
        for j := 0; j < m; j++ {
            fmt.Printf("%d ", arr[j])
        }
        fmt.Println()
    }
}
