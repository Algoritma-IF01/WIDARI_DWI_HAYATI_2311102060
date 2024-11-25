package main

import (
    "fmt"
)

func selectionSort(arr []int, ascending bool) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        minOrMaxIndex := i
        for j := i + 1; j < n; j++ {
            if (ascending && arr[j] < arr[minOrMaxIndex]) || (!ascending && arr[j] > arr[minOrMaxIndex]) {
                minOrMaxIndex = j
            }
        }
        arr[i], arr[minOrMaxIndex] = arr[minOrMaxIndex], arr[i]
    }
}

func main() {
    var n int
    fmt.Scan(&n)
    
    for i := 0; i < n; i++ {
        var m int
        fmt.Scan(&m)
        
        oddNumbers := []int{}
        evenNumbers := []int{}
        
        for j := 0; j < m; j++ {
            var num int
            fmt.Scan(&num)
            if num%2 == 1 {
                oddNumbers = append(oddNumbers, num)
            } else {
                evenNumbers = append(evenNumbers, num)
            }
        }
        
        // Sort odd numbers in ascending order
        selectionSort(oddNumbers, true)
        
        // Sort even numbers in descending order
        selectionSort(evenNumbers, false)
        
        // Print the sorted odd numbers
        for _, num := range oddNumbers {
            fmt.Printf("%d ", num)
        }
        
        // Print the sorted even numbers
        for _, num := range evenNumbers {
            fmt.Printf("%d ", num)
        }
        
        fmt.Println()
    }
}
