package main

import (
    "fmt"
)

func main() {
    // A pointer stores the memory address of another variable
    b := 255
    var a *int = &b
    fmt.Println(a) // shows address of b
    // dereferencing pointer:
    fmt.Println(*a) // shows value of b

    // change value of b
    *a++
    fmt.Println(b)
    change(a)
    fmt.Println(b)

    var c *int
    fmt.Println(c) // zero value of pointer is nil

    d := [3]int{10, 20, 30}

    modifyArrayByPointer(&d)
    fmt.Println(d)

    // use slice instead of pointer
    modifyArrayBySlice(d[:])
    fmt.Println(d)

    // NOTE: go does not support pointer arithmetic like C++
}

func change(val *int) {
    *val = 55
}

func modifyArrayByPointer(arr *[3]int) {
    (*arr)[0] = 80 // shorthand version is arr[0]
}

func modifyArrayBySlice(arr []int) {
    arr[0] = 90
}

