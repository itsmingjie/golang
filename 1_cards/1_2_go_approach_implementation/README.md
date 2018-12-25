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

### Type Declaration

```go
type deck []string
```
* We want to "extend" a basic data type and add functionality to it
    - Define a new `deck` type, which is a slice of strings. This tels Go we want to create an array of strings and add a bunch of functions specifically made to work with the type.
    - Function with `deck` as a "receiver" - it's like a "method", a function designed to work with the type specifically

### Receiver Function Declaration

```go
func (d deck) print() {
    ...
}
```
* Any variable of type `deck` now gets access to the `print` "method"
* `d` - The actual copy of the deck we're working with, as in "the instance of a deck". `d` is the reference. Very similar to `this` or `self`.
    - In Go, conventionally we don't name it `this`.
    - Usually (by convention) refer to the receiver as a **1 or 2 letter abbreviation of the type**.