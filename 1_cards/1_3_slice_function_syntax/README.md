# 1.3 Slice & Function Syntax
12/25/2018

**Goal: Split the deck so that a card and the rest of the deck is separated**

## Slice Syntax

* Index starts at 0 (as always)
* Accessing element in slice: `array[0]`
* Selecting a portion of the slice
    ```go
    array[startingIncluding : upToNotIncluding]
    ```
    - Can optionally leave out the number at beginning or end to select the rest
        * `array[2:]` - everything from index 2 until the end
        * `array[:5]` - everything from beginning until index **4**

## Function Syntax

```go
func deal(d deck, handSize int) (deck, deck) {
    ...
}
```
* `d deck` - argument taken by the function, calls must match the data type(s) of the function declaration
* `(deck, deck)` - Go has support of **multiple return values**
* You have free access to functions constructed in the same package
* Retrieving multiple return data
    - Compiler will still determine the data type of the variables
```go
hand, remainingCards := deal(cards, 5)
```
