package main

import (
    "fmt"
)

// Fungsi f(x) = x^2
func f(x int) int {
    return x * x
}

// Fungsi g(x) = x - 2
func g(x int) int {
    return x - 2
}

// Fungsi h(x) = x + 1
func h(x int) int {
    return x + 1
}

// Fungsi komposisi fogoh(x) = f(g(h(x)))
func fogoh(x int) int {
    return f(g(h(x)))
}

// Fungsi komposisi gohof(x) = g(h(f(x)))
func gohof(x int) int {
    return g(h(f(x)))
}

// Fungsi komposisi hofog(x) = h(f(g(x)))
func hofog(x int) int {
    return h(f(g(x)))
}

func main() {
    var a, b, c int

    fmt.Print("Masukkan tiga bilangan bulat (a b c): ")
    fmt.Scanf("%d %d %d", &a, &b, &c)

    fmt.Printf("(fogoh)(%d) = %d\n", a, fogoh(a))
    fmt.Printf("(gohof)(%d) = %d\n", b, gohof(b))
    fmt.Printf("(hofog)(%d) = %d\n", c, hofog(c))
}
