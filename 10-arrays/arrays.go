package main

import "fmt"

func main() {
    // declare int array with length of 5
    // by default all values will be set to zero
    // 0 is the default value of int
    var a [5]int
    // set values to specific index
    a[0] = 1
    a[4] = 2
    fmt.Println(a)

    // short-hand declaration
    b := [5]int{1, 2}
    fmt.Println(b)

    // compiler can also find the length of an array
    // using [...]
    c := [...]int{2, 4, 6}
    fmt.Println(c)

    d := [...]int{} // declare blank array
    // same with: var d []int
    fmt.Println(d)

    // NOTE: size of the array is part
    // of the data type:
    x := [3]int{10, 20, 30}
    var y [5]int
    // y = x is not valid because [3]int != [5]int
    fmt.Println(x)
    fmt.Println(y)

    // NOTE: arrays are value types, not reference types

    // getting length of array using len()
    fmt.Println("length of x is", len(x))

    // iterate over the array
    fmt.Println("type 1 loop")
    for i := 0; i < len(x); i++ {
        fmt.Println(i, "=>", x[i])
    }

    // another way of iterating using for..range
    fmt.Println("type 2 loop")
    for index, value := range x {
        fmt.Println(index, "=>", value)
    }

    // index can be ignored:
    for _, v := range x {
        fmt.Println(v)
    }

    // value can be ignored also
    for i, _ := range x {
        fmt.Println(i)
    }

    // multidimension array:
    mdArray := [2][2]int {
        {1, 2},
        {3, 4},
    }
    mdArray[0][0] = 2
    mdArray[1][0] = 4
    fmt.Println(mdArray)

    // mixing types is also possible by:
    var mixed [3]interface{}
    mixed[0] = 12
    mixed[1] = "Hello"
    fmt.Println(mixed)
}


