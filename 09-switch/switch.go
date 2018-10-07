package main

import "fmt"

func main() {
    num := 10
    switch num {
        case 10:
            fmt.Println("Ten")
            fallthrough // continue to next case
        case 20:
            fmt.Println("Twenty")
        case 30:
            fmt.Println("Thirty")
        // default statement is optional
        default:
            fmt.Println("Default")
    }

    // expression-less switch is possible
    switch {
        case num == 10:
            fmt.Println("yup, it works")
    }
}

