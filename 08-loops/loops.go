package main

import (
    "fmt"
)

func main() {
    // there's only one type of loop in go, `for`
    for i := 0; i <= 10; i++ {
        if i == 0 {
            continue // skip iteration
        }
        if i == 10 {
            break // break the loop
        }
        fmt.Println(i)
    }

    // multiple variables are possible also
    for j, k := 2, 2; j > 1 && k <= 10; j, k = j + 1, k + 1 {
	fmt.Println("j=", j, "k=", k)
    }

    fmt.Println("CTRL + C to exit")
    for {
	// infinite loop
    }
}

