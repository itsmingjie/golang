# 1.4 File IO
12/25/2018

## `WriteFile` function

```go
func WriteFile(filename string, data []byte, perm os.FileMode) error
```
* `WriteFile` is a function from the standard package `io/ioutil`
* Can only use when can convert the data into a **byte slice**
    - A byte slice is a slice of **ASCII values** from a certain string
    - In the case of the cards project, we need to run the deck type into a byte slice
    - Type Conversion
        * Used to convert a type of data to another
            ```go
            []byte("Hi there!") // converts a string to a byte slice
            ```
        * In the case of the deck, we have a `deck`. We need to turn the `deck` into a slice of strings, then turn the slice into a single string, and convert the string into a slice of bytes.
            - It's relatively easy to convert a string to a byte slice, so it'd be helpful enough if a helper function can just return a string, since that way it can be used by future methods too
                ``` go
                []string(d) // d is a type "deck"
                ```
                * This conversion is possible because `deck` **is a** slice of strings. You **can** convert a type back to a basic type.
            - Need to use the `Join` function from the `strings` package to combine strings
                ```go
                func Join(a []string, sep string) string
                ```
                * Concatenates the elements of the first argument into a single string, separated by the second argument
                * Importing Multiple Libraries
                    ```go
                    import ( // no separators in between
                        "fmt"
                        "io/ioutil" // ioutil is a subpackage of the io package
                        "strings"
                        ...
                    )
                    ```
* Permission `0666` sets the permission to a basic readable & writable permission level


## `ReadFile` Function
```go
func ReadFile(filename string) ([]byte, error)
```
* `ReadFile` is a function from the standard package `io/ioutil`
* Reads the file named by filename and returns the contents in a byte slice
* `error` - the error being generated from the function call
    - Value of type `error` is the error information
    - if no error is being generated, the value will be `nil`
        * Usually check with an if statement to make sure no error generated
            ```go
            if err != nil {
                ... // execute only if there's something wrong
            }
            ```
            - Option #1 - log the error and do something (e.g. return a new deck/empty deck)
            - **Option #2 - log the error, quit the program entirely**
                * Utilize `Exit` function from the `os` package to quit the program
                * If `Exit` has an argument of 0, the program ran successfully
                * Otherwise it will indicate that something went wrong
                ```go
                os.Exit(1)
                ```
* Need to convert the byte slice back into a string, then into a string slice, then a deck that's returned
    - Can convert the byte slice back to a string
        ```go
        string(bs)
        ```
    - Need to use the `Split` function from the `strings` package to convert the string into a string slice
        ```go
        s := strings.Split(string(bs))
        ```
    - Then cast the string slice back to a `deck` (which is a string slice)
        ```go
        deck(s)
        ```