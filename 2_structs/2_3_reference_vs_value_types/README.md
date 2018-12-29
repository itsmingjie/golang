# 2.3 Reference vs Value Types
12/28/2018

## Arrays vs Slices

* Arrays: cannot be resized
    - **Primitive** data structure
    - Can't be resized
    - Rarely used directly
* Slices: can grow or shrink
    - Used 99% of the time for lists of elements
    - Has three elements inside
        * A pointer to the head of the array
        * Capacity
        * Length

## "Big Gotcha"

When updating a slice element in a function, a function will get direct access to the slice data, modifying it directly.
```go
func main() {
    m := []string{"Hi", "There", "How", "Are", "You"}

    updateSlice(m)

    fmt.Println(m)
}

func updateSlice(s []string) {
    s[0] = "Bye"
}
```
* In this case `m[0]` will be **successfully updated** and replaced with "bye" even if no pointer operations are used.
* Because when the slice is being duplicated in the function, although the slice is duplicated, the pointer to the head of the actual data array is still pointed towards the same original array, which makes it possible for us to directly update the values.

## Reference Types
* maps, channels, pointers, functions, & slices are all reference types.
    - Don't worry about pointers with these!
* ints, floats, strings, bools, and structs are all value types.
    - Use pointers to change these things in a function.