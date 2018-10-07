package sample

import "fmt"

var testInt = 123

func init() {
    // package-level variables are initialized first
    // before init()
    fmt.Println("initialized", testInt)
}

// NOTE: variables/functions that start with caps
// means that they are "exported" names in go
func SayHi() (string) {
    return "Hi"    
}
