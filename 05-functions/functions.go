package main

import (
    "fmt"
)

func main() {
    fmt.Println("sum of 1 and 2 is", sumNumbers(1, 2))
    
    a, b := addOne(3, 5)
    fmt.Println("a =", a, "b =", b)

    x, _ := addTwo(3, 5) // discarding values is also possible
    fmt.Println("x =", x)
}

// you can omit the type on the parameter if consecutive
// parameters have the same type..
func sumNumbers(a, b int) int {
    return a + b
}

// you can also return multiple values
func addOne(a, b int) (int, int) {
    return a + 1, b + 1
} 

// named return values
func addTwo(a, b int) (x, y int) {
    x = a + 2
    y = b + 2
    return
}
