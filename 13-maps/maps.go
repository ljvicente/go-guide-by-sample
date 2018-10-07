package main

import "fmt"

func main() {
    // map is a built-in type which
    // associates a value to a key

    // to create a map:
    salaries := make(map[string]int)

    // adding elements:
    salaries["joe"] = 1000
    salaries["john"] = 2000
    fmt.Println(salaries)

    // maps can be initialized 
    // during the declaration
    ages := map[string]int{
        "joe": 39,
        "alan": 50,
    }
    ages["mike"] = 22
    fmt.Println(ages)

    // if the element is not present,
    // the default value of the type will be used.
    // in this case, 0
    fmt.Println("Age of Ted is", ages["ted"])

    // to check whether the element is existing,
    // use the following syntax:
    value, ok := ages["mike"]
    fmt.Println("Mike is existing:", ok, "Age of:", value)

    // for..range loops is used to iterate
    // over all the elements of a map
    for key, value := range ages {
        fmt.Println(key, ":", value)
    }

    // to delete element
    delete(ages, "mike")

    // to get the length of the map, use len()
    fmt.Println("length is", len(ages))

    // maps are reference types like slice
    years := ages
    years["joe"] = 69
    fmt.Println(ages)

    // maps cannot be compared by ==
    // if ages == years {}
}
