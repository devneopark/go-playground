package main

import "fmt"

func Multiple(a int, b int) int {
    return a * b
}

func Fibonacci(index int) int {
    if index < 2 {
        return index
    }
    return Fibonacci(index - 2) + Fibonacci(index - 1)
}

func main() {
    fmt.Println(Multiple(3, 3))
    fmt.Println(Fibonacci(9))
}
