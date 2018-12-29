# 2.2 Pointers
12/28/2018

## RAM & Pointers

* The RAM contains a lot of fields, each with a unique address and has a specific value.
* When a new variable is initialized, the data is stored into a container as the value, and the variable is pointed at the value.
* Go is a **pass by value language**
    - When a variable is passed into a function, Go is going to copy the value of that variable and place it into a new container.
    - The function above is in fact **modifying the copy** of the original struct, instead of the original version.

### The Problem
```go
// does not update correctly! A copy is constructed for the person and is updated
func (p person) updateName(newFirstName string) {
    p.firstName = newFirstName
}
```

### The Solution
```go
jimPointer := &jim
...

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
```

## Pointer Operators

* `&variable` - operator that gives access to the memory address of the value this variable is pointing at
* `*pointer` - operator that gives the value this memory address is pointing at
    - In the code above, `*pointerToPerson` is the address that the variable `jim` exists at. Therefore, it has an actual struct stored.
    - **However, if `*` is placed in where a type should be, e.g. `func (pointerToPerson *person)`, it represents a description of the type, meaning that we're working with a pointer to person.**
* In short...
    - Turn **address** into **value** with `*address`
    - Turn **value** into **address** with `&value`

## Shortcut

```go
jim.updateName("Jimmy")
...

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
```
* With go, if a receiver function with type of a pointer is declared, it is allowed for the function to be either called by a pointer **or the variable itself**.
    - `updateName` can be called by either a type `person` or a type `*person`