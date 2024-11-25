package main

import (
    "fmt"
)

func insertionSort(arr []int) {
    for i := 1; i < len(arr); i++ {
        key := arr[i]
        j := i - 1
        for j >= 0 && arr[j] > key {
            arr[j+1] = arr[j]
            j = j - 1
        }
        arr[j+1] = key
    }
}

func checkEqualSpacing(arr []int) (bool, int) {
    if len(arr) < 2 {
        return true, 0
    }
    distance := arr[1] - arr[0]
    for i := 2; i < len(arr); i++ {
        if arr[i] - arr[i-1] != distance {
            return false, 0
        }
    }
    return true, distance
}

func main() {
    var num int
    numbers := []int{}
    
    // Reading input numbers
    for {
        fmt.Scan(&num)
        if num < 0 {
            break
        }
        numbers = append(numbers, num)
    }
    
    // Sort the array using insertion sort
    insertionSort(numbers)
    
    // Check if the data is equally spaced
    equalSpacing, distance := checkEqualSpacing(numbers)
    
    // Print the sorted array
    for _, num := range numbers {
        fmt.Printf("%d ", num)
    }
    fmt.Println()
    
    // Print the status of spacing
    if equalSpacing {
        fmt.Printf("Data berjarak %d\n", distance)
    } else {
        fmt.Println("Data berjarak tidak tetap")
    }
}
