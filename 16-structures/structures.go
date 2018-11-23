package main

import (
    "fmt"
    "16-structures/records"
)

// Structure is a user defined type
// which represents a collection of fields

// to declare a "named" structure:
type Employee struct {
    name string
    age int
}

func main() {
    var testStruct Employee
    fmt.Println(testStruct) // shows zero value of struct

    // creating structure using field names
    e := Employee {
        name: "Jack",
        age: 33,
    }

    // creating structure without field names
    e2 := Employee {
        "Joe",
        26,
    }

    fmt.Println(e)
    fmt.Println(e2)

    // to declare an "anonymous" structure:
    e3 := struct {
        name string
        age int
    } {
        "John",
        12,
    }

    fmt.Println(e3)

    // accessing struct fields using "."
    fmt.Println(e3.name)

    // values can be also set later
    e3.name = "Logan"
    fmt.Println(e3)

    // anonymous fields is also possible:
    type Address struct {
        string
        int
    }
    var addr Address
    fmt.Println(addr)

    // by default, name of anonymous field
    // is the name of its type
    addr.string = "Antartica"
    fmt.Println(addr)

    // nesting struct is also possible
    type Info struct {
        employee Employee
        address Address
    }
    // assign value to nest
    info := Info {
        employee: Employee {
            "Rudy",
            12,
        },
        address: Address {
            "Tokyo, Japan",
            762345,
        },
    }
    // just another way of assigning on nest
    info.employee = Employee {
        "Julie",
        30,
    }
    fmt.Println(info)

    // Promoted fields are anonymous fields
    // which are directly accessible parent anonymous struct
    type A struct {
        num int
        str string
    }

    type B struct {
        name string
        A
    }

    var b B
    b.num = 12 // promoted field
    fmt.Println(b)

    // Imported structures, see records/salary.go
    var salary records.Salary
    salary.Amount = 10000
    salary.Currency = "USD"
    fmt.Println(salary)

    // Structs are comparable if their fields are comparable
    type Animal struct {
        name string
    }

    animal1 := Animal {"Lion"}
    animal2 := Animal {}

    if animal1 == animal2 {
        fmt.Println("equal")
    } else {
        fmt.Println("not equal")
    }

    // ..except for maps. If a struct has field type "map",
    // it is not comparable
}

