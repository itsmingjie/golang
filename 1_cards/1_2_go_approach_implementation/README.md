# 1.2 Go Approach Implementation
12/25/2018

**Go is not an object-oriented language.**

## Object-Oriented Approach

* Deck Class > Deck Instance
    - Instance Variable
        * `cards` - string
    - Methods
        * `print`
        * `shuffle`
        * `saveToFile`

## Go Approach

* We want to "extend" a basic data type and add functionality to it
    - Define a new `deck` type, which is a slice of strings. This tels Go we want to create an array of strings and add a bunch of functions specifically made to work with the type.
        ```go
        type deck []string
        ```
    - Function with `deck` as a "receiver" - it's like a "method", a function designed to work with the type specifically

### Function Declaration
```go
func (d deck) print() {
    ...
}
```