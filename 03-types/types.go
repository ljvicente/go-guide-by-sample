package main

import (
    "fmt"
    "unsafe"
)

func main() {
    var a int = 1

    fmt.Printf("type of a is %T\n", a)
    fmt.Printf("size of a is %d\n", unsafe.Sizeof(a))
    // NOTE: size of int depends on the number
    // of bits in a microprocessor

    b := 1.1 // defaults to float64
    fmt.Printf("type of b is %T\n", b)

    var c uint = 244
    var d int = 1
    sum := int(c) + d // go doesn't allow operations on different types. needs to cast first
    fmt.Printf("sum of c and d is %d\n", sum)
}

