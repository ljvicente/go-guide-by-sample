package main

import (
    "fmt"
)

func main() {
    var a = 1 + 2
    fmt.Println("Hoi", a)

    const b = 1 + 2
    fmt.Println("const add:", b)
    // NOTE: const b = math.Sqrt(25) is not valid
    // as the const needs to be assigned before runtime

    var c int = 1
    fmt.Printf("type %T\n", c)

    const d = true // this doesn't have a type yet
    fmt.Printf("type %T\n", d) // the command demands a type, that means const supplies the type of the value by default

    type trueOrFalse bool
    var e trueOrFalse = d
    fmt.Printf("type %T, value %v\n", e, e)
}
