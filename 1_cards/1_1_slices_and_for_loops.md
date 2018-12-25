# 1.1 Slices & For Loops
12/24/2018

## Slice
* Array?
    - Array **is a fixed length** list of things
    - A **slice** is an array that can **grow or shrink**
* Every element in a slice must be of **same type**

### Slice Declaration
```go
cards := []string{"Ace of Diamonds", newCard()}
```
* Warning: "declared and not used"
    - Go will call out any unused variables, must do something with it

### Append
```go
cards = append(cards, "Six of Spades")
```
- `append` does NOT modify the original slice, but it creates a new one.
- Therefore, if we're looking to "modify" the original, the result of `append` needs to replace the original slice.

### For Loop - Range
```go
for i, card := range cards {
    fmt.Println(i, card)
}
```
* `i` - the index of the current element in the array
* `card` - the value of the current element in the array
* `range cards` - keyword that takes the slice `cards` and loops over it

#### Why `:=`?
By using the `:=`, we are able to reinitialize the variables `i`, and `cards` while iterating through the slice.