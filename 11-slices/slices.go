package main

import "fmt"

func main() {
    // "Slices" are "references" to existing array
    // Slices do not have their own data
    a := [5]int{1, 2, 3, 4, 5}
    var b []int = a[1:4] // creates slice from a[1] to a[3]
    fmt.Println(b)

    // modify the original array using their slice
    c := [...]int{1, 2, 3, 4, 5, 6}
    cSlice := c[2:4]
    fmt.Println("before:", c)
    cSlice[0] = 30
    cSlice[1] = 40
    fmt.Println("after:", c)

    // create a slice of entire array
    cWholeSlice := c[:]
    fmt.Println(cWholeSlice)
    // NOTE: array[first param is 0 by default, second param is len(array)]

    // len() of slice is the number of elements in slice

    // cap() of slice is the number of elements in the original array
    // starting from the index where the slice is created
    fmt.Printf("len() of cSlice: %d cap() of cSlice: %d\n", len(cSlice), cap(cSlice))

    fmt.Println(cSlice)
    cSlice = cSlice[:cap(cSlice)] // re-slice cSlice from start to its max index
    fmt.Println(cSlice)

    // create slice using make()
    // by passing the type, length and capacity
    d := make([]int, 5, 5)
    fmt.Println(d)

    // array length is fixed
    // slice length is dynamic by appending new element(s) using append()
    e := []string{"John", "Jane"}
    fmt.Println("Names starts with J:", e)
    fmt.Println("Initial capacity of e is", cap(e))
    e = append(e, "Joe")
    // when append() is used, it creates a new slice IF..
    // the capacity of the slice is NOT sufficient
    // https://golang.org/pkg/builtin/#append
    fmt.Println("Names starts with J now is:", e)
    fmt.Println("New capacity of e is", cap(e))

    // NOTE: when the capacity of the slice is reached
    // upon append(), the capacity is automatically doubled.
    // Growing slices means more allocation.

    // ... operator can be used to append one slice to another
    arr1 := []string{"cat", "dog"}
    arr2 := []string{"snake"}
    animals := append(arr1, arr2...)
    fmt.Println(animals)

    nos := []int{1, 2, 3}
    fmt.Println(nos)

    // slices are referencing arrays
    // so, even if passed on to a function
    // the original array will still be update
    addOne(nos)
    fmt.Println(nos)

    // NOTE: most of the times slices are being used.
    // The only time you are dealing with an array
    // is when you create it with a size. Eg.:
    // variable := [2]int{}
    //
    // The rest is a slice.

    // NOTE: an array cannot be garbage collected if
    // an slice is still in the memory. To solve this issue,
    // we can copy() the slice to a new slice, so the 
    // original array will be garbage collected.
    countries := []string{"USA", "Singapore", "Germany"}
    neededCountries := countries[:2]
    countriesCopy := make([]string, len(neededCountries))
    copy(countriesCopy, neededCountries)

    fmt.Println(countriesCopy)
}

func addOne(numbers []int) {
    for i := range numbers {
        numbers[i] += 1
    }

    // in this case, 6 is not visible outside the functions
    // because append() creates a "new" slice, as mentioned above.
    numbers = append(numbers, 6)
}

