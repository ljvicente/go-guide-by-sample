package main

import "fmt"

func main () {
    var age int // defaults to '0' basing on the type
    age = 69
    fmt.Println("my age is", age)

    var name = "leo" // type can be ommited if initial value is set 
    fmt.Println("my name is", name)
    
    var height, weight int = 156, 70 // multiple declaration
    fmt.Println("my height is", height, "my weight is", weight)

    // multiple declaration with multiple types
    var (
        teacher = "lexi"
        teacherAge = 25
    )
    fmt.Println("my teacher's name is", teacher, "her age is", teacherAge)

    // multiple declaration shorthand
    a, b := 1, 2
    fmt.Println("a =", a, "b =", b)
    // NOTE: `:=` can be only used if there's at least one
    // on the left side is newly declared.
}
