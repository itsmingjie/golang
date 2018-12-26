# 1.6 Writing Tests
12/26/2018

* Testing with Go is not like RSpec, mocha, jasmine, selenium, etc.
* To make a test, create a new file ending in `..._test.go`
* To run **all tests** in a package, run the command `go test`

## What to test?

* Think about the function - what do we care about?
    - `newDeck()` function
        * The size of the returned deck is exactly n elements
        * Make sure that the first card is an Ace of Spades
        * Make sure that the last card is the King of Clubs

## Testing flow

* `newDeck()` test
    - Create a new deck
    - Write new statement to see if the deck has the right number of cards
    - If it doesn't, tell the Go test handler that something went wrong

## Writing a test
```go
func TestNewDeck(t *testing.T) {
    ...
}
```
* Name the function with `Test...()`
* `t *testing.T` - Test handler, called automatically by the Go test handler
    - When something went wrong, use
        ```
        t.Errorf("Error message %v", ...)
        ```
        * Formatted string - `%v` will be replaced by the arguments that comes after

## Testing File I/O

* File I/O testing flow
    - Create a deck
    - Save to file
    - File created
    - Attempt to load file
    - Error, crashed
        * The file is created... but it's never cleaned up - may interfere with the next test.
* Improved flow - static file name to avoid duplicates
    - Delete any files in current working directory with the name "_decktesting"
    - Create a deck
    - Save to file "_decktesting"
    - Load from file
    - Assert deck length
    - Delete any files in current working directory with the name "_decktesting"
* Deleting a file
    - Use the `Remove` function in the `os` standard package
        ```go
        func Remove(name string) error
        ```
    - Called at start and end of the test