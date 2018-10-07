package main

import "fmt"

func main() {
    // Variadic functions are functions that
    // can accept variable number of arguments
    find(10, 10, 20, 30)
    find(5, 10, 20, 30)

    // slice can be also passed to a variadic function
    // instead of passing a list of arguments
    find(5, []int{10, 20, 30}...)

    sampleNumbers := []int{9, 8, 5}
    find(0, sampleNumbers...)

    // variadic functions also recognizes a slice
    // as a passed by reference argument
    modify(sampleNumbers...)
    fmt.Println("new value of sampleNumbers:", sampleNumbers)
}

// checks if a number is existing
// in a list of integers
func find(number int, numbers ...int) {
    fmt.Printf("type of numbers is %T\n", numbers)
    found := false
    for i, v := range numbers {
        if v == number {
            fmt.Println(number, "found at index", i, "in", numbers)
            found = true
        }
    }
    if !found {
        fmt.Println(number, "not found in", numbers)
    }
}

func modify(numbers ...int) {
    numbers[0] = 1
    numbers[1] = 2
    numbers[2] = 3

    // and as usual, appended slice is not visible outside.
    numbers = append(numbers, 4)
}

