package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    // Strings can be initialize inside " "
    message := "Hello World"
    fmt.Println(message)

    // a string is a slice of bytes
    fmt.Printf("%x\n", message[0]) // output will be in UTF-8 encoded
    fmt.Printf("%c\n", message[0]) // output will be in normal

    message = "Ola Señor" // 'ñ' occupies 2 bytes
    printBytes(message)
    printChars(message)
    printCharsUsingRune(message)

    fmt.Println(len("Señor")) // returns 6 because 'ñ' is 2 bytes
    // to get the correct length of string, use utf.RuneCountInString()
    fmt.Println(utf8.RuneCountInString("Señor")) // returns 5

    // NOTE: strings are immutable once created
    // To workaround, convert string to slice of runes
    m := []rune(message)
    m[0] = 'A'
    fmt.Println(string(m)) // cast rune back to string

    // construct string using slice of bytes
    byteSlice := []byte{65, 66, 67} // code point equivalent of ABC in decimal 
    fmt.Println(string(byteSlice))

    runeSlice := []rune{0x0041, 0x0042, 0x0043} // hex equivalent of ABC
    fmt.Println(string(runeSlice))
}

func printBytes(s string) {
    for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
    fmt.Printf("\n")
}

func printChars(s string) {
    for i := 0; i < len(s); i++ {
        fmt.Printf("%c ", s[i])
    }
    fmt.Printf("\n")
}

func printCharsUsingRune(s string) {
    runes := []rune(s)
    for i := 0; i < len(runes); i++ {
        fmt.Printf("%c ", runes[i])
    }
    fmt.Printf("\n")

    // other way to iterate
    for _, rune := range s {
        fmt.Printf("%c ", rune)
    }
    fmt.Printf("\n")
}
