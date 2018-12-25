# 1.0 Variables & Functions
12/24/2018

## Method Descriptions
* `newDeck`
    - Create a list of playing cards. Essentially and array of strings.
* `print`
    - Log out of the contents of a deck of cards
* `shuffle`
    - Shuffles all the cards in a deck
* `deal`
    - Create a "hand" of cards
* `saveToFile`
    - Save a list of cards to a file on the local machine
* `newDeckFromFile`
    - Load a list of cards from the local machine

## Variable Declaration
```go
var card string = "Ace of Spades"
```
* `var` - creating a new variable
* `card` - variable identifier
* `string` - variable datatype

### "Should omit type string?"
* Dynamicly Typed Languages
    - Javascript, Ruby, Python
    - Developer don't care about what value is assigned to a variable
    ```javascript
    var number = 123 // integer passed in
    number = "abcd" // reassigned a string

    // number >> "abcd"
    ```
* Staticly Typed Languages
    - C++, Java, Go
    - Variables are declared to datatype and data passed in always have to match the datatype
* **Code above should be rewritten**:
    ```go
    card := "Ace of Spades"
    ```
    - The data already indicates that "card" is supposed to contain String
    - `:=` syntax tells the interpreter to figure out datatype based on data passed in
    - `:=` only used when declaring a **new variable**
        * To reassign data, use `=`
            ```go
            card = "Five of Diamonds"
            ```

### GoLang Data Types
* `bool` - true, false
* `string` - "Hi"   
* `int` - 0, -10000, 99999
* `float64` - 10.000001, 0.00009, -100.003

### Additional Notes
* A variable cannot be initialized with data outside a function, but it can be declared without any value
    ```go
    var num int

    func main() {
        num = 3
        ...
    }
    ```

## Function Declaration
```go
func newCard() string {
    ...
}
```
* `func` - start function declaration
* `newCard` - name of new function
* `string` - return type
    - If being directly assigned to a `:=` syntax, the return type will still be picked up as data type

