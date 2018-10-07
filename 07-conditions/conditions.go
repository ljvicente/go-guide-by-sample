package main

import "fmt"

func main() {
    x := 1
    if x == 1 {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }

    // another variant of above condition 
    if y := 1; y == 1 {
        fmt.Println("true")
    } else { // else must be on the same line as the followed closing, otherwise, go will cry
        fmt.Println("false")
    }
}

