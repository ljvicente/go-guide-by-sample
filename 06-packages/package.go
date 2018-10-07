package main

import (
    "fmt"
    "06-packages/sample"
    "math"
    _ "unsafe" // unused imports are not allowed, but you can silence error..
)

// just another error silencer..
var _ = math.Sqrt

func main() {
    fmt.Println(sample.SayHi())
}


